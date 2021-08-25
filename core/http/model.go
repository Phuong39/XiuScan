package http

type Request struct {
	Method  string
	Url     *UrlObj
	Body    string
	Headers map[string]string
}

type Response struct {
	Code    int
	Reason  string
	Body    []byte
	Headers map[string]string
}
