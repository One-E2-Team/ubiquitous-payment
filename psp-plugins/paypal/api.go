package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var expiresIn = 0
var accessToken string
var lastRefreshTime time.Time = time.Now()

func extractTime(s string) error {
	timestampZulu := strings.Split(s, "Z")
	zuluParts := strings.Split(timestampZulu[0], "T")
	dateParts := strings.Split(zuluParts[0], "-")
	timeParts := strings.Split(zuluParts[1], ":")
	year, err := strconv.Atoi(dateParts[0])
	if err != nil {
		return err
	}
	month, err := strconv.Atoi(dateParts[1])
	if err != nil {
		return err
	}
	day, err := strconv.Atoi(dateParts[2])
	if err != nil {
		return err
	}
	hour, err := strconv.Atoi(timeParts[0])
	if err != nil {
		return err
	}
	minutes, err := strconv.Atoi(timeParts[1])
	if err != nil {
		return err
	}
	seconds, err := strconv.Atoi(timeParts[2])
	if err != nil {
		return err
	}
	lastRefreshTime = time.Date(year, time.Month(month), day, hour, minutes, seconds, 0, time.UTC)
	return nil
}

func hasExpired() bool {
	return time.Now().After(lastRefreshTime.Add(time.Duration(expiresIn * int(time.Second))))
}

func getNewAccessToken() error {
	params := url.Values{}
	params.Add("grant_type", `client_credentials`)
	body := strings.NewReader(params.Encode())

	id, ok := os.LookupEnv("PAYPAL_CLIENT_ID")
	if !ok {
		return errors.New("no paypal client id in env")
	}
	secret, ok := os.LookupEnv("PAYPAL_CLIENT_SECRET")
	if !ok {
		return errors.New("no paypal secret in env")
	}

	req, err := http.NewRequest("POST", "https://api-m.sandbox.paypal.com/v1/oauth2/token", body)
	if err != nil {
		return err
	}
	req.SetBasicAuth(id, secret)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en_US")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	data := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return err
	}

	token, ok := data["access_token"].(string)
	if !ok {
		return errors.New("access token conversion error")
	}
	timestamp, ok := data["nonce"].(string)
	if !ok {
		return errors.New("last refresh time conversion error")
	}

	err = extractTime(timestamp)
	if err != nil {
		return err
	}
	accessToken = token

	return nil
}

func getAccessToken() (string, error) {
	if hasExpired() {
		err := getNewAccessToken()
		if err != nil {
			return "", err
		}
	}
	return accessToken, nil
}

type Payload struct {
	Intent        string          `json:"intent"`
	PurchaseUnits []PurchaseUnits `json:"purchase_units"`
}
type Amount struct {
	CurrencyCode string `json:"currency_code"`
	Value        string `json:"value"`
}
type PurchaseUnits struct {
	Amount Amount `json:"amount"`
}

func CallPayPalAPI(method string, url string, data interface{}) (map[string]interface{}, error) {
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	token, err := getAccessToken()
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	responseJson := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&responseJson)
	if err != nil {
		return nil, err
	}

	return responseJson, nil
}

const OrdersApiUrl = "https://api-m.sandbox.paypal.com/v2/checkout/orders"
