package transactions

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
	"ubiquitous-payment/psp-plugins/bank/dto"
	"ubiquitous-payment/psp-plugins/pspdto"
)

func PrepareTransaction(data pspdto.TransactionDTO, context *map[string]string, pluginName string) (pspdto.TransactionCreatedDTO, error) {
	amount, err := strconv.ParseFloat(data.Amount, 32)
	if err != nil {
		return pspdto.TransactionCreatedDTO{}, err
	}
	req := dto.PspRequestDTO{
		MerchantId:        data.PayeeId,
		MerchantPassword:  data.PayeeSecret,
		Amount:            float32(amount),
		Currency:          data.Currency,
		MerchantOrderID:   data.OrderId,
		MerchantTimestamp: time.Now(),
		SuccessURL:        data.SuccessUrl,
		FailURL:           data.FailUrl,
		ErrorURL:          data.ErrorUrl,
		Method:            pluginName,
	}
	targetUrl, err := determineBankEndpoint(data.PayeeId, context)
	if err != nil {
		return pspdto.TransactionCreatedDTO{}, err
	}
	var ret dto.PspResponseDTO
	err = CallBankAPI(http.MethodPost, targetUrl, req, &ret)
	if err != nil {
		return pspdto.TransactionCreatedDTO{}, err
	}
	checkBaseUrl := ret.PaymentCheckUrl[:len(ret.PaymentCheckUrl)-4] // remove '{id}' at the end
	return pspdto.TransactionCreatedDTO{TransactionId: checkBaseUrl + ret.PaymentId, RedirectUrl: ret.PaymentUrl}, nil
}

func determineBankEndpoint(id string, context *map[string]string) (string, error) {
	for k, v := range *context {
		if id[:len(k)] == k {
			return v, nil
		}
	}
	return "", errors.New("pan prefix could not be found in context")
}

func CheckPaymentStatusSuccess(id string) (bool, error) {
	var ret dto.PaymentResponseDTO
	fmt.Println("prije calla za provjeru bla bla")
	err := CallBankAPI(http.MethodGet, id, nil, &ret)
	fmt.Println("posle calla ya provjeru sa err ", err, "i ret", ret)
	if err != nil {
		return false, err
	}
	if ret.TransactionStatus != dto.FULFILLED {
		return false, nil
	} else {
		return true, nil
	}
}

func CallBankAPI(method string, url string, data interface{}, ret interface{}) error {
	var body io.Reader = nil
	if data != nil {
		payloadBytes, err := json.Marshal(data)
		if err != nil {
			return err
		}
		body = bytes.NewReader(payloadBytes)
	}

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 && resp.StatusCode != 201 {
		//return nil, errors.New("response is not http 200 or http 201")
		fmt.Println(resp.StatusCode)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	if resp.StatusCode == 204 {
		return nil
	}

	err = json.NewDecoder(resp.Body).Decode(ret)
	if err != nil {
		return err
	}
	fmt.Println(ret)
	return nil
}
