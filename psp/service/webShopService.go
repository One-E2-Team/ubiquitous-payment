package service

import "ubiquitous-payment/util"

func (service *Service) ChangeWebShopAcceptance(webShopID string, isAccepted bool) error {
	mongoID, err := util.String2MongoID(webShopID)
	if err != nil {
		return err
	}
	return service.PSPRepository.ChangeWebShopAcceptance(mongoID, isAccepted)
}
