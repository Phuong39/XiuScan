package ysoserial

import "github.com/EmYiQing/XiuScan/core/ysoserial/gadget"

func GetCC1(cmd string) []byte {
	return gadget.GetCommonsCollections1(cmd)
}

func GetCC2(cmd string) []byte {
	return gadget.GetCommonsCollections2(cmd)
}
