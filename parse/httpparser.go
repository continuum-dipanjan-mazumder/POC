package parse

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func GetBodyFromHttpResponse(resp *http.Response) (string, error) {
	bodyString := ""
	if resp.StatusCode == 200 { // OK
		bodyBytes, err2 := ioutil.ReadAll(resp.Body)
		if err2 != nil {
			return "", err2
		}
		bodyString = string(bodyBytes)
	}
	return bodyString, nil
}

func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}
