package utils

import (
	"net"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
)

func ResponseWasConflict(resp autorest.Response) bool {
	return responseWasStatusCode(resp, http.StatusConflict)
}

func ResponseWasNotFound(resp autorest.Response) bool {
	return responseWasStatusCode(resp, http.StatusNotFound)
}

func ResponseErrorIsRetryable(err error) bool {
	if arerr, ok := err.(autorest.DetailedError); ok {
		err = arerr.Original
	}

	switch e := err.(type) {
	case net.Error:
		if e.Temporary() || e.Timeout() {
			return true
		}
	}

	return false
}

func responseWasStatusCode(resp autorest.Response, statusCode int) bool {
	if r := resp.Response; r != nil {
		if r.StatusCode == statusCode {
			return true
		}
	}

	return false
}
