package model

import (
	"gorm.io/gorm"
	"time"
	"ubiquitous-payment/webshop/model"
)

type PccOrder struct {
	gorm.Model
	AcquirerTransactionId     uint 					`json:"acquirerTransactionId"`
	AcquirerTimestamp         time.Time 			`json:"acquirerTimestamp"`
	Amount					  float32   			`json:"amount"`
	Currency			      string			    `json:"currency"`
	IssuerPAN			      string			    `json:"issuerPan"`
	IssuerCVC			      string			    `json:"issuerCvc"`
	IssuerValidUntil		  string			    `json:"issuerValidUntil"`
	IssuerHolderName		  string			    `json:"issuerHolderName"`
	IssuerOrderId		  	  uint			        `json:"issuerOrderId"`
	IssuerTimestamp           time.Time 			`json:"issuerTimestamp"`
	OrderStatus				  model.OrderStatus		`json:"orderStatus"`
}
