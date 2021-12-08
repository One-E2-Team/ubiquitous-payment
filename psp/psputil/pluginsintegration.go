package psputil

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"plugin"
	"ubiquitous-payment/psp-plugins/pspdto"
)

type Plugin interface {
	Test() string
	SupportsPlanPayment() bool
	ExecuteTransaction(data pspdto.TransactionDTO) (pspdto.TransactionCreatedDTO, error)
}

var plugins = make(map[string]Plugin, 0)

func loadPlugin(pluginName string) (Plugin, error) {
	var p Plugin = nil
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	exPath := filepath.Dir(ex)
	fmt.Println(exPath)
	plug, err := plugin.Open(fmt.Sprintf("temp/%s.so", pluginName))
	if err != nil {
		return p, err
	}
	pluginSymbol, err := plug.Lookup("Plugin")
	if err != nil {
		return p, err
	}
	p, ok := pluginSymbol.(Plugin)
	if !ok {
		return p, errors.New("invalid plugin type")
	}
	plugins[pluginName] = p
	return p, nil
}

func GetPlugin(pluginName string) (Plugin, error) {
	var err error = nil
	p, ok := plugins[pluginName]
	if ok != true {
		p, err = loadPlugin(pluginName)
	}
	return p, err
}
