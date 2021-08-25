package struts2

func setS2048() {
	s := &Struts2{
		Code: "S2-048",
		Desc: "Possible RCE in the Struts Showcase app in the Struts 1 plugin example in Struts 2.3.x series",
	}
	Modules = append(Modules, s)
}
