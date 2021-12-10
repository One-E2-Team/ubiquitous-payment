package main

import (
	"fmt"
	"log"
	bitcoind_rpc "ubiquitous-payment/psp-plugins/bitcoin/bitcoind-rpc"
	"ubiquitous-payment/psp-plugins/pspdto"
)

type plugin struct {
}

func (p plugin) Test() string {
	fmt.Println("Bit-in bit-out wasaaaaaaaaaa")
	return "bitups"
}

func (p plugin) SupportsPlanPayment() bool {
	return true
}

func (p plugin) ExecuteTransaction(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error) {
	return pspdto.TransactionCreatedDTO{}, nil
}

var Plugin plugin

func main() {
	client, err := bitcoind_rpc.GetClient()
	if err != nil {
		fmt.Println(err)
	}
	defer bitcoind_rpc.CloseClient()

	address, err := client.GetNewAddress("wtf")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(address)
}
