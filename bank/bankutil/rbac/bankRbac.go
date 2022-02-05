package rbac

import (
	"net/http"
	"ubiquitous-payment/bank/service"
	"ubiquitous-payment/util"
)

type CopyBankService struct {
	RealBankService *service.Service
}

var copyBankService CopyBankService

func InitRbacService(realBankService *service.Service) {
	copyBankService = CopyBankService{RealBankService: realBankService}
}

type BankRbacService struct{}

var bankRbacService BankRbacService

func (BankRbacService) IsPrivilegeValid(privilege string, request *http.Request) bool {
	id := util.GetLoggedUserIDFromToken(request)
	if id == 0 {
		return false
	}
	return isPrivilegeValid(privilege, *copyBankService.RealBankService.GetPrivileges(id))
}

func BankRbac(handler func(http.ResponseWriter, *http.Request), privilege string) func(http.ResponseWriter, *http.Request) {
	return util.GenericRBAC(handler, privilege, bankRbacService)
}

func isPrivilegeValid(privilege string, validPrivileges []string) bool {
	for _, val := range validPrivileges {
		if val == privilege {
			return true
		}
	}
	return false
}
