package struts2

type Struts2 struct {
	Code    string
	Desc    string
	Payload string
}

var Modules []*Struts2

func init() {
	setS2001()
	setS2005()
	setS2007()
	setS2008()
	setS2009()
	setS2012()
	setS2013()
	setS2015()
	setS2016()
	setS2032()
	setS2045()
	setS2046()
	setS2048()
	setS2052()
	setS2053()
	setS2057()
}

func GetModules() []*Struts2 {
	return Modules
}
