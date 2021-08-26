package tool

import (
	"github.com/thoas/go-funk"
	"math/rand"
	"strings"
	"time"
)

var (
	lowerList  []rune
	upperList  []rune
	letterList []rune
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
	baseLetter := "abcdefghigklmnopqrstuvwxyz"
	lowerList = []rune(baseLetter)
	upperList = []rune(strings.ToUpper(baseLetter))
	letterList = []rune(baseLetter + strings.ToUpper(baseLetter))
}

func GetRandomInt() int {
	return getRandomInt(1000, 2000)
}

func GetRandomIntStr(len int) string {
	return funk.RandomString(len, []rune("1234567890"))
}

func getRandomInt(start, end int) int {
	return funk.RandomInt(start, end)
}

func GetRandomLower(len int) string {
	return funk.RandomString(len, lowerList)
}

func GetRandomUpper(len int) string {
	return funk.RandomString(len, upperList)
}

func GetRandomLetter(len int) string {
	return funk.RandomString(len, letterList)
}
