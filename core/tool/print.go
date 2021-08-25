package tool

import (
	"bytes"
	"strings"
)

func FormatPayload(payload string) string {
	space := "\n               "
	if len(payload) < 60 {
		return payload
	}
	res := bytes.Buffer{}
	for i := 0; i <= len(payload)/60; i++ {
		res.WriteString(space)
		if i == len(payload)/60 {
			res.WriteString(payload[i*60:])
		} else {
			res.WriteString(payload[i*60 : (i+1)*60])
		}
	}
	return strings.TrimRight(res.String(), space)
}
