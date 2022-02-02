package pccutil

import (
	"net/http"
	"ubiquitous-payment/util"
)

type PccRbacService struct{}

var pccRbacService PccRbacService

func (PccRbacService) IsPrivilegeValid(privilege string, request *http.Request) bool {
	ok, _ := util.ValidateCsToken(request, []string{privilege})
	return ok
}

func PccRbac(handler func(http.ResponseWriter, *http.Request), privilege string) func(http.ResponseWriter, *http.Request) {
	return util.GenericRBAC(handler, privilege, pccRbacService)
}
