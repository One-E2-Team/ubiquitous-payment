package transactions

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	bitcoind_rpc "ubiquitous-payment/psp-plugins/bitcoin/bitcoind-rpc"
	"ubiquitous-payment/psp-plugins/pspdto"
)

func PrepareTransaction(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	client, err := bitcoind_rpc.GetClient()
	if err != nil {
		return pspdto.TransactionCreatedDTO{}, err
	}
	address, err := client.GetNewAddress("wtf")
	if err != nil {
		return pspdto.TransactionCreatedDTO{}, err
	}
	amount, err := convertValueToBTC(data.Amount, data.Currency)
	if err != nil {
		return pspdto.TransactionCreatedDTO{}, err
	}
	var ret = pspdto.TransactionCreatedDTO{
		TransactionId: address,
		RedirectUrl:   "bitcoin:" + strings.ToUpper(address) + "?amount=" + amount,
	}
	go sendFundsToMerchantWhenReceived(data, ret)
	return ret, nil
}

func convertValueToBTC(value string, currency string) (string, error) {
	url := "https://bitcoinaverage-global-bitcoin-index-v1.p.rapidapi.com/indices/global/ticker/BTC" + currency

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("x-rapidapi-host", "bitcoinaverage-global-bitcoin-index-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "8af238044dmsha877b551370cb40p1b7ee4jsne090c7346514")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	responseJson := make(map[string]interface{})
	err = json.NewDecoder(res.Body).Decode(&responseJson)
	if err != nil {
		return "", err
	}
	averages, ok := responseJson["averages"].(map[string]interface{})
	if !ok {
		return "", errors.New("could not convert currency averages in BTC to " + currency)
	}
	dayAverage, ok := averages["day"].(float64)
	if !ok {
		return "", errors.New("could not convert day currency average (float64) for BTC to " + currency)
	}
	s, err := strconv.ParseFloat(value, 64)
	if err != nil {
		return "", errors.New("could not convert given value - function parameter - to float64")
	}
	return fmt.Sprintf("%.8f", s/dayAverage), nil
}

func sendFundsToMerchantWhenReceived(data pspdto.TransactionDTO, preparedData pspdto.TransactionCreatedDTO) {
	var receivedAmount float64 = 0
	i := 0
	b, err := bitcoind_rpc.GetClient()
	if err != nil {
		// TODO error
	}
	amount, err := strconv.ParseFloat(data.Amount, 64)
	if err != nil {
		// TODO error
	}
	for {
		time.Sleep(5 * time.Second)
		i++
		receivedAmount, err = b.GetReceivedByAddress(preparedData.TransactionId, 6)
		if err != nil {
			// TODO error
		}
		if receivedAmount >= amount {
			break
		}
		if i >= 60*10/5 {
			//TODO fail
		}
	}
	// TODO success
	sendFundsToMerchant(data)
}

func sendFundsToMerchant(data pspdto.TransactionDTO) {
	// TODO
}
