package http

import (
	"strconv"
	"strings"
)

// UrlObj 对URL的封装
// Example: https://www.xxx.com:8080/admin/index.php?id=1&page=2
type UrlObj struct {
	// https://www.xxx.com:8080/admin/index.php?id=1&page=2
	Url string
	// https
	Protocol string
	// www.xxx.com
	Host string
	// 8080
	Port int
	// /admin/index.php
	Path string
	// {id:1,page:2}
	Params map[string]string
}

func NewUrl(url string) *UrlObj {
	instance := &UrlObj{Url: url}
	protocol := strings.Split(url, "://")[0]
	instance.Protocol = protocol
	temp := strings.Split(strings.Split(url, "://")[1], "/")[0]
	if strings.Contains(temp, ":") {
		split := strings.Split(temp, ":")
		instance.Host = split[0]
		instance.Port, _ = strconv.Atoi(split[1])
	} else {
		instance.Host = temp
		instance.Port = 80
	}
	path := strings.Split(strings.Split(url, "?")[0], temp)[1]
	instance.Path = path
	split := strings.Split(url, "?")
	var params string
	if len(split) > 1 {
		params = split[1]
		split = strings.Split(params, "&")
		paramsMap := make(map[string]string)
		for _, v := range split {
			if !strings.Contains(v, "=") {
				continue
			}
			item := strings.Split(v, "=")
			paramsMap[item[0]] = strings.Join(item[1:], "=")
		}
		instance.Params = paramsMap
	} else {
		instance.Params = nil
	}
	return instance
}
