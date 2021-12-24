package util

import (
	"net/http"
)

type IRbac interface {
	IsPrivilegeValid(string, *http.Request) bool
}

func finalHandler(handler func(http.ResponseWriter, *http.Request), pass bool) func(http.ResponseWriter, *http.Request) {
	if pass {
		return handler
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		MarshalResult(writer, "{\"status\":\"fail\", \"reason\":\"unauthorized\"}")
	}
}

func GenericRBAC(handler func(http.ResponseWriter, *http.Request), privilege string, rbacEntity IRbac) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		handleFunc := finalHandler(handler, rbacEntity.IsPrivilegeValid(privilege, request))
		handleFunc(writer, request)
	}
}
