package handler

import (
	"encoding/json"
	"net/http"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/dto"
)

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var dto dto.RegistrationDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		util.HandleErrorInHandler(err, w)
		return
	}

	err = handler.WSService.Register(w, dto)

	if err != nil {
		util.HandleErrorInHandler(err, w)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("{\"message\":\"ok\"}"))
	}
}