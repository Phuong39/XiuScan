package util

import "strings"

// CheckUrl 获得结尾不是/的Url
func CheckUrl(url string) string {
	var result string
	url = strings.TrimSpace(url)
	if !strings.HasPrefix(url, "http") {
		return result
	}
	for strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}
	result = url
	return result
}
