package input

import "encoding/json"

type Payload struct {
	Name string
	Age  int
}

func GetMessage(name string) string {
	p := new(Payload)
	p.Name = name
	p.Age = 33

	payloadJSON, _ := json.Marshal(p)
	return string(payloadJSON)
}
