package service

import (
	"ubiquitous-payment/webshop/repository"
)

type Service struct {
	WSRepository *repository.Repository
}