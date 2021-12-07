package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	HttpGet         = "GET"
	HttpPost        = "POST"
	HttpPut         = "PUT"
	HttpDelete      = "DELETE"
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

func HandleErrorInHandler(err error, responseWriter http.ResponseWriter) {
	fmt.Println(err)
	responseWriter.WriteHeader(http.StatusBadRequest)
}

func MarshalResult(w http.ResponseWriter, result interface{}) {
	w.Header().Set("Content-Type", "application/json")
	js, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)
}

func UnmarshalResponse(resp *http.Response, result interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if err = json.Unmarshal(body, &result); err != nil {
		return  err
	}
	return nil
}
