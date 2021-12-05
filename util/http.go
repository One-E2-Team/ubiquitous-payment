package util

import (
	"fmt"
	"net/http"
)

const (
	HttpGet    = "GET"
	HttpPost   = "POST"
	HttpPut    = "PUT"
	HttpDelete = "DELETE"
)

func HandleErrorInHandler(err error, responseWriter http.ResponseWriter) {
	fmt.Println(err)
	responseWriter.WriteHeader(http.StatusBadRequest)
}
