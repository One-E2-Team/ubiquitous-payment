package bitcoind_rpc

import (
	"github.com/btcsuite/btcd/rpcclient"
	"ubiquitous-payment/util"
)

type BitcoinRPC struct {
	client *rpcclient.Client
}

var rpc = BitcoinRPC{}

func (b *BitcoinRPC) clientInit() (BitcoinRPC, error) {
	var host string = "localhost"
	if util.DockerChecker() {
		host = "host.docker.internal"
	}
	connCfg := &rpcclient.ConnConfig{
		Host:                host + ":18332/wallet/secondarytest",
		User:                "root",
		Pass:                "root",
		DisableConnectOnNew: true,
		HTTPPostMode:        true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:          true, // Bitcoin core does not provide TLS by default
	}
	c, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return *b, err
	}
	b.client = c
	return *b, nil
}

func (b *BitcoinRPC) CloseClient() {
	b.client.Shutdown()
}

func GetClient() (BitcoinRPC, error) {
	if rpc.client != nil {
		return rpc, nil
	} else {
		return rpc.clientInit()
	}
}
