package struts2

func setS2013() {
	s := &Struts2{
		Code: "S2-013",
		Desc: "A vulnerability, present in the includeParams " +
			"attribute of the URL and Anchor Tag, allows remote command execution",
	}
	Modules = append(Modules, s)
}
