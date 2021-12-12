package wsutil

import "ubiquitous-payment/webshop/service"

type WebShopUtilService struct {
	WSService *service.Service
}

var UtilService WebShopUtilService

func InitWebShopUtilService(wsService *service.Service) {
	UtilService = WebShopUtilService{WSService: wsService}
}
