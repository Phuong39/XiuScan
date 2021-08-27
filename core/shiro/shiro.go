package shiro

import (
	"github.com/EmYiQing/XiuScan/core/http"
)

func CheckShiroKey(target string) string {
	var key string
	for _, v := range keys {
		data := encodeCBCShiro(v, simplePrincipalCollectionHex)
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
