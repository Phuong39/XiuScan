package struts2

func setS2012() {
	s := &Struts2{
		Code: "S2-012",
		Desc: "Showcase app vulnerability allows remote command execution",
	}
	Modules = append(Modules, s)
}
