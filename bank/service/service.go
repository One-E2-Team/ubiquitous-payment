package service

import (
	"encoding/json"
	"errors"
	"io"
	"math"
	"net/http"
	"strconv"
	"strings"
	"ubiquitous-payment/bank/repository"
)

type Service struct {
	BankRepository *repository.Repository
}

func (service *Service) currencyConversion(from string, to string, amount float32) (float32, error) {
	strAmount := strconv.FormatFloat(float64(amount), 'f', 2, 64)

	url := "https://currency-converter27.p.rapidapi.com/currency/convert"

	payload := strings.NewReader("{\"from\": \"" + from + "\", \"to\": \"" + to + "\",\"amount\": " + strAmount + "}")

	req, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return 0, err
	}

	req.Header.Add("content-type", "application/json")
	req.Header.Add("x-rapidapi-host", "currency-converter27.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "8af238044dmsha877b551370cb40p1b7ee4jsne090c7346514")

	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	responseJson := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&responseJson)
	if err != nil {
		return 0, err
	}

	respPayload, ok := responseJson["payload"].(map[string]interface{})
	if !ok {
		return 0, errors.New("no payload obj in currency conv response")
	}
	status, ok := respPayload["status"].(string)
	if !ok {
		return 0, errors.New("no status string in response payload of currency converter")
	}
	if status != "success" {
		return 0, errors.New("unsuccessful currency conversion")
	}
	amountTo, ok := respPayload["amountTo"].(string)
	if !ok {
		return 0, errors.New("unsuccessful conversion of amountTo to string in currency conversion task")
	}
	retAmount, err := strconv.ParseFloat(amountTo, 32)
	if err != nil {
		return 0, err
	}
	return float32(math.Round(retAmount*100) / 100), nil
}
