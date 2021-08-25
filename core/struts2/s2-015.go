package struts2

func setS2015() {
	s := &Struts2{
		Code: "S2-015",
		Desc: "A vulnerability introduced by wildcard matching " +
			"mechanism or double evaluation of OGNL Expression allows remote command execution",
	}
	Modules = append(Modules, s)
}
