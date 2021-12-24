package rbac

import (
	"net/http"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/service"
)

type CopyWebShopService struct {
	RealWebShopService *service.Service
}

var copyWebShopService CopyWebShopService

func InitRbacService(realWebShopService *service.Service) {
	copyWebShopService = CopyWebShopService{RealWebShopService: realWebShopService}
}

type WebShopRbacService struct{}

var webShopRbacService WebShopRbacService

func (WebShopRbacService) IsPrivilegeValid(privilege string, request *http.Request) bool {
	id := util.GetLoggedUserIDFromToken(request)
	if id == 0 {
		return false
	}
	return isPrivilegeValid(privilege, *copyWebShopService.RealWebShopService.GetPrivileges(id))
}

func WebShopRbac(handler func(http.ResponseWriter, *http.Request), privilege string) func(http.ResponseWriter, *http.Request) {
	return util.GenericRBAC(handler, privilege, webShopRbacService)
}

func isPrivilegeValid(privilege string, validPrivileges []string) bool {
	for _, val := range validPrivileges {
		if val == privilege {
			return true
		}
	}
	return false
}
