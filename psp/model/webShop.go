package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type WebShop struct {
	ID             primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name           string             `json:"name"`
	PSPAccessToken string             `json:"pspAccessToken"` //TODO: rename to uuid
	Accepted       bool               `json:"accepted"`
	PaymentTypes   []PaymentType      `json:"paymentTypes"`
	Accounts       []Account          `json:"accounts"`
}
