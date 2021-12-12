package psputil

import (
	"net/http"
	"ubiquitous-payment/psp/model"
	"ubiquitous-payment/psp/service"
)

type UtilService struct {
	PspService *service.Service
}

var utilService UtilService

func InitPspUtilService(wsService *service.Service) {
	utilService = UtilService{PspService: wsService}
}

//TODO: create general RBAC if possible

func RBAC(handler func(http.ResponseWriter, *http.Request), privilege string, returnCollection bool) func(http.ResponseWriter, *http.Request) {
	finalHandler := func(pass bool) func(http.ResponseWriter, *http.Request) {
		if pass {
			return handler
		} else {
			return func(writer http.ResponseWriter, request *http.Request) {
				writer.WriteHeader(http.StatusOK)
				writer.Header().Set("Content-Type", "application/json")
				if returnCollection {
					_, _ = writer.Write([]byte("[{\"status\":\"fail\", \"reason\":\"unauthorized\"}]"))
				} else {
					_, _ = writer.Write([]byte("{\"status\":\"fail\", \"reason\":\"unauthorized\"}"))
				}
			}
		}
	}

	return func(writer http.ResponseWriter, request *http.Request) {
		var handleFunc func(http.ResponseWriter, *http.Request)
		id := GetLoggedUserIDFromToken(request)
		if id == "" {
			handleFunc = finalHandler(false)
		} else {
			user, err := utilService.PspService.GetUserByID(id)
			if err != nil {
				handleFunc = finalHandler(false)
			}
			handleFunc = finalHandler(isPrivilegeValid(privilege, *user))
		}
		handleFunc(writer, request)
	}
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
