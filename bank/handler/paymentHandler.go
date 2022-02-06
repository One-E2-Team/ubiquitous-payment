package handler

import (
	"errors"
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
	redirectUrl := handler.BankService.Pay(issuerCard, paymentUrlId)
	if redirectUrl == "" {
		util.HandleErrorInHandler(errors.New("redirect url from acquirer is empty"), w, loggingClass+"Pay", loggingService)
		return
	}
	util.MarshalResult(w, redirectUrl)
}

func (handler *Handler) IssuerPay(w http.ResponseWriter, r *http.Request) {
	var pccOrderDto dto.PccOrderDTO
	err := util.UnmarshalRequest(r, &pccOrderDto)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"IssuerPay", loggingService)
		return
	}

	pccResponse := handler.BankService.IssuerPay(pccOrderDto)
	util.MarshalResult(w, pccResponse)
}

func (handler *Handler) GetPaymentDetails(w http.ResponseWriter, r *http.Request) {
	paymentDetails, err := handler.BankService.GetPaymentDetails(util.GetPathVariable(r, "payment-url-id"))
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetPaymentDetails", loggingService)
		return
	}

	util.MarshalResult(w, paymentDetails)
}
