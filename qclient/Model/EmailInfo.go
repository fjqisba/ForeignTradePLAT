package Model


type EmailInfo struct {
	//邮箱地址
	Email string	`json:"email"`
	//邮箱对应的人名
	PersonName string	`json:"personName,omitempty"`
	//邮箱对应的手机号
	PhoneNumber string	`json:"phoneNumber,omitempty"`
	//职位
	Position string	`json:"position,omitempty"`
	//部门
	Department string	`json:"department,omitempty"`
	//类型
	Type string	`json:"type,omitempty"`
	//领英信息
	Linkedin string	`json:"linkedin,omitempty"`
	//推特信息
	Twitter string	`json:"twitter,omitempty"`
}
