package bitcoind_rpc

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
	return info.Labels[0], nil
}
