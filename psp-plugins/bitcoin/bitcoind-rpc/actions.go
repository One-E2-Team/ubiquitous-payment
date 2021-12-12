package bitcoind_rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func (b *BitcoinRPC) GetNewAddress(label string) (string, error) {
	address, err := b.client.GetNewAddress(label)
	if err != nil {
		return "", err
	}
	return address.String(), nil
}

func (b *BitcoinRPC) GetReceivedByAddress(address string, minimumConfirmations int) (float64, error) {
	value, err := b.client.GetReceivedByAccountMinConf(address, minimumConfirmations)
	if err != nil {
		return 0, err
	}
	return value.ToBTC(), nil
}

func (b *BitcoinRPC) GetLabelForAddress(address string) (string, error) {
	info, err := b.client.GetAddressInfo(address)
	if err != nil {
		return "", err
	}
	fmt.Println(*info)
	return info.Labels[0], nil
}

func (b *BitcoinRPC) SendAmountToAddressAndSubtractFees(address string, amount string) error {
	type Payload struct {
		Jsonrpc string        `json:"jsonrpc"`
		ID      string        `json:"id"`
		Method  string        `json:"method"`
		Params  []interface{} `json:"params"`
	}

	data := Payload{
		Jsonrpc: "2.0",
		ID:      "0",
		Method:  "sendtoaddress",
		Params:  []interface{}{address, amount, "", "", true, true, 6, "unset", false},
	}
	payloadBytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	body := bytes.NewReader(payloadBytes)

	req, err := http.NewRequest("POST", "host.docker.internal:18332/wallet/secondarytest", body)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Basic cm9vdDpyb290")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	responseJson := make(map[string]interface{})
	err = json.NewDecoder(resp.Body).Decode(&responseJson)
	if err != nil {
		return err
	}
	fmt.Println(responseJson)
	if responseJson["error"] == nil {
		return nil
	} else {
		return errors.New("wallet side error occurred")
	}
}
