package ceye

import "fmt"

type Ceye struct {
	Identifier string
	Token      string
	Type       string
	Filter     string
	Url        string
}

func (c *Ceye) NewHttpCeye(identifier, token, filter string) {
	url := fmt.Sprintf("http://api.ceye.io/v1/records?token=%s&type=http&filter=%s", token, filter)
	c.Identifier = identifier
	c.Token = token
	c.Filter = filter
	c.Url = url
}

func (c *Ceye) ChangeFilter(filter string) {
	url := fmt.Sprintf("http://api.ceye.io/v1/records?token=%s&type=http&filter=%s", c.Token, filter)
	c.Filter = filter
	c.Url = url
}
