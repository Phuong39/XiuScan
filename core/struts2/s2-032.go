package struts2

func setS2032() {
	s := &Struts2{
		Code: "S2-032",
		Desc: "Remote Code Execution can be performed via " +
			"method: prefix when Dynamic Method Invocation is enabled",
	}
	Modules = append(Modules, s)
}
