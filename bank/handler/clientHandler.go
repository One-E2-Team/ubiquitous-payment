package handler

import (
	"net/http"
	"ubiquitous-payment/util"
)

func (handler *Handler) GetMyAccount(w http.ResponseWriter, r *http.Request) {
	account, err := handler.BankService.GetMyAccount(util.GetLoggedUserIDFromToken(r))
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetMyAccount", loggingService)
		return
	}
	util.MarshalResult(w, account)
}

func (handler *Handler) GetMyTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := handler.BankService.GetMyTransactions(util.GetLoggedUserIDFromToken(r))
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetMyTransactions", loggingService)
		return
	}
	util.MarshalResult(w, transactions)
}

func (handler *Handler) GetAllTransactions(w http.ResponseWriter, r *http.Request) {
	transactions, err := handler.BankService.GetAllTransactions()
	if err != nil {
		util.HandleErrorInHandler(err, w, loggingClass+"GetMyTransactions", loggingService)
		return
	}
	util.MarshalResult(w, transactions)
}
