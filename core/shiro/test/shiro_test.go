package test

import (
	"github.com/EmYiQing/XiuScan/core/shiro"
	"github.com/EmYiQing/XiuScan/log"
	"testing"
)

func TestFindShiro(t *testing.T) {
	key := shiro.CheckShiroKey("http://192.168.222.132:8080/")
	if key != "" {
		log.Info("find key: %s", key)
	}
}
