package Model

import "time"

type CustomerInformation struct {
	//公司名称
	CompanyName    string    `json:"company_name"`
	//成立年份
	YearFounded    int    `json:"year_founded"`
	//员工数量
	EmployeesNum   int       `json:"employees_num"`
	//年营业额(美元)
	AnnualRevenue  int		`json:"annual_revenue"`
	//联系电话
	PhoneNumber    string		`json:"phone_number"`
	//客户邮箱
	Email          string    `json:"email"`
	//联系人名字
	KeyContact     string    `json:"key_contact"`
	//Hunter大数据
	HunterData     string    `json:"hunter_data"`
	//客户网址
	CompanyDomain  string    `json:"company_domain"`
	//客户类型
	CustomerType   string    `json:"customer_type"`
	//客户等级
	CustomerRating string    `json:"customer_rating"`
	//最后跟进日期
	FollowupDate   time.Time `json:"followup_date"`
	//备注
	Remarks        string    `json:"remarks"`
	//客户介绍
	Description    string    `json:"description"`
	//Facebook主页
	FacebookPage	string	`json:"facebook_page"`
	//领英主页
	LinkedInPage	string	`json:"linkedin_page"`
	//推特账号
	TwitterHandle	string	`json:"twitter_handle"`
	//行业
	Industry	string		`json:"industry"`
	//国家
	Country		string		`json:"country"`
	//城市
	City		string		`json:"city"`
	//详细地址
	StreetAddress	string	`json:"street_address"`
	//邮编
	PostalCode	string		`json:"postal_code"`
	//跟进人
	Coordinator	string		`json:"coordinator"`
	//导入日期
	CreatTime	time.Time		`json:"creat_time"`
}