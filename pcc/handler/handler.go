package handler

import (
	"net/http"
	"ubiquitous-payment/pcc/service"
)

var loggingService = "pcc"
var loggingClass = "Handler."

type Handler struct {
	Service *service.Service
}

func (handler *Handler) Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	a := "wasssuuuup"
	_, err := w.Write([]byte(a))
	if err != nil {
		return
	}
}