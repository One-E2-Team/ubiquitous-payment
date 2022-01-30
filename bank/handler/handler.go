package handler

import "ubiquitous-payment/bank/service"

type Handler struct {
	BankService *service.Service
}
