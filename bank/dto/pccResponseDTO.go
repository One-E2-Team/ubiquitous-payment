package dto

import (
	"time"
	"ubiquitous-payment/bank/model"
)

type PccResponseDTO struct {
	IssuerOrderId   uint                    `json:"issuerOrderId"`
	IssuerTimestamp time.Time               `json:"issuerTimestamp"`
	OrderStatus     model.TransactionStatus `json:"orderStatus"`
}
