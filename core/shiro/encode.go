package shiro

import (
	"encoding/base64"
)

func EncodeCBCShiro(key string, payload []byte) string {
	decodeKey, _ := base64.StdEncoding.DecodeString(key)
	res := doAESCBCEncrypt(decodeKey, payload)
	return res
}
