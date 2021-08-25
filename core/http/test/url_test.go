package test

import (
	"fmt"
	"github.com/EmYiQing/XiuScan/core/http"
	"testing"
)

func TestUrl(t *testing.T) {
	//urlObj := http.NewUrl("http://www.xxx.com/")
	//urlObj := http.NewUrl("https://www.xxx.com:8080/")
	//urlObj := http.NewUrl("https://www.xxx.com:8080/admin/index.php")
	urlObj := http.NewUrl("https://www.xxx.com:8080/admin/index.php?id=1&page=2")
	fmt.Println(urlObj.Url)
	fmt.Println(urlObj.Protocol)
	fmt.Println(urlObj.Host)
	fmt.Println(urlObj.Port)
	fmt.Println(urlObj.Path)
	fmt.Println(urlObj.Params)
}
