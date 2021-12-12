package service

import (
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/util"
)

func (service *Service) GetUserByID(userID string) (*model.User, error) {
	return service.PSPRepository.GetUserByID(util.String2MongoID(userID))
}
