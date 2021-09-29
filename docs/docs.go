// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/message": {
            "put": {
                "security": [
                    {
                        "AdminUserAuth": []
                    }
                ],
                "description": "Updates a message",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "operationId": "UpdateMessage",
                "parameters": [
                    {
                        "description": "body json",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AdminUserAuth": []
                    }
                ],
                "description": "Creates a message",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "operationId": "CreateMessage",
                "parameters": [
                    {
                        "description": "body json",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/admin/message/{id}": {
            "get": {
                "security": [
                    {
                        "AdminUserAuth": []
                    }
                ],
                "description": "Retrieves a message by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Admin"
                ],
                "operationId": "GetMessage",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AdminUserAuth": []
                    }
                ],
                "description": "Deletes a message with id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Admin"
                ],
                "operationId": "DeleteMessage",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/admin/messages": {
            "get": {
                "security": [
                    {
                        "AdminUserAuth": []
                    }
                ],
                "description": "Gets all messages",
                "tags": [
                    "Admin"
                ],
                "operationId": "GetMessages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "user - filter by user",
                        "name": "user",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "topic - filter by topic",
                        "name": "topic",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit - limit the result",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "order - Possible values: asc, desc. Default: desc",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "start_date - Start date filter in milliseconds as an integer epoch value",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "end_date - End date filter in milliseconds as an integer epoch value",
                        "name": "end_date",
                        "in": "query"
                    }
                ]
            }
        },
        "/admin/topic": {
            "put": {
                "security": [
                    {
                        "AdminUserAuth": []
                    }
                ],
                "description": "Updated the topic.",
                "tags": [
                    "Admin"
                ],
                "operationId": "UpdateTopic",
                "parameters": [
                    {
                        "description": "body json",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Topic"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/Topic"
                        }
                    }
                }
            }
        },
        "/admin/topics": {
            "get": {
                "security": [
                    {
                        "AdminUserAuth": []
                    }
                ],
                "description": "Gets all topics",
                "tags": [
                    "Admin"
                ],
                "operationId": "AdminGetTopics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Topic"
                            }
                        }
                    }
                }
            }
        },
        "/int/message": {
            "post": {
                "security": [
                    {
                        "InternalAuth": []
                    }
                ],
                "description": "Sends a message to a user, list of users or a topic",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Internal"
                ],
                "operationId": "InternalSendMessage",
                "parameters": [
                    {
                        "description": "body json",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/message": {
            "post": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "Creates a message",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "operationId": "createMessage",
                "parameters": [
                    {
                        "description": "body json",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            }
        },
        "/message/{id}": {
            "get": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "Retrieves a message by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Client"
                ],
                "operationId": "GetUserMessage",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Message"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "Removes the current user from the recipient list of the message",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Client"
                ],
                "operationId": "DeleteUserMessage",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/messages": {
            "get": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "Gets all messages to the authenticated user.",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "operationId": "GetUserMessages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit - limit the result",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "order - Possible values: asc, desc. Default: desc",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "start_date - Start date filter in milliseconds as an integer epoch value",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "end_date - End date filter in milliseconds as an integer epoch value",
                        "name": "end_date",
                        "in": "query"
                    },
                    {
                        "description": "body json of the all message ids that need to be filtered",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/getMessagesRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Message"
                            }
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "UserAuth": []
                    }
                ],
                "description": "Removes the current user from the recipient list of all described messages",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "operationId": "DeleteUserMessages",
                "parameters": [
                    {
                        "description": "body json of the all message ids that need to be filtered",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/getMessagesRequestBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/token": {
            "post": {
                "security": [
                    {
                        "RokwireAuth UserAuth": []
                    }
                ],
                "description": "Stores a firebase token and maps it to a idToken if presents",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "operationId": "Token",
                "parameters": [
                    {
                        "description": "body json",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/storeTokenBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/topic/{topic}/messages": {
            "get": {
                "security": [
                    {
                        "RokwireAuth UserAuth": []
                    }
                ],
                "description": "Gets all messages for topic",
                "tags": [
                    "Client"
                ],
                "operationId": "GetTopicMessages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "topic",
                        "name": "topic",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit - limit the result",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "order - Possible values: asc, desc. Default: desc",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "start_date - Start date filter in milliseconds as an integer epoch value",
                        "name": "start_date",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "end_date - End date filter in milliseconds as an integer epoch value",
                        "name": "end_date",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Message"
                            }
                        }
                    }
                }
            }
        },
        "/topic/{topic}/subscribe": {
            "post": {
                "security": [
                    {
                        "RokwireAuth UserAuth": []
                    }
                ],
                "description": "Subscribes the current user to a topic",
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Client"
                ],
                "operationId": "Subscribe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "topic",
                        "name": "topic",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body json",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/storeTokenBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/topic/{topic}/unsubscribe": {
            "post": {
                "security": [
                    {
                        "RokwireAuth UserAuth": []
                    }
                ],
                "description": "Unsubscribes the current user to a topic",
                "tags": [
                    "Client"
                ],
                "operationId": "Unsubscribe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "topic",
                        "name": "topic",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "body json",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/storeTokenBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/topics": {
            "get": {
                "security": [
                    {
                        "RokwireAuth": []
                    }
                ],
                "description": "Gets all topics",
                "tags": [
                    "Client"
                ],
                "operationId": "GetTopics",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Topic"
                            }
                        }
                    }
                }
            }
        },
        "/version": {
            "get": {
                "security": [
                    {
                        "RokwireAuth": []
                    }
                ],
                "description": "Gives the service version.",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "Client"
                ],
                "operationId": "Version",
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        }
    },
    "definitions": {
        "CoreUserRef": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "Recipient": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "uid": {
                    "type": "string"
                }
            }
        },
        "Topic": {
            "type": "object",
            "properties": {
                "date_created": {
                    "type": "string"
                },
                "date_updated": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "getMessagesRequestBody": {
            "type": "object",
            "properties": {
                "ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "model.Message": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "data": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "date_created": {
                    "type": "string"
                },
                "date_updated": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "priority": {
                    "type": "integer"
                },
                "recipients": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/Recipient"
                    }
                },
                "sender": {
                    "$ref": "#/definitions/model.Sender"
                },
                "subject": {
                    "type": "string"
                },
                "topic": {
                    "type": "string"
                }
            }
        },
        "model.Sender": {
            "type": "object",
            "properties": {
                "type": {
                    "description": "user or system",
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/CoreUserRef"
                }
            }
        },
        "storeTokenBody": {
            "type": "object",
            "properties": {
                "previous_token": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AdminUserAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header (add admin core access token with Bearer prefix to the Authorization value. The token must contain notifications_admin as a permission)"
        },
        "InternalAuth": {
            "type": "apiKey",
            "name": "INTERNAL-API-KEY",
            "in": "header"
        },
        "RokwireAuth": {
            "type": "apiKey",
            "name": "ROKWIRE-API-KEY",
            "in": "header"
        },
        "UserAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header (add core access token with Bearer prefix to the Authorization value. The token must represent either anonymous or authenticated user )"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "0.1.10",
	Host:        "localhost",
	BasePath:    "/notifications/api",
	Schemes:     []string{"https"},
	Title:       "Rokwire Notifications Building Block API",
	Description: "Rokwire Notifications Building Block API Documentation.",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
