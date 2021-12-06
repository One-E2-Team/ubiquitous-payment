package util

import (
	"encoding/json"
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

func MarshalResult(w http.ResponseWriter, result interface{}){
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)
}
