package rbac

import (
	"net/http"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/psputil"
	"ubiquitous-payment/psp/service"
	"ubiquitous-payment/util"
)

type CopyPspService struct {
	RealPspService *service.Service
}

var copyPspService CopyPspService

func InitRbacService(realPspService *service.Service) {
	copyPspService = CopyPspService{RealPspService: realPspService}
}

type PspRbacService struct{}

var pspRbacService PspRbacService

func (PspRbacService) IsPrivilegeValid(privilege string, request *http.Request) bool {
	id := psputil.GetLoggedUserIDFromToken(request)
	if id == "" {
		return false
	}
	user, err := copyPspService.RealPspService.GetUserByID(id)
	if err != nil {
		return false
	}
	return isPrivilegeValid(privilege, *user)
}

func PspRbac(handler func(http.ResponseWriter, *http.Request), privilege string) func(http.ResponseWriter, *http.Request) {
	return util.GenericRBAC(handler, privilege, pspRbacService)
}

func isPrivilegeValid(privilege string, user model.User) bool {
	for _, role := range user.Roles {
		for _, userPrivilege := range role.Privileges {
			if userPrivilege.Name == privilege {
				return true
			}
		}
	}
	return false
}
