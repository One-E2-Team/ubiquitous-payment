package util

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

const (
	ContentType     = "Content-Type"
	ApplicationJson = "application/json"
	Authorization   = "Authorization"
)

func GetPathVariable(request *http.Request, variableName string) string {
	return mux.Vars(request)[variableName]
}

func HandleErrorInHandler(err error, responseWriter http.ResponseWriter, resourceMethod string, service string) {
	fmt.Println(err)
	Logging(ERROR, resourceMethod, err.Error(), service)
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
