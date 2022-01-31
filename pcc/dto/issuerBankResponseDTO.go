package dto

import (
	"time"
	"ubiquitous-payment/pcc/model"
)

type IssuerBankResponseDto struct {
	IssuerOrderId		  	  uint			        `json:"issuerOrderId"`
	IssuerTimestamp           time.Time 			`json:"issuerTimestamp"`
	OrderStatus				  model.OrderStatus	    `json:"orderStatus"`
}
