package handler

import (
	"net/http"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/util"
)

func (handler *Handler) Pay(w http.ResponseWriter, r *http.Request) {
	paymentUrlId := util.GetPathVariable(r, "payment-url-id")
	var issuerCard dto.IssuerCardDTO
	err := util.UnmarshalRequest(r, &issuerCard)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"Pay", loggingService)
		return
	}
	paymentDto, err := handler.BankService.Pay(issuerCard, paymentUrlId)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"Pay", loggingService)
		return
	}
	util.MarshalResult(w, paymentDto)
}

func (handler *Handler) IssuerPay(w http.ResponseWriter, r *http.Request) {
	var pccOrderDto dto.PccOrderDTO
	err := util.UnmarshalRequest(r, &pccOrderDto)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"IssuerPay", loggingService)
		return
	}

	pccResponse, err := handler.BankService.IssuerPay(pccOrderDto)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"IssuerPay", loggingService)
		return
	}
	util.MarshalResult(w, pccResponse)
}
