package wsutil

import (
	"net/http"
	"strconv"
	"ubiquitous-payment/util"
	"ubiquitous-payment/webshop/service"
)

type RbacService struct {
	WSService *service.Service
}

var rbacService RbacService

func InitRbacService(wsService *service.Service) {
	rbacService = RbacService{WSService: wsService}
}

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
		writer.Header().Set("initiator", "NO_TOKEN")
		var handleFunc func(http.ResponseWriter, *http.Request)
		id := util.GetLoggedUserIDFromToken(request)
		if id == 0 {
			writer.Header().Set("initiator", "UNAUTHORIZED")
			handleFunc = finalHandler(false)
		} else {
			writer.Header().Set("initiator", strconv.Itoa(int(id)))
			handleFunc = finalHandler(isPrivilegeValid(privilege, *rbacService.WSService.GetPrivileges(id)))
		}
		handleFunc(writer, request)
	}
}

func isPrivilegeValid(privilege string, validPrivileges []string) bool {
	for _, val := range validPrivileges {
		if val == privilege {
			return true
		}
	}
	return false
}
