package http

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"strings"
)

func Get(url string) *Response {
	req := &Request{
		Method: "GET",
		Url:    NewUrl(url),
	}
	return DoRequest(req)
}

func DoRequest(req *Request) *Response {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	request, err := http.NewRequest(req.Method, req.Url.Url, strings.NewReader(req.Body))
	checkErr(err)
	for k, v := range req.Headers {
		request.Header.Set(k, v)
	}
	if _, ok := req.Headers["User-Agent"]; !ok {
		request.Header.Set("User-Agent", "Mozilla/5.0 (Linux; Android 6.0; "+
			"Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) "+
			"Chrome/92.0.4515.131 Mobile Safari/537.36")
	}
	resp, err := client.Do(request)
	checkErr(err)
	response := &Response{
		Code:   resp.StatusCode,
		Reason: resp.Status,
	}
	respHeader := make(map[string]string)
	for k, v := range resp.Header {
		respHeader[k] = strings.Join(v, "; ")
	}
	response.Headers = respHeader
	respBody, err := ioutil.ReadAll(resp.Body)
	checkErr(err)
	response.Body = respBody
	return response
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
