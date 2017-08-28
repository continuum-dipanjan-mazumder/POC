package restclient

import (
	"sync"
)

type restClientFactory struct {
}

var instance *restClientFactory
var once sync.Once

func GetClientFactoryInstance() *restClientFactory {
	once.Do(func() {
		instance = &restClientFactory{}
	})
	return instance
}

func (*restClientFactory) GetRestClient(library int) RestClient {
	switch library {
	case 1:
		return &httpClient{}

	case 2:
		return &restyClient{}
	}

	return &httpClient{}

}
