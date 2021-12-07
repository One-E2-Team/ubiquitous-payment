package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recurring struct {
	ID                      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Type                    RecurringType      `json:"type"`
	InstallmentCount        uint               `json:"installmentCount"`
	DelayedInstallmentCount uint               `json:"delayedInstallmentCount"`
}
