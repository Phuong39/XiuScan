package shiro

import (
	"encoding/base64"
	"encoding/hex"
)

func encodeCBCShiro(key string, payload string) string {
	decodeKey, _ := base64.StdEncoding.DecodeString(key)
	data, _ := hex.DecodeString(payload)
	res := doAESCBCEncrypt(decodeKey, data)
	return res
}
