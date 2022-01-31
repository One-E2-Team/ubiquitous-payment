package handler

import (
	"errors"
	"net/http"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/bank/handler/mapper"
	"ubiquitous-payment/util"
)

func (handler *Handler) PspRequest(w http.ResponseWriter, r *http.Request) {
	var pspRequest dto.PspRequestDTO
	err := util.UnmarshalRequest(r, &pspRequest)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"PspRequest", loggingService)
		return
	}
	t, err := mapper.PspRequestDTOToTransaction(pspRequest)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"PspRequest", loggingService)
		return
	}
	response, err := handler.BankService.PspRequest(t)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"PspRequest", loggingService)
		return
	}
	util.MarshalResult(w, response)
}

func (handler *Handler) CheckPayment(w http.ResponseWriter, r *http.Request) {
	id := util.GetPathVariable(r, "id")
	if id == "" {
		util.HandleErrorInHandler(errors.New("no id path var"), w, loggingClass+"CheckPayment", loggingService)
		return
	}
	response, err := handler.BankService.CheckPaymentStatus(id)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"CheckPayment", loggingService)
		return
	}
	util.MarshalResult(w, response)
}
