package service

import (
	"github.com/google/uuid"
	"ubiquitous-payment/psp/model"
)

func (service *Service) CreateEmptyTransaction() (string, error) {
	orderId := uuid.NewString()
	err := service.PSPRepository.CreateTransaction(&model.Transaction{PSPId: orderId})
	if err != nil {
		return "", err
	}
	return orderId, nil
}
