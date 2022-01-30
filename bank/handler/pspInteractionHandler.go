package handler

import (
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

	response, err := handler.BankService.PspRequest(mapper.PspRequestDTOToTransaction(pspRequest))
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"PspRequest", loggingService)
		return
	}
	util.MarshalResult(w, response)
}
