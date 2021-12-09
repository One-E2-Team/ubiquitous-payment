package service

import (
	"ubiquitous-payment/psp/repository"
)

type Service struct {
	PSPRepository *repository.Repository
}
