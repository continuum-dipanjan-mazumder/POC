package restclient

type RestClient interface {
	Get(url string, queryParams map[string]string) (string, error)
}
