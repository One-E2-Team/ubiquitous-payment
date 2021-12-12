package handler

import (
	"net/http"
	"ubiquitous-payment/psp/service"
)

type Handler struct {
	PSPService *service.Service
}

const loggingService = "psp"
const loggingClass = "Handler."

func (handler *Handler) Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	a := "wasssuuuup"
	_, err := w.Write([]byte(a))
	if err != nil {
		return
	}
}
