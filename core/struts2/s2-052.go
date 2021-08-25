package struts2

func setS2052() {
	s := &Struts2{
		Code: "S2-052",
		Desc: "Possible Remote Code Execution attack when using the " +
			"Struts REST plugin with XStream handler to handle XML payloads",
	}
	Modules = append(Modules, s)
}
