package util

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
)

func HandleErrorInHandler(err error, responseWriter http.ResponseWriter) {
	fmt.Println(err)
	responseWriter.WriteHeader(http.StatusBadRequest)
}

func MarshalResult(w http.ResponseWriter, result interface{}) {
	w.Header().Set(ContentType, ApplicationJson)
	js, err := json.Marshal(result)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(js)
}

func UnmarshalRequest(request *http.Request, resultObject interface{}) error {
	return json.NewDecoder(request.Body).Decode(&resultObject)
}

func UnmarshalResponse(resp *http.Response, resultObject interface{}) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		_ = Body.Close()
	}(resp.Body)
	if err = json.Unmarshal(body, &resultObject); err != nil {
		return err
	}
	return nil
}
