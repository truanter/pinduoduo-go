package sdk

import (
	"net/http"

	"github.com/truanter/pinduoduo-go/pkg/config"

	"github.com/truanter/pinduoduo-go/pkg/api"
)

type SdkOption struct {
	IsDebug       bool
	DebugFilePath string
}

type Client struct {
	http.RoundTripper
	Option         *SdkOption
	API            *api.API
	ClientId       string
	Secret         string
	PublicParams   map[string]interface{}
	AccessToken    string
	middlewareList []Middleware
}

func NewSdkClient(option *SdkOption) *Client {
	apiClient := api.NewAPIClient()
	sdkClient := &Client{
		Option:       option,
		RoundTripper: http.DefaultTransport,
		API:          apiClient,
		PublicParams: config.DefaultParams,
	}
	sdkClient.API.HttpClient.Transport = sdkClient
	sdkClient.middlewareList = []Middleware{
		&AuthMiddleware{sdkClient},
		&LogMiddleware{sdkClient},
	}
	return sdkClient
}

func (c *Client) UseProduction() {
	c.Option.IsDebug = false
}

func (c *Client) UseDebug() {
	c.Option.IsDebug = true
}

func (c *Client) IsDebug() bool {
	return c.Option.IsDebug
}

func (c *Client) SetDebugFile(filePath string) {
	c.Option.DebugFilePath = filePath
}

func (c *Client) GetDebugFile() string {
	return c.Option.DebugFilePath
}

func (c *Client) SetClientIdAndSecret(clientId, secret string) {
	c.ClientId = clientId
	c.Secret = secret
}

func (c *Client) SetAccessToken(accessToken string) {
	c.AccessToken = accessToken
}

func (c *Client) genMiddlewareFunc(mid Middleware, beforeFunc func(*http.Request) (*http.Response, error)) func(*http.Request) (*http.Response, error) {
	return func(r *http.Request) (*http.Response, error) {
		return mid.Handle(r, beforeFunc)
	}
}

func (c *Client) RoundTrip(req *http.Request) (rsp *http.Response, err error) {
	beforeFunc := func(r *http.Request) (*http.Response, error) {
		return c.RoundTripper.RoundTrip(r)
	}
	middlewareCnt := len(c.middlewareList)
	for i := middlewareCnt - 1; i >= 0; i-- {
		beforeFunc = c.genMiddlewareFunc(c.middlewareList[i], beforeFunc)
	}
	return beforeFunc(req)
}
