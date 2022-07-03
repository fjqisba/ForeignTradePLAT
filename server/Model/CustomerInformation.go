package Model

import "time"

type CustomerInformation struct {
	//公司名称
	CompanyName    string    `json:"company_name" db:"company_name"`
	//成立年份
	YearFounded    int    `json:"year_founded" db:"year_founded"`
	//员工数量
	EmployeesNum   int       `json:"employees_num" db:"employees_num"`
	//年营业额(美元)
	AnnualRevenue  int		`json:"annual_revenue" db:"annual_revenue"`
	//联系电话
	PhoneNumber    string		`json:"phone_number" db:"phone_number"`
	//客户邮箱
	Email          string    `json:"email" db:"email"`
	//联系人名字
	KeyContact     string    `json:"key_contact" db:"key_contact"`
	//Hunter大数据
	HunterData     string    `json:"hunter_data" db:"hunter_data"`
	//客户网址
	CompanyDomain  string    `json:"company_domain" db:"company_domain"`
	//客户类型
	CustomerType   string    `json:"customer_type" db:"customer_type"`
	//客户等级
	CustomerRating string    `json:"customer_rating" db:"customer_rating"`
	//最后跟进日期
	FollowupDate   time.Time `json:"followup_date" db:"followup_date"`
	//备注
	Remarks        string    `json:"remarks" db:"remarks"`
	//客户介绍
	Description    string    `json:"description" db:"description"`
	//Facebook主页
	FacebookPage	string	`json:"facebook_page" db:"facebook_page"`
	//领英主页
	LinkedInPage	string	`json:"linkedin_page" db:"linkedin_page"`
	//推特账号
	TwitterHandle	string	`json:"twitter_handle" db:"twitter_handle"`
	//行业
	Industry	string		`json:"industry" db:"industry"`
	//国家
	Country		string		`json:"country" db:"country"`
	//城市
	City		string		`json:"city" db:"city"`
	//详细地址
	StreetAddress	string	`json:"street_address" db:"street_address"`
	//邮编
	PostalCode	string		`json:"postal_code" db:"postal_code"`
	//跟进人
	Coordinator	string		`json:"coordinator" db:"coordinator"`
	//导入日期
	CreatTime	time.Time		`json:"creat_time" db:"creat_time"`
}

type AssignTaskReq struct {
	Target string	`json:"target"`
	Domain []string	`json:"domain"`
}