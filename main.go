package main

import (
	"github.com/EmYiQing/XiuScan/log"
)

const (
	AUTHOR  = "4ra1n,XiuJun"
	VERSION = "0.0.1"
)

func main() {
	log.PrintLogo(VERSION, AUTHOR)
	log.Info("start xiuscan success")
	log.Warn("no input")
	log.Error("exit")
}
