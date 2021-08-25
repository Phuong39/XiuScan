package struts2

func setS2001() {
	s := &Struts2{
		Code: "S2-001",
		Desc: "Remote code exploit on form validation error",
	}
	Modules = append(Modules, s)
}
