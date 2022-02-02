package bankutil

import (
	"net/http"
	"ubiquitous-payment/util"
)

type BankRbacService struct{}

var bankRbacService BankRbacService

func (BankRbacService) IsPrivilegeValid(privilege string, request *http.Request) bool {
	ok, _ := util.ValidateCsToken(request, []string{privilege})
	return ok
}

func BankRbac(handler func(http.ResponseWriter, *http.Request), privilege string) func(http.ResponseWriter, *http.Request) {
	return util.GenericRBAC(handler, privilege, bankRbacService)
}
