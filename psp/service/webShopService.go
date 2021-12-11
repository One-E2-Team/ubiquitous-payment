package service

import "ubiquitous-payment/util"

func (service *Service) ChangeWebShopAcceptance(webShopID string, isAccepted bool) error {
	return service.PSPRepository.ChangeWebShopAcceptance(util.String2MongoID(webShopID), isAccepted)
}
