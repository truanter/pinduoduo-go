package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	error2 "github.com/truanter/pinduoduo-go/error"

	"github.com/truanter/pinduoduo-go/pkg/config"
)

type api interface {
	GetMethod() string
}

type API struct {
	HttpClient *http.Client
	common     service

	DDKGoodsSearch               *DDKGoodsSearch
	DDKGoodsPromotionUrlGenerate *DDKGoodsPromotionUrlGenerate
	DDKResourceUrlGen            *DDKResourceUrlGen
	DDKGoodsRecommendGet         *DDKGoodsRecommendGet
	DDKGoodsDetail               *DDKGoodsDetail
	DDKRpPromUrlGenerate         *DDKRpPromUrlGenerate
	DDKMemberAuthorityQuery      *DDKMemberAuthorityQuery
	DDKCmsPromUrlGenerate        *DDKCmsPromUrlGenerate

	GeneralAPI *GeneralAPI
}

func NewAPIClient() *API {
	api := &API{}
	api.HttpClient = &http.Client{}
	api.common.client = api

	api.DDKGoodsSearch = (*DDKGoodsSearch)(&api.common)
	api.DDKGoodsPromotionUrlGenerate = (*DDKGoodsPromotionUrlGenerate)(&api.common)
	api.DDKResourceUrlGen = (*DDKResourceUrlGen)(&api.common)
	api.DDKGoodsRecommendGet = (*DDKGoodsRecommendGet)(&api.common)
	api.DDKGoodsDetail = (*DDKGoodsDetail)(&api.common)
	api.DDKRpPromUrlGenerate = (*DDKRpPromUrlGenerate)(&api.common)
	api.DDKMemberAuthorityQuery = (*DDKMemberAuthorityQuery)(&api.common)
	api.DDKCmsPromUrlGenerate = (*DDKCmsPromUrlGenerate)(&api.common)

	api.GeneralAPI = (*GeneralAPI)(&api.common)

	return api
}

type service struct {
	client *API
	Method string
}

func (a *API) call(req *http.Request) (*http.Response, error) {
	return a.HttpClient.Do(req)
}

func (a *API) prepareRequest(method string, requestData interface{}) (r *http.Request, err error) {
	url := config.GetHost()

	body := &bytes.Buffer{}
	if err = json.NewEncoder(body).Encode(requestData); err != nil {
		return
	}

	headers := http.Header{
		"Content-Type": []string{"application/json"},
	}
	if r, err = http.NewRequest(http.MethodPost, url, body); err != nil {
		return
	}
	r.Header = headers
	q := r.URL.Query()
	q.Set("type", method)
	r.URL.RawQuery = q.Encode()
	return
}

func (a *API) do(api api, reqData, respData interface{}) error {
	method := api.GetMethod()
	req, err := a.prepareRequest(method, reqData)
	if err != nil {
		return error2.NewRuntimeError(err.Error())
	}
	resp, err := a.call(req)
	if err != nil {
		return error2.NewRuntimeError(fmt.Sprintf("http request error: %s", err.Error()))
	}
	defer func() {
		if resp != nil {
			_ = resp.Body.Close()
		}
	}()

	if resp.StatusCode == 200 {
		body, _ := ioutil.ReadAll(resp.Body)
		decoder := json.NewDecoder(bytes.NewBuffer(body))
		decoder.UseNumber()
		if err := decoder.Decode(respData); err != nil {
			return error2.NewRuntimeError(fmt.Sprintf("parse response data error: %s", err.Error()))
		}
		return nil
	} else {
		return error2.NewHttpRequestError(resp.StatusCode)
	}
}
