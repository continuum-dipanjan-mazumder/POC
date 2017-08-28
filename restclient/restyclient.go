package restclient

import (
	"fmt"

	"gopkg.in/resty.v0"
)

type restyClient struct{}

func (*restyClient) Get(url string, queryParams map[string]string) (string, error) {
	var resp *resty.Response
	var err error

	//resty.SetDebug(true)

	if queryParams != nil {
		resp, err = resty.R().
			SetQueryParams(queryParams).
			SetHeader("Accept", "application/json").
			Get(url)
		defer resp.RawResponse.Body.Close()

	} else {
		resp, err = resty.R().Get(url)
		defer resp.RawResponse.Body.Close()
	}
	if err != nil {
		return "", err
	}

	//	bodyBytes, err2 := ioutil.ReadAll(bytes.NewReader(resp.Body))
	//	if err2 != nil {
	//		return "", err2
	//	}
	bodyString := string(resp.Body())

	fmt.Println(bodyString)
	return bodyString, nil
}
