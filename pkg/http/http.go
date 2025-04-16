package http

import (
	"github.com/go-resty/resty/v2"
)

func Post(url string, body map[string]string, headers map[string]string) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetBody(body).SetHeaders(headers).Post(url)
}

func Get(url string, query map[string]string, headers map[string]string) (*resty.Response, error) {
	client := resty.New()
	return client.R().SetQueryParams(query).SetHeaders(headers).Get(url)
}
