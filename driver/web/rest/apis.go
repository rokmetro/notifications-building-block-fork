/*
 *   Copyright (c) 2020 Board of Trustees of the University of Illinois.
 *   All rights reserved.

 *   Licensed under the Apache License, Version 2.0 (the "License");
 *   you may not use this file except in compliance with the License.
 *   You may obtain a copy of the License at

 *   http://www.apache.org/licenses/LICENSE-2.0

 *   Unless required by applicable law or agreed to in writing, software
 *   distributed under the License is distributed on an "AS IS" BASIS,
 *   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *   See the License for the specific language governing permissions and
 *   limitations under the License.
 */

package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"notifications/core"
	"notifications/core/model"
	"strings"
)

// ApisHandler handles the rest APIs implementation
type ApisHandler struct {
	app *core.Application
}

// NewApisHandler creates new rest Handler instance
func NewApisHandler(app *core.Application) ApisHandler {
	return ApisHandler{app: app}
}

type getMessagesRequestBody struct {
	IDs []string `json:"ids"`
} // @name getMessagesRequestBody

type storeTokenBody struct {
	PreviousToken *string `json:"previous_token"`
	Token         *string `json:"token"`
} // @name storeTokenBody

type tokenBody struct {
	Token *string `json:"token"`
} // @name tokenBody

// Version gives the service version
// @Description Gives the service version.
// @Tags Client
// @ID Version
// @Produce plain
// @Success 200
// @Security RokwireAuth
// @Router /version [get]
func (h ApisHandler) Version(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte(h.app.Services.GetVersion()))
}

// StoreFirebaseToken Sends a message to a user, list of users or a topic
// @Description Stores a firebase token and maps it to a idToken if presents
// @Tags Client
// @ID Token
// @Param data body storeTokenBody true "body json"
// @Accept  json
// @Success 200
// @Security RokwireAuth UserAuth
// @Router /token [post]
func (h ApisHandler) StoreFirebaseToken(user *model.CoreToken, w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error on marshal token data - %s\n", err.Error())
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var tokenBody storeTokenBody
	err = json.Unmarshal(data, &tokenBody)
	if err != nil {
		log.Printf("Error on unmarshal the create student guide request data - %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if tokenBody.Token == nil || len(*tokenBody.Token) == 0 {
		log.Printf("token is empty or null")
		http.Error(w, "token is empty or null\n", http.StatusBadRequest)
		return
	}

	err = h.app.Services.StoreFirebaseToken(*tokenBody.Token, tokenBody.PreviousToken, user)
	if err != nil {
		log.Printf("Error on creating student guide: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Subscribe Subscribes the current user to a topic
// @Description Subscribes the current user to a topic
// @Tags Client
// @ID Subscribe
// @Param topic path string true "topic"
// @Param data body storeTokenBody true "body json"
// @Accept  json
// @Success 200
// @Security RokwireAuth UserAuth
// @Router /topic/{topic}/subscribe [post]
func (h ApisHandler) Subscribe(user *model.CoreToken, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	topic := params["topic"]
	if len(topic) == 0 {
		log.Println("topic is required")
		http.Error(w, "topic is required", http.StatusBadRequest)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error on reading message data - %s\n", err.Error())
		http.Error(w, fmt.Sprintf("Error on reading message data - %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	var body tokenBody
	err = json.Unmarshal(data, &body)
	if err != nil {
		log.Printf("Error on unmarshal the message request data - %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(*body.Token) == 0 {
		log.Printf("Missing token in the body")
		http.Error(w, "Missing token in the body", http.StatusBadRequest)
		return
	}

	err = h.app.Services.SubscribeToTopic(*body.Token, user, topic)
	if err != nil {
		log.Printf("Error on subscribe to topic (%s): %s\n", topic, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Unsubscribe Unsubscribes the current user to a topic
// @Description Unsubscribes the current user to a topic
// @Tags Client
// @ID Unsubscribe
// @Param topic path string true "topic"
// @Param data body storeTokenBody true "body json"
// @Success 200
// @Security RokwireAuth UserAuth
// @Router /topic/{topic}/unsubscribe [post]
func (h ApisHandler) Unsubscribe(user *model.CoreToken, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	topic := params["topic"]
	if len(topic) == 0 {
		log.Println("topic is required")
		http.Error(w, "topic is required", http.StatusBadRequest)
		return
	}

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error on reading message data - %s\n", err.Error())
		http.Error(w, fmt.Sprintf("Error on reading body data - %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	var body tokenBody
	err = json.Unmarshal(data, &body)
	if err != nil {
		log.Printf("Error on unmarshal the body request data - %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(*body.Token) == 0 {
		log.Printf("Missing token in the json body")
		http.Error(w, "Missing token in the json body", http.StatusBadRequest)
		return
	}

	err = h.app.Services.UnsubscribeToTopic(*body.Token, user, topic)
	if err != nil {
		log.Printf("Error on unsubscribe to topic (%s): %s\n", topic, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// GetUserMessages Gets all messages for the user
// @Description Gets all messages to the authenticated user.
// @Tags Client
// @ID GetUserMessages
// @Param offset query string false "offset"
// @Param limit query string false "limit - limit the result"
// @Param order query string false "order - Possible values: asc, desc. Default: desc"
// @Param start_date query string false "start_date - Start date filter in milliseconds as an integer epoch value"
// @Param end_date query string false "end_date - End date filter in milliseconds as an integer epoch value"
// @Param data body getMessagesRequestBody false "body json of the all message ids that need to be filtered"
// @Accept  json
// @Success 200 {array} model.Message
// @Security UserAuth
// @Router /messages [get]
func (h ApisHandler) GetUserMessages(user *model.CoreToken, w http.ResponseWriter, r *http.Request) {
	offsetFilter := getInt64QueryParam(r, "offset")
	limitFilter := getInt64QueryParam(r, "limit")
	orderFilter := getStringQueryParam(r, "order")
	startDateFilter := getInt64QueryParam(r, "start_date")
	endDateFilter := getInt64QueryParam(r, "end_date")

	var messageIDs []string
	bodyData, _ := ioutil.ReadAll(r.Body)
	if bodyData != nil {
		var body getMessagesRequestBody
		bodyErr := json.Unmarshal(bodyData, &body)
		if bodyErr == nil {
			messageIDs = body.IDs
		}
	}

	var err error
	var messages []model.Message
	if user != nil {
		messages, err = h.app.Services.GetMessages(user.UserID, messageIDs, startDateFilter, endDateFilter, nil, offsetFilter, limitFilter, orderFilter)
		if err != nil {
			log.Printf("Error on getting user messages: %s", err)
			http.Error(w, fmt.Sprintf("Error on getting user messages: %s", err), http.StatusInternalServerError)
			return
		}
	}
	if messages == nil {
		messages = []model.Message{}
	}

	data, err := json.Marshal(messages)
	if err != nil {
		log.Printf("Error on marshal messages: %s\n", err)
		http.Error(w, fmt.Sprintf("Error on marshal messages: %s\n", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// GetTopics Gets all topics
// @Description Gets all topics
// @Tags Client
// @ID GetTopics
// @Success 200 {array} model.Topic
// @Security RokwireAuth
// @Router /topics [get]
func (h ApisHandler) GetTopics(_ *model.CoreToken, w http.ResponseWriter, _ *http.Request) {

	topics, err := h.app.Services.GetTopics()
	if err != nil {
		log.Printf("Error on retrieving all topics: %s\n", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	data, err := json.Marshal(topics)
	if err != nil {
		log.Println("Error on marshal topics")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// GetTopicMessages Gets all messages for topic
// @Description Gets all messages for topic
// @Tags Client
// @ID GetTopicMessages
// @Param topic path string true "topic"
// @Param offset query string false "offset"
// @Param limit query string false "limit - limit the result"
// @Param order query string false "order - Possible values: asc, desc. Default: desc"
// @Param start_date query string false "start_date - Start date filter in milliseconds as an integer epoch value"
// @Param end_date query string false "end_date - End date filter in milliseconds as an integer epoch value"// @Produce plain
// @Success 200 {array} model.Message
// @Security RokwireAuth UserAuth
// @Router /topic/{topic}/messages [get]
func (h ApisHandler) GetTopicMessages(_ *model.CoreToken, w http.ResponseWriter, r *http.Request) {
	offsetFilter := getInt64QueryParam(r, "offset")
	limitFilter := getInt64QueryParam(r, "limit")
	orderFilter := getStringQueryParam(r, "order")
	startDateFilter := getInt64QueryParam(r, "start_date")
	endDateFilter := getInt64QueryParam(r, "end_date")

	params := mux.Vars(r)
	topic := params["topic"]
	if len(topic) == 0 {
		log.Println("topic is required")
		http.Error(w, "topic is required", http.StatusBadRequest)
		return
	}

	messages, err := h.app.Services.GetMessages(nil, nil, startDateFilter, endDateFilter, &topic, offsetFilter, limitFilter, orderFilter)
	if err != nil {
		log.Printf("Error on getting messages: %s", err)
		http.Error(w, fmt.Sprintf("Error on getting messages: %s", err), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(messages)
	if err != nil {
		log.Printf("Error on marshal messages: %s\n", err)
		http.Error(w, fmt.Sprintf("Error on marshal messages: %s\n", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// GetMessage Retrieves a message by id
// @Description Retrieves a message by id
// @Tags Client
// @ID GetUserMessage
// @Param id path string true "id"
// @Accept  json
// @Produce plain
// @Success 200 {object} model.Message
// @Security UserAuth
// @Router /message/{id} [get]
func (h ApisHandler) GetMessage(user *model.CoreToken, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if len(id) == 0 {
		log.Println("Message id is required")
		http.Error(w, "Message id is required", http.StatusBadRequest)
		return
	}

	message, err := h.app.Services.GetMessage(id)
	if err != nil {
		log.Printf("Error on get message with id (%s): %s\n", id, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if message == nil || !message.HasUser(user) {
		log.Printf("Error on get message with id (%s): %s\n", id, err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	data, err := json.Marshal(message)
	if err != nil {
		log.Println("Error on marshal message")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// DeleteUserMessages Removes the current user from the recipient list of all described messages
// @Description Removes the current user from the recipient list of all described messages
// @Tags Client
// @ID DeleteUserMessages
// @Param data body getMessagesRequestBody false "body json of the all message ids that need to be filtered"
// @Accept  json
// @Success 200
// @Security UserAuth
// @Router /messages [delete]
func (h ApisHandler) DeleteUserMessages(user *model.CoreToken, w http.ResponseWriter, r *http.Request) {
	var messageIDs []string
	bodyData, _ := ioutil.ReadAll(r.Body)
	if bodyData != nil {
		var body getMessagesRequestBody
		bodyErr := json.Unmarshal(bodyData, &body)
		if bodyErr == nil {
			messageIDs = body.IDs
		}
	}

	errStrings := []string{}
	if len(messageIDs) > 0 {
		for _, id := range messageIDs {
			err := h.app.Services.DeleteUserMessage(user, id)
			if err != nil {
				errStrings = append(errStrings, fmt.Sprintf("%s\n", err.Error()))
				log.Printf("Error on delete message with id (%s) for recipuent (%s): %s\n", id, *user.UserID, err)
			}
		}
	} else {
		log.Printf("Missing ids inthe request body")
		http.Error(w, "Missing ids inthe request body", http.StatusBadRequest)
		return
	}
	if len(errStrings) > 0 {
		http.Error(w, strings.Join(errStrings, ""), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// CreateMessage Creates a message
// @Description Creates a message
// @Tags Client
// @ID createMessage
// @Accept  json
// @Param data body model.Message true "body json"
// @Success 200 {object} model.Message
// @Security UserAuth
// @Router /message [post]
func (h ApisHandler) CreateMessage(user *model.CoreToken, w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error on reading message data - %s\n", err.Error())
		http.Error(w, fmt.Sprintf("Error on reading message data - %s\n", err.Error()), http.StatusBadRequest)
		return
	}

	var message *model.Message
	err = json.Unmarshal(data, &message)
	if err != nil {
		log.Printf("Error on unmarshal the message request data - %s\n", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	message, err = h.app.Services.CreateMessage(user, message)
	if err != nil {
		log.Printf("Error on create message: %s\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err = json.Marshal(message)
	if err != nil {
		log.Println("Error on marshal message")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// DeleteUserMessage Removes the current user from the recipient list of the message
// @Description Removes the current user from the recipient list of the message
// @Tags Client
// @ID DeleteUserMessage
// @Param id path string true "id"
// @Produce plain
// @Success 200
// @Security UserAuth
// @Router /message/{id} [delete]
func (h ApisHandler) DeleteUserMessage(user *model.CoreToken, w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if len(id) == 0 {
		log.Println("Message id is required")
		http.Error(w, "Message id is required", http.StatusBadRequest)
		return
	}

	err := h.app.Services.DeleteUserMessage(user, id)
	if err != nil {
		log.Printf("Error on delete message with id (%s) for recipuent (%s): %s\n", id, *user.UserID, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
