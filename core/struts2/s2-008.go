package struts2

func setS2008() {
	s := &Struts2{
		Code: "S2-008",
		Desc: "Multiple critical vulnerabilities in Struts2",
	}
	Modules = append(Modules, s)
}
