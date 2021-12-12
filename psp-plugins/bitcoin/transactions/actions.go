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
	amount, err := convertValueToBTC(data.Amount, data.Currency)
	if err != nil {
		return pspdto.TransactionCreatedDTO{}, err
	}
	address, err := client.GetNewAddress(amount)
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
		fmt.Println(err)
		// TODO error
	}
	label, err := b.GetLabelForAddress(preparedData.TransactionId)
	if err != nil {
		fmt.Println(err)
		// TODO error
	}
	amount, err := strconv.ParseFloat(label, 64)
	if err != nil {
		fmt.Println(err)
		// TODO error
	}
	for {
		time.Sleep(5 * time.Second)
		i++
		receivedAmount, err = b.GetReceivedByAddress(preparedData.TransactionId, 1)
		if err != nil {
			fmt.Println(err)
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
	fmt.Println("ZAPOCINJE PLACANJE")
	sendFundsToMerchant(data, receivedAmount)
}

func sendFundsToMerchant(data pspdto.TransactionDTO, amount float64) {
	b, err := bitcoind_rpc.GetClient()
	if err != nil {
		panic(err)
	}
	err = b.SendAmountToAddressAndSubtractFees(data.PayeeId, fmt.Sprintf("%.8f", amount))
	if err != nil {
		panic(err)
	}
}

func CaptureTransactionSuccess(id string) (bool, error) {
	b, err := bitcoind_rpc.GetClient()
	if err != nil {
		return false, err
	}
	label, err := b.GetLabelForAddress(id)
	if err != nil {
		return false, err
	}
	amount, err := strconv.ParseFloat(label, 64)
	if err != nil {
		return false, err
	}
	receivedAmount, err := b.GetReceivedByAddress(id, 1)
	if err != nil {
		return false, err
	}
	if amount <= receivedAmount {
		return true, nil
	} else {
		return false, nil
	}
}
