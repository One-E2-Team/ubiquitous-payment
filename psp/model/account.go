package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Account struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	AccountID   string             `json:"accountId"`
	Secret      string             `json:"secret"`
	PaymentType PaymentType		   `json:"paymentType"`
}
