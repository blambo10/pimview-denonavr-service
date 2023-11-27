package utils

import (
	"github.com/go-resty/resty/v2"
	"strconv"
)

func GetRestyClient() *resty.Client {
	client := resty.New()

	client.EnableTrace()

	return client
}

func GetNumber(s string) (float64, error) {
	number, err := strconv.ParseFloat(s, 64)
	return number, err
}
