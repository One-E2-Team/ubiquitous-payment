package handler

import (
	"net/http"
	"ubiquitous-payment/bank/dto"
	"ubiquitous-payment/util"
)

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var requestDTO dto.RegistrationDTO
	err := util.UnmarshalRequest(r, &requestDTO)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"Register", loggingService)
		return
	}

	err = handler.BankService.Register(requestDTO, w)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"Register", loggingService)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set(util.ContentType, util.ApplicationJson)
	_, _ = w.Write([]byte("{\"message\":\"ok\"}"))
}
