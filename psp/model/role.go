package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"ubiquitous-payment/webshop/model"
)

type Role struct {
	ID   			primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name 			string             `json:"name"`
	Privileges		[]model.Privilege  `json:"privileges"`
}
