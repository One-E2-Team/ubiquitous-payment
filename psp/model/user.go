package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID               primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Username         string             `json:"username"`
	Password         string             `json:"password"`
	IsDeleted        bool               `json:"isDeleted"`
	ValidationUuid   string             `json:"validationUuid"`
	ValidationExpire time.Time          `json:"validationExpire"`
	Roles            []Role             `json:"roles"`
	WebShopId        string             `json:"webShopId"`
}
