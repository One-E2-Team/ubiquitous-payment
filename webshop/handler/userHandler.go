package handler

import (
	"encoding/json"
	"net/http"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/dto"
)

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var requestDTO dto.RegistrationDTO
	err := json.NewDecoder(r.Body).Decode(&requestDTO)
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"Register", loggingService)
		return
	}

	err = handler.WSService.Register(w, requestDTO)

	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"Register", loggingService)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set(util.ContentType, util.ApplicationJson)
	_, _ = w.Write([]byte("{\"message\":\"ok\"}"))
}
