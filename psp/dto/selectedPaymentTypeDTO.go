package dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type SelectedPaymentTypeDTO struct {
	ID					primitive.ObjectID `json:"id,omitempty"`
	PaymentTypeName		string			   `json:"name"`
}
