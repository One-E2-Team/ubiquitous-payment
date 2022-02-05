package handler

import (
	"net/http"
	"os"
	"ubiquitous-payment/bank/model"
	"ubiquitous-payment/bank/service"
	"ubiquitous-payment/util"
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

func (handler *Handler) TestEncryption(w http.ResponseWriter, r *http.Request) {
	var test string
	err := util.UnmarshalRequest(r, &test)
	if err != nil {
		w.WriteHeader(401)
		return
	}
	t := model.Test{Name: util.EncryptedString{Data: test}}
	err = handler.BankService.Repository.Create(&t)
	if err != nil {
		w.WriteHeader(402)
		return
	}

	var x []model.Test
	if err = handler.BankService.Repository.Database.Raw("select * from tests").Scan(&x).Error; err != nil {
		w.WriteHeader(403)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain")
	a := "wasssuuuup" + os.Getenv("PAN_PREFIX")
	_, err = w.Write([]byte(a))
	if err != nil {
		return
	}
}
