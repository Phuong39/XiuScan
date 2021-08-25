package ysoserial

import "github.com/EmYiQing/XiuScan/core/ysoserial/gadget"

func GetCC1(cmd string) []byte {
	return gadget.GetCommonCollections1(cmd)
}
