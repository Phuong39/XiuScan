package struts2

func setS2009() {
	s := &Struts2{
		Code: "S2-009",
		Desc: "ParameterInterceptor vulnerability allows remote command execution",
	}
	Modules = append(Modules, s)
}
