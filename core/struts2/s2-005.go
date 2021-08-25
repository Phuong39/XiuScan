package struts2

func setS2005() {
	s := &Struts2{
		Code: "S2-005",
		Desc: "XWork parameterInterceptors bypass allows remote command execution",
	}
	Modules = append(Modules, s)
}
