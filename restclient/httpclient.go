package restclient

import (
	"POC/parse"
	"log"
	"net/http"
)

type httpClient struct{}

func (*httpClient) Get(url string, queryParams map[string]string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print(err)

	}
	req.Header.Add("Accept", "application/json")
	q := req.URL.Query()
	for key, value := range queryParams {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := parse.GetBodyFromHttpResponse(resp)
	if err != nil {

		return "", err
	}
	return body, nil
}
