package struts2

func setS2057() {
	s := &Struts2{
		Code: "S2-057",
		Desc: "Possible Remote Code Execution when alwaysSelectFullNamespace is true",
	}
	Modules = append(Modules, s)
}
