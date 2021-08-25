package struts2

func setS2045() {
	s := &Struts2{
		Code: "S2-045",
		Desc: "Possible Remote Code Execution when performing file upload based on Jakarta Multipart parser",
	}
	Modules = append(Modules, s)
}
