package handler

import (
	"net/http"
	"os"
	"ubiquitous-payment/bank/service"
)

type Handler struct {
	BankService *service.Service
}

var loggingService = "bank_" + os.Getenv("PAN_PREFIX")
var loggingClass = "Handler."

func (handler *Handler) Test(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	a := "wasssuuuup" + os.Getenv("PAN_PREFIX")
	_, err := w.Write([]byte(a))
	if err != nil {
		return
	}
}
