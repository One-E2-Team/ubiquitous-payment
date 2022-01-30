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
	CaptureTransaction(id string, plan bool) (bool, error)
	InitContextData(context map[string]string)
}

var plugins = make(map[string]Plugin, 0)
var PluginInterfaceContext PluginContext

type PluginContext interface {
	GetAllBanksKeyValue() (*map[string]string, error)
}

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
	context, err := getContext()
	if err != nil {
		return p, err
	}
	p.InitContextData(*context)
	plugins[pluginName] = p
	return p, nil
}

func getContext() (*map[string]string, error) {
	ret, err := PluginInterfaceContext.GetAllBanksKeyValue()
	if err != nil {
		return nil, err
	}
	return ret, err
}

func GetPlugin(pluginName string) (Plugin, error) {
	var err error = nil
	p, ok := plugins[pluginName]
	if ok != true {
		p, err = loadPlugin(pluginName)
	}
	return p, err
}
