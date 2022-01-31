package service

import (
	"ubiquitous-payment/bank/repository"
)

type Service struct {
	BankRepository *repository.Repository
}
