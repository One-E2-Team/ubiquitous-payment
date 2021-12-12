package wsutil

import (
	"bytes"
	"fmt"
	"net/http"
	"ubiquitous-payment/util"
)

var pspAccessToken = ""

func SetPspAccessToken(accessToken string) {
	pspAccessToken = accessToken
}

func PSPRequest(method string, path string, data []byte, headers map[string]string) (*http.Response, error) {
	if pspAccessToken == "" {
		initAccessToken()
	}
	client := &http.Client{}
	pspHost, pspPort := util.GetPSPHostAndPort()
	pspPrefix := util.GetPSPProtocol() + "://" + pspHost + ":" + pspPort
	req, err := http.NewRequest(method, pspPrefix+path, bytes.NewBuffer(data))
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+pspAccessToken)
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	return client.Do(req)
}

func initAccessToken() {
	accessToken, err := UtilService.WSService.GetPSPAccessToken()
	if err != nil {
		pspAccessToken = ""
	}
	pspAccessToken = accessToken
}
