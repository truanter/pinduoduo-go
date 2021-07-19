package sdk

import "net/http"

type Middleware interface {
	Handle(req *http.Request, next func(*http.Request) (*http.Response, error)) (*http.Response, error)
}
