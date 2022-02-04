package handler

import (
	"net/http"
	"ubiquitous-payment/util"
)

func (handler *Handler) GetMyAccount(w http.ResponseWriter, r *http.Request) {
	accounts, err := handler.BankService.GetMyAccount(util.GetLoggedUserIDFromToken(r))
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetMyAccount", loggingService)
		return
	}
	util.MarshalResult(w, accounts)
}

func (handler *Handler) GetMyTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := handler.BankService.GetMyTransactions(util.GetLoggedUserIDFromToken(r))
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetMyTransactions", loggingService)
		return
	}
	util.MarshalResult(w, transactions)
}
