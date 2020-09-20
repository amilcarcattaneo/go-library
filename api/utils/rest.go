package utils

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	resty "github.com/go-resty/resty"
)

var (
	defaultTimeout       = time.Duration(2 * time.Second)
	defaultRetry         = 3
	defaultRetryWaitTime = time.Duration(300 * time.Millisecond)
)

func MakeGETRequest(url string, headers map[string]string) (int, []byte, error) {
	client := getClient()

	resp, err := client.R().SetHeaders(headers).Get(url)
	if err != nil {
		return 0, nil, err
	}

	if err = ReturnErrorFromStatusCode(resp.StatusCode(), resp.Body()); err != nil {
		return resp.StatusCode(), nil, err
	}

	return resp.StatusCode(), resp.Body(), nil
}

func getClient() *resty.Client {

	return resty.New().AddRetryCondition(
		func(r *resty.Response, err error) bool {
			return r.StatusCode() != http.StatusTooManyRequests
		},
	).SetRetryCount(defaultRetry).SetRetryWaitTime(defaultRetryWaitTime).SetTimeout(defaultTimeout)
}

func ReturnErrorFromStatusCode(statusCode int, body []byte) error {
	if statusCode >= 400 && statusCode <= 599 {
		return errors.New(fmt.Sprintf("[status_code:%d][response:%s]", statusCode, string(body)))
	}
	return nil
}
