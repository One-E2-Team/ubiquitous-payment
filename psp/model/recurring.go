package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recurring struct {
	ID                      primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Type                    RecurringType      `json:"type"`
	InstallmentCount        int                `json:"installmentCount"`
	DelayedInstallmentCount int                `json:"delayedInstallmentCount"`
}
