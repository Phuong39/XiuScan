package struts2

func setS2046() {
	s := &Struts2{
		Code: "S2-046",
		Desc: "Possible RCE when performing file upload based on Jakarta Multipart parser (similar to S2-045)",
	}
	Modules = append(Modules, s)
}
