package proxy

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/sirupsen/logrus"
)

type RequestPayloadStruct struct {
	ProxyCondition string `json:"proxy_condition"`
}

func ServeReverseProxy(target string, resp http.ResponseWriter, req *http.Request) {
	url, _ := url.Parse(target)

	proxy := httputil.NewSingleHostReverseProxy(url)

	// Update the headers to allow for SSL redirection
	req.URL.Host = url.Host
	req.URL.Scheme = url.Scheme
	req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
	req.Host = url.Host

	proxy.ServeHTTP(resp, req)
}

func RequestBodyDecoder(req *http.Request) *json.Decoder {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		logrus.Printf("Error reading the request body: %v\n", err)
		panic(err)
	}

	req.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	return json.NewDecoder(ioutil.NopCloser(bytes.NewBuffer(body)))
}

func ParseRequestBody(req *http.Request) RequestPayloadStruct {
	var payload RequestPayloadStruct

	decoder := RequestBodyDecoder(req)
	err := decoder.Decode(&payload)
	if err != nil {
		logrus.Printf("Error decoding the json: %v\n", err)
		panic(err)
	}
	return payload
}

func HandleRequestAndRedirect(resp http.ResponseWriter, req *http.Request) {
	payload := ParseRequestBody(req)
	url := GetProxyURL(payload.ProxyCondition)
	LogRequestPayload(payload, url)
	ServeReverseProxy(url, resp, req)
}
