package struts2

func setS2016() {
	s := &Struts2{
		Code: "S2-016",
		Desc: "A vulnerability introduced by manipulating parameters " +
			"prefixed with \"action:\"/\"redirect:\"/\"redirectAction:\" allows remote command execution",
	}
	Modules = append(Modules, s)
}
