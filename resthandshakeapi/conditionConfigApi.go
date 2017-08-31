package resthandshakeapi

import (
	"POC/config"
	"encoding/json"
	"net/http"
)

//condition config Policy service implementation struct
type ConditionConfigService struct {
	ConditionPolicyMap map[string]config.ConditionPolicy
}

//Get the supported conditions
func (cs *ConditionConfigService) getPolicies() []string {
	i := 0
	supportedConditions := make([]string, len(cs.ConditionPolicyMap))
	for key := range cs.ConditionPolicyMap {
		supportedConditions[i] = key
		i++
	}
	return supportedConditions
}

//Method to handle http requests
func (cs *ConditionConfigService) ServeHTTP(responseWrite http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		supportedConditions := cs.getPolicies()
		if len(supportedConditions) != 0 {
			json.NewEncoder(responseWrite).Encode(supportedConditions)
		}
	}
}
