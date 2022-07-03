package DataBase

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"server/Model"
	"strings"
	"time"
)

var(
	Instance DataBase
)

const sqlCreateTable_UserROLE = `CREATE TABLE TABLE_USERROLE (
  username VARCHAR (64) PRIMARY KEY NOT NULL,
  role int	COMMENT '1表示管理员,2表示普通用户'
);`

const sqlCreateTable_CustomerInfo = `CREATE TABLE TABLE_CUSTOMER (
	company_domain		VARCHAR (64) PRIMARY KEY,
	company_name		VARCHAR (64),
	year_founded		INTEGER,
	employees_num		INTEGER,
	annual_revenue		INTEGER,
	phone_number		VARCHAR (32),
	email				VARCHAR (32),
	key_contact			VARCHAR (32),
	hunter_data			TEXT,
	customer_type		VARCHAR (32),
	customer_rating		VARCHAR (32),
	followup_date		DATE,
	remarks				TEXT,
	description			TEXT,
	facebook_page		VARCHAR (64),
	linkedin_page		VARCHAR (64),
	twitter_handle		VARCHAR (64),
	industry			VARCHAR (64),
	country				VARCHAR (64),
	city				VARCHAR (64),
	street_address		VARCHAR (64),
	postal_code			VARCHAR (64),
	coordinator			VARCHAR (64),
	creat_time			DATE
);`

type DataBase struct {
	sql sqlx.DB
}

func init()  {
	err := Instance.initTable()
	if err != nil{
		log.Panicln(err)
	}
}

func (this *DataBase)initTable()error  {
	var err error
	this.sql.DB, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/Trade")
	if err != nil{
		return err
	}

	err = this.sql.DB.Ping()
	if err != nil{
		return err
	}

	//创建表TABLE_USER
	_,err = this.sql.Exec(sqlCreateTable_UserROLE)
	if err != nil{
		if strings.Contains(err.Error(),"already exists") == false{
			return err
		}
	}

	//创建表用户信息
	_,err = this.sql.Exec(sqlCreateTable_CustomerInfo)
	if err != nil{
		if strings.Contains(err.Error(),"already exists") == false{
			return err
		}
	}
	return nil
}

func (this *DataBase)DeleteUser(userName string)error  {
	_,err := this.sql.Exec("delete from TABLE_USERROLE where username=? and role = 2",userName)
	if err != nil{
		return err
	}
	return nil
}


func (this *DataBase)GetDataCount(userName string)(int,error)  {
	var rowsCount int
	err := this.sql.Get(&rowsCount,"select count(*) from TABLE_CUSTOMER where coordinator=?",userName)
	if err != nil{
		return 0,err
	}
	return rowsCount,nil
}

//获取数据总数

func (this *DataBase)GetAllDataCount()(int,error)  {
	var rowsCount int
	err := this.sql.Get(&rowsCount,"select count(*) from TABLE_CUSTOMER")
	if err != nil{
		return 0,err
	}
	return rowsCount,nil
}

func (this *DataBase)AddUser(userName string)error  {
	_,err := this.sql.Exec("insert into TABLE_USERROLE(username,role) VALUES(?,?)",userName,Model.USER_USER)
	if err != nil{
		return err
	}
	return nil
}

func (this *DataBase)AssignTask(assignInfo Model.AssignTaskReq)(error)  {

	if len(assignInfo.Domain) == 0{
		return nil
	}
	const sql_assignTask = `update TABLE_CUSTOMER set coordinator=? where company_domain=?`
	tx,err := this.sql.Begin()
	if err != nil{
		return err
	}
	stmt,err := tx.Prepare(sql_assignTask)
	if err != nil{
		return err
	}
	defer stmt.Close()
	for _,eDomain := range assignInfo.Domain{
		_,err = stmt.Exec(assignInfo.Target,eDomain)
		if err != nil{
			log.Println("[AssignTask]Exec:",err)
		}
	}
	err = tx.Commit()
	if err != nil{
		log.Println("[AssignTask]Commit:",err)
	}
	return nil
}


func (this *DataBase)GetUserList()([]string,error)  {
	var retUserList []string
	err := this.sql.Select(&retUserList,"select username from TABLE_USERROLE where role = 2")
	if err != nil{
		return nil,err
	}
	return retUserList,nil
}

func (this *DataBase)GetUserLevel(userName string)(Model.UserLevel,error)  {
	var retRole Model.UserLevel
	err := this.sql.Get(&retRole,"select role from TABLE_USERROLE where username=?",userName)
	if err != nil{
		return Model.USER_INVALID,err
	}
	return retRole,nil
}

func (this *DataBase)UpdateCustomerData(dataList []Model.CustomerInformation)error{
	if len(dataList) == 0{
		return nil
	}
	const sql_updateCustomerData = `update TABLE_CUSTOMER set company_domain=?,company_name=?,year_founded=?,employees_num=?,
annual_revenue=?,phone_number=?,email=?,key_contact=?,hunter_data=?,customer_type=?,customer_rating=?,followup_date=?,
remarks=?,description=?,facebook_page=?,linkedin_page=?,twitter_handle=?,industry=?,country=?,city=?,street_address=?,
postal_code=?,coordinator=?,creat_time=? where company_domain = ?`
	tx,err := this.sql.Begin()
	if err != nil{
		return err
	}
	stmt,err := tx.Prepare(sql_updateCustomerData)
	if err != nil{
		return err
	}
	defer stmt.Close()
	for _,eData := range dataList{
		_,err = stmt.Exec(eData.CompanyDomain,eData.CompanyName,eData.YearFounded,eData.EmployeesNum,eData.AnnualRevenue,
			eData.PhoneNumber,eData.Email,eData.KeyContact,eData.HunterData,eData.CustomerType,eData.CustomerRating,
			eData.FollowupDate,eData.Remarks,eData.Description,eData.FacebookPage,eData.LinkedInPage,eData.TwitterHandle,
			eData.Industry,eData.Country,eData.City,eData.StreetAddress,eData.PostalCode,eData.Coordinator,eData.CreatTime,
		eData.CompanyDomain)
		if err != nil{
			log.Println("[UpdateCustomerData]Exec:",err)
		}
	}
	err = tx.Commit()
	if err != nil{
		log.Println("[UpdateCustomerData]Commit:",err)
	}
	return nil
}

func (this *DataBase)QueryCustomerData(from int,count int,orderBy string,userName string)([]Model.CustomerInformation,error)  {
	const sql_queryAllData = `select company_domain,company_name,
year_founded,employees_num,annual_revenue,phone_number,email,key_contact,hunter_data,customer_type,
customer_rating,followup_date,remarks,description,facebook_page,linkedin_page,twitter_handle,
industry,country,city,street_address,postal_code,coordinator,creat_time from TABLE_CUSTOMER where coordinator=? limit ?,?`
	var retCustomerList []Model.CustomerInformation
	rows,err := this.sql.Query(sql_queryAllData,userName,from,count)
	if err != nil{
		return nil,err
	}
	var tmpFollowupDate string
	var tmpCreateTime string
	for rows.Next(){
		var tmpInfo Model.CustomerInformation
		err = rows.Scan(&tmpInfo.CompanyDomain,&tmpInfo.CompanyName,&tmpInfo.YearFounded,&tmpInfo.EmployeesNum,
			&tmpInfo.AnnualRevenue,&tmpInfo.PhoneNumber,&tmpInfo.Email,&tmpInfo.KeyContact, &tmpInfo.HunterData,
			&tmpInfo.CustomerType,&tmpInfo.CustomerRating,&tmpFollowupDate,&tmpInfo.Remarks, &tmpInfo.Description,
			&tmpInfo.FacebookPage,&tmpInfo.LinkedInPage,&tmpInfo.TwitterHandle,&tmpInfo.Industry, &tmpInfo.Country,
			&tmpInfo.City,&tmpInfo.StreetAddress,&tmpInfo.PostalCode,&tmpInfo.Coordinator,&tmpCreateTime)
		if err != nil{
			continue
		}
		tmpInfo.FollowupDate, _ = time.Parse("2006-01-02",tmpFollowupDate)
		tmpInfo.CreatTime, _ = time.Parse("2006-01-02",tmpCreateTime)
		retCustomerList = append(retCustomerList, tmpInfo)
	}
	return retCustomerList,nil
}

func (this *DataBase)QueryAllCustomerData(from int,count int,orderBy string)([]Model.CustomerInformation,error)  {
	const sql_queryAllData = `select company_domain,company_name,
year_founded,employees_num,annual_revenue,phone_number,email,key_contact,hunter_data,customer_type,
customer_rating,followup_date,remarks,description,facebook_page,linkedin_page,twitter_handle,
industry,country,city,street_address,postal_code,coordinator,creat_time from TABLE_CUSTOMER limit ?,?`
	var retCustomerList []Model.CustomerInformation
	rows,err := this.sql.Query(sql_queryAllData,from,count)
	if err != nil{
		return nil,err
	}
	var tmpFollowupDate string
	var tmpCreateTime string
	for rows.Next(){
		var tmpInfo Model.CustomerInformation
		err = rows.Scan(&tmpInfo.CompanyDomain,&tmpInfo.CompanyName,&tmpInfo.YearFounded,&tmpInfo.EmployeesNum,
			&tmpInfo.AnnualRevenue,&tmpInfo.PhoneNumber,&tmpInfo.Email,&tmpInfo.KeyContact, &tmpInfo.HunterData,
			&tmpInfo.CustomerType,&tmpInfo.CustomerRating,&tmpFollowupDate,&tmpInfo.Remarks, &tmpInfo.Description,
			&tmpInfo.FacebookPage,&tmpInfo.LinkedInPage,&tmpInfo.TwitterHandle,&tmpInfo.Industry, &tmpInfo.Country,
			&tmpInfo.City,&tmpInfo.StreetAddress,&tmpInfo.PostalCode,&tmpInfo.Coordinator,&tmpCreateTime)
		if err != nil{
			continue
		}
		tmpInfo.FollowupDate, _ = time.Parse("2006-01-02",tmpFollowupDate)
		tmpInfo.CreatTime, _ = time.Parse("2006-01-02",tmpCreateTime)
		retCustomerList = append(retCustomerList, tmpInfo)
	}
	return retCustomerList,nil
}

func (this *DataBase)ImportCustomerInfo(customerList []Model.CustomerInformation)error  {

	if len(customerList) == 0{
		return nil
	}

	const sqlInsertCustomerInfo = `INSERT IGNORE INTO TABLE_CUSTOMER(company_domain,company_name,
year_founded,employees_num,annual_revenue,phone_number,email,key_contact,hunter_data,customer_type,
customer_rating,followup_date,remarks,description,facebook_page,linkedin_page,twitter_handle,
industry,country,city,street_address,postal_code,coordinator,creat_time) VALUES(?,?,?,?,?,
?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

	tx,err := this.sql.Begin()
	if err != nil{
		return err
	}
	stat,err := tx.Prepare(sqlInsertCustomerInfo)
	if err != nil{
		return err
	}
	defer stat.Close()

	for _,eCust := range customerList{
		_, err = stat.Exec(eCust.CompanyDomain, eCust.CompanyName,eCust.YearFounded,eCust.EmployeesNum,
			eCust.AnnualRevenue,eCust.PhoneNumber,eCust.Email,eCust.KeyContact,eCust.HunterData,eCust.CustomerType,
			eCust.CustomerRating,eCust.FollowupDate,eCust.Remarks,eCust.Description,eCust.FacebookPage,
			eCust.LinkedInPage,eCust.TwitterHandle,eCust.Industry,eCust.Country,eCust.City,eCust.StreetAddress,
			eCust.PostalCode,eCust.Coordinator,eCust.CreatTime)
		if err != nil{
			log.Println("insertCustomer failed:",err)
			continue
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}
	return nil
}