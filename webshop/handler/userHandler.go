package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ubiquitous-payment/webshop/dto"
)

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request) {
	var dto dto.RegistrationDto
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{\"message\":\"Server error while decoding.\"}"))
		return
	}

	err = handler.WSService.Register(w, dto)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{\"message\":\"Server error while registering.\"}"))
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("{\"message\":\"ok\"}"))
	}
}