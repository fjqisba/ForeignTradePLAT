package TradeApp

import (
	"fmt"
	"qclient/Const"
	"qclient/Model"
	"qclient/api"
	"qclient/global"
	"strconv"
	"time"
)

//更新页面
//page是第几页
//count是每页的个数
//orderBy是排序依据

func (this *TradeApp)UpdatePage(page int,count int,orderBy string,userName string)  {
	var retResp Model.CustomerDataRsp
	var err error
	if global.GUserLevel == Model.USER_ADMIN{
		retResp,err = api.GetAllCustomerData(page,count,orderBy)
	}else{
		retResp,err = api.GetCustomerData(page,count,orderBy,userName)
	}
	if err != nil{
		return
	}
	//清空列表
	this.table_CSV.SetRowCount(0)
	this.table_CSV.ClearContents()
	for _,eData := range retResp.Data{
		this.insertCsvItem(&eData)
	}
	//设置页数
	totalPage := retResp.Count / Const.PageNum
	if retResp.Count % Const.PageNum != 0{
		totalPage = totalPage + 1
	}
	global.GCurrentPageIndex = page
	global.GTotalPageCount = totalPage
	this.label_CsvCount.SetText(fmt.Sprintf("共%d页,当前第%d页",totalPage,page))
}

func (this *TradeApp)getCustomerData(row int)(retInfo Model.CustomerInformation) {
	retInfo.CompanyName = this.table_CSV.Item(row,COLUMN_CompanyName).Text()
	retInfo.YearFounded, _ = strconv.Atoi(this.table_CSV.Item(row,COLUMN_YearFounded).Text())
	retInfo.EmployeesNum, _ = strconv.Atoi(this.table_CSV.Item(row,COLUMN_EmployeesNum).Text())
	retInfo.AnnualRevenue, _ = strconv.Atoi(this.table_CSV.Item(row,COLUMN_AnnualRevenue).Text())
	retInfo.PhoneNumber = this.table_CSV.Item(row,COLUMN_PhoneNumber).Text()
	retInfo.Email = this.table_CSV.Item(row,COLUMN_Email).Text()
	retInfo.KeyContact = this.table_CSV.Item(row,COLUMN_KeyContact).Text()
	retInfo.HunterData = this.table_CSV.Item(row,COLUMN_HunterData).Text()
	retInfo.CompanyDomain = this.table_CSV.Item(row,COLUMN_CompanyDomain).Text()
	retInfo.CustomerType = this.table_CSV.Item(row,COLUMN_CustomerType).Text()
	retInfo.CustomerRating = this.table_CSV.Item(row,COLUMN_CustomerRating).Text()
	retInfo.FollowupDate, _ = time.Parse("2006-01-02",this.table_CSV.Item(row,COLUMN_FollowupDate).Text())
	retInfo.Remarks = this.table_CSV.Item(row,COLUMN_Remarks).Text()
	retInfo.Description = this.table_CSV.Item(row,COLUMN_Description).Text()
	retInfo.FacebookPage = this.table_CSV.Item(row,COLUMN_FacebookPage).Text()
	retInfo.LinkedInPage = this.table_CSV.Item(row,COLUMN_LinkedInPage).Text()
	retInfo.TwitterHandle = this.table_CSV.Item(row,COLUMN_TwitterHandle).Text()
	retInfo.Industry = this.table_CSV.Item(row,COLUMN_Industry).Text()
	retInfo.Country = this.table_CSV.Item(row,COLUMN_Country).Text()
	retInfo.City = this.table_CSV.Item(row,COLUMN_City).Text()
	retInfo.StreetAddress = this.table_CSV.Item(row,COLUMN_StreetAddress).Text()
	retInfo.PostalCode = this.table_CSV.Item(row,COLUMN_PostalCode).Text()
	retInfo.Coordinator = this.table_CSV.Item(row,COLUMN_Coordinator).Text()
	retInfo.CreatTime, _ = time.Parse("2006-01-02",this.table_CSV.Item(row,COLUMN_CreatTime).Text())
	return retInfo
}