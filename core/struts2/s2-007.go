package struts2

func setS2007() {
	s := &Struts2{
		Code: "S2-007",
		Desc: "User input is evaluated as an OGNL expression when there's a conversion error",
	}
	Modules = append(Modules, s)
}
