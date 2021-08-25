package struts2

func setS2053() {
	s := &Struts2{
		Code: "S2-053",
		Desc: "A possible Remote Code Execution attack when using " +
			"an unintentional expression in Freemarker tag instead of string literals",
	}
	Modules = append(Modules, s)
}
