package model

// Recipient represent recipient of a message
type Recipient struct {
	UID  *string `json:"user_id" bson:"user_id"`
	Name *string `json:"name" bson:"name"`
} //@name Recipient
