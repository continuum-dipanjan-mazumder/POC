package main

import (
	"POC/config"
	handshake "POC/resthandshakeapi"
	"net/http"
)

func main() {
	//Starting producer
	StartProducer()

	go StartConsumer()

	cs := &handshake.ConditionConfigService{
		ConditionPolicyMap: map[string]config.ConditionPolicy{
			"cond2005": config.ConditionPolicy{
				PolicyId: "cond2005",
			},
		},
	}

	http.ListenAndServe(":8088", cs)
}
