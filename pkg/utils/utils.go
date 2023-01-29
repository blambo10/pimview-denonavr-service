package utils

import "github.com/go-resty/resty/v2"

func GetRestyClient() *resty.Client {
	client := resty.New()

	client.EnableTrace()

	return client
}
