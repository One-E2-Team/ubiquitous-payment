package service

import (
	"ubiquitous-payment/bank/repository"
)

type Service struct {
	Repository *repository.Repository
}
