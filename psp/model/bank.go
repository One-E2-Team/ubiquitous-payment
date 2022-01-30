package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bank struct {
	ID        primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	PANPrefix string             `json:"panPrefix"`
	URL       string             `json:"url"`
	IsActive  bool               `json:"isActive"`
}
