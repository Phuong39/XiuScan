package util

import (
	"encoding/binary"
	"encoding/hex"
)

func Int32ToBytes(data uint32) string {
	s := make([]byte, 4)
	binary.BigEndian.PutUint32(s, data)
	return hex.EncodeToString(s)
}

func Int16ToBytes(data uint16) string {
	s := make([]byte, 2)
	binary.BigEndian.PutUint16(s, data)
	return hex.EncodeToString(s)
}
