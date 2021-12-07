package dto

type WebShopOrderDTO struct {
	PspOrderId		string						`json:"pspOrderId"`
	Amount			float32						`json:"amount"`
	Currency		string						`json:"currency"`
	PaymentMode		string						`json:"paymentMode,omitempty"`
	RecurringType	string						`json:"recurringType,omitempty"`
	RecurringTimes	string						`json:"recurringTimes,omitempty"`
	PaymentTo		map[string]interface{}		`json:"paymentTo,omitempty"`
	SuccessUrl		string						`json:"successUrl"`
	FailedUrl		string						`json:"failedUrl"`
	ErrorUrl		string						`json:"errorUrl"`
}
