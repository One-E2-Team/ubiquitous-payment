package util

import (
	"bytes"
	"fmt"
	"net/http"
)

//
//import (
//	"bytes"
//	"fmt"
//	"github.com/dgrijalva/jwt-go"
//	"net/http"
//	"os"
//	"time"
//)
//
var pspAccessToken = ""

func SetPspJwt(pspJwt string) {
	pspAccessToken = pspJwt
}

func PSPRequest(method string, path string, data []byte, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	pspHost, pspPort := GetPSPHostAndPort()
	urlPath := GetPSPProtocol() + "://" + pspHost + ":" + pspPort
	req, err := http.NewRequest(method, urlPath+path, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+pspAccessToken)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return client.Do(req)
}

//
//func PSPAuth(handler func(http.ResponseWriter, *http.Request), webShops []string) func(http.ResponseWriter, *http.Request) {
//
//	finalHandler := func(pass bool) func(http.ResponseWriter, *http.Request) {
//		if pass {
//			return handler
//		} else {
//			return func(writer http.ResponseWriter, request *http.Request) {
//				writer.WriteHeader(http.StatusOK)
//				writer.Header().Set("Content-Type", "application/json")
//				_, _ = writer.Write([]byte("{\"status\":\"fail\", \"reason\":\"unauthorized\"}"))
//			}
//		}
//	}
//
//	return func(writer http.ResponseWriter, request *http.Request) {
//		if check := ValidatePSPToken(request, webShops); check {
//			finalHandler(true)(writer, request)
//		} else {
//			finalHandler(false)(writer, request)
//		}
//	}
//}
