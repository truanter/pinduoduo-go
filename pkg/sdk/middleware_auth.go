package sdk

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"time"
)

type AuthMiddleware struct {
	client *Client
}

func (a *AuthMiddleware) Handle(req *http.Request, next func(*http.Request) (*http.Response, error)) (*http.Response, error) {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	_ = req.ParseForm()

	allData := make(map[string]string)
	if a.client.AccessToken != "" {
		allData["access_token"] = a.client.AccessToken
	}
	allData["timestamp"] = timestamp
	allData["client_id"] = a.client.ClientId

	q := req.URL.Query()

	for k, v := range a.client.PublicParams {
		allData[k] = v.(string)
	}

	for k, v := range q {
		if len(v) > 0 {
			allData[k] = v[0]
		}
	}

	jsonData := make(map[string]interface{})
	b, _ := ioutil.ReadAll(req.Body)
	decoder := json.NewDecoder(bytes.NewBuffer(b))
	decoder.UseNumber()
	_ = decoder.Decode(&jsonData)

	for k, v := range jsonData {
		data := ""
		switch v.(type) {
		case []interface{}, reflect.StructField:
			b, _ := json.Marshal(v)
			s := string(b)
			data = s
			break
		default:
			data = fmt.Sprintf("%v", v)
			break
		}
		allData[k] = data
	}

	multipartFormData := make(map[string]interface{})
	if req.MultipartForm != nil {
		for k, v := range req.MultipartForm.Value {
			if len(v) > 0 {
				multipartFormData[k] = v[0]
				allData[k] = v[0]
			}
		}
	}

	formData := make(map[string]interface{})
	if len(req.PostForm) > 0 {
		for k, v := range req.PostForm {
			if len(v) > 0 {
				formData[k] = v[0]
				allData[k] = v[0]
			}
		}
	}

	keys := make([]string, 0, len(allData))
	for k, _ := range allData {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	strKey := a.client.Secret
	for _, v := range keys {
		strKey += v + allData[v]
	}
	strKey += a.client.Secret
	h := md5.New()
	h.Write([]byte(strKey))
	sign := hex.EncodeToString(h.Sum(nil))
	sign = strings.ToUpper(sign)
	allData["sign"] = sign

	data := url.Values{}
	for k, v := range allData {
		data.Set(k, fmt.Sprintf("%v", v))
	}

	req.URL.RawQuery = ""
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Content-Length", strconv.Itoa(len(data.Encode())))
	req.ContentLength = int64(len(data.Encode()))

	req.Body = ioutil.NopCloser(strings.NewReader(data.Encode()))

	resp, err := next(req)
	return resp, err
}
