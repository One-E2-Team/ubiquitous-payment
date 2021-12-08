package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type PaymentType struct {
	ID   primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Name string      `json:"name"`
}

/*
type PaymentMethod string

const (
	CreditCard PaymentMethod = "creditcard"
	QRCode PaymentMethod = "qrcode"
	Bitcoin PaymentMethod = "bitcoin"
	PayPal PaymentMethod = "paypal"
)*/
