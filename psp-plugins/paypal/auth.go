package main

import (
	"encoding/json"
	"errors"
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
	defer resp.Body.Close()

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

func GetAccessToken() (string, error) {
	if hasExpired() {
		err := getNewAccessToken()
		if err != nil {
			return "", err
		}
	}
	return accessToken, nil
}
