package http

import "strings"

func ContainsCookie(resp *Response, cookie string) bool {
	var flag bool
	for _, c := range resp.Cookies {
		if strings.Contains(c.Raw, cookie) {
			flag = true
		}
	}
	return flag
}
