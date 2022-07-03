package TradeApp

import (
	"encoding/csv"
	"github.com/tealeg/xlsx/v3"
	"github.com/therecipe/qt/widgets"
	"os"
	"qclient/Model"
	"strconv"
	"strings"
	"time"
)

var(
	//读取的顺序
	gIndexMap map[string]int = map[string]int{
		"Company Name":0,
		"Year Founded":1,
		"Number of Employees":2,
		"Annual Revenue":3,
		"Phone Number":4,
		"Email":5,
		"Key Contact":6,
		"Hunter Data":7,
		"Company Domain Name":8,
		"Customer Type":9,
		"Customer Rating":10,
		"Followup Date":11,
		"Remarks":12,
		"Description":13,
		"Facebook Page":14,
		"LinkedIn Company Page":15,
		"Twitter Handle":16,
		"Industry":17,
		"Country":18,
		"City":19,
		"Street Address":20,
		"Postal Code":21,
		"Coordinator":22,
		"Creation Date":23,
	}
)

func getCustomerInformationCsv(records []string,indexArray []int)(retInfo Model.CustomerInformation) {

	maxRecordsLen := len(records)

	fGetString := func(col int)string {
		if col < 0 || col >= maxRecordsLen{
			return ""
		}
		return records[col]
	}
	fGetInt := func(col int)int {
		if col < 0 || col >= maxRecordsLen{
			return 0
		}
		i, err := strconv.Atoi(records[col])
		if err != nil{
			return 0
		}
		return i
	}

	fGetTime := func(col int)time.Time {
		if col < 0 || col >= maxRecordsLen{
			return time.Time{}
		}
		iValue,err := time.Parse("2006-01-02",records[col])
		if err != nil{
			return time.Time{}
		}
		return iValue
	}

	idx := 0
	retInfo.CompanyName = fGetString(indexArray[idx]);idx++
	retInfo.YearFounded = fGetInt(indexArray[idx]);idx++
	retInfo.EmployeesNum = fGetInt(indexArray[idx]);idx++
	retInfo.AnnualRevenue = fGetInt(indexArray[idx]);idx++
	retInfo.PhoneNumber = fGetString(indexArray[idx]);idx++
	retInfo.Email = fGetString(indexArray[idx]);idx++
	retInfo.KeyContact = fGetString(indexArray[idx]);idx++
	retInfo.HunterData = fGetString(indexArray[idx]);idx++
	retInfo.CompanyDomain = fGetString(indexArray[idx]);idx++
	retInfo.CustomerType = fGetString(indexArray[idx]);idx++
	retInfo.CustomerRating = fGetString(indexArray[idx]);idx++
	retInfo.FollowupDate = fGetTime(indexArray[idx]);idx++
	retInfo.Remarks = fGetString(indexArray[idx]);idx++
	retInfo.Description = fGetString(indexArray[idx]);idx++
	retInfo.FacebookPage = fGetString(indexArray[idx]);idx++
	retInfo.LinkedInPage = fGetString(indexArray[idx]);idx++
	retInfo.TwitterHandle = fGetString(indexArray[idx]);idx++
	retInfo.Industry = fGetString(indexArray[idx]);idx++
	retInfo.Country = fGetString(indexArray[idx]);idx++
	retInfo.City = fGetString(indexArray[idx]);idx++
	retInfo.StreetAddress = fGetString(indexArray[idx]);idx++
	retInfo.PostalCode = fGetString(indexArray[idx]);idx++
	retInfo.Coordinator = fGetString(indexArray[idx]);idx++
	retInfo.CreatTime = fGetTime(indexArray[idx]);idx++
	if retInfo.CreatTime.Format("20060102") == "00010101"{
		retInfo.CreatTime = time.Now()
	}
	retInfo.PhoneNumber = strings.TrimSpace(retInfo.PhoneNumber)
	return retInfo
}

func getCustomerInformation(hRow *xlsx.Row,indexArray []int)(retInfo Model.CustomerInformation)  {

	fGetString := func(col int)string {
		if col < 0{
			return ""
		}
		return hRow.GetCell(col).String()
	}

	fGetInt := func(col int)int {
		if col < 0{
			return 0
		}
		iValue,err := hRow.GetCell(col).Int()
		if err != nil{
			return 0
		}
		return iValue
	}

	fGetTime := func(col int)time.Time {
		if col < 0{
			return time.Time{}
		}
		iValue,err := hRow.GetCell(col).GetTime(false)
		if err != nil{
			return time.Time{}
		}
		return iValue
	}

	idx := 0
	retInfo.CompanyName = fGetString(indexArray[idx]);idx++
	retInfo.YearFounded = fGetInt(indexArray[idx]);idx++
	retInfo.EmployeesNum = fGetInt(indexArray[idx]);idx++
	retInfo.AnnualRevenue = fGetInt(indexArray[idx]);idx++
	retInfo.PhoneNumber = fGetString(indexArray[idx]);idx++
	retInfo.Email = fGetString(indexArray[idx]);idx++
	retInfo.KeyContact = fGetString(indexArray[idx]);idx++
	retInfo.HunterData = fGetString(indexArray[idx]);idx++
	retInfo.CompanyDomain = fGetString(indexArray[idx]);idx++
	retInfo.CustomerType = fGetString(indexArray[idx]);idx++
	retInfo.CustomerRating = fGetString(indexArray[idx]);idx++
	retInfo.FollowupDate = fGetTime(indexArray[idx]);idx++
	retInfo.Remarks = fGetString(indexArray[idx]);idx++
	retInfo.Description = fGetString(indexArray[idx]);idx++
	retInfo.FacebookPage = fGetString(indexArray[idx]);idx++
	retInfo.LinkedInPage = fGetString(indexArray[idx]);idx++
	retInfo.TwitterHandle = fGetString(indexArray[idx]);idx++
	retInfo.Industry = fGetString(indexArray[idx]);idx++
	retInfo.Country = fGetString(indexArray[idx]);idx++
	retInfo.City = fGetString(indexArray[idx]);idx++
	retInfo.StreetAddress = fGetString(indexArray[idx]);idx++
	retInfo.PostalCode = fGetString(indexArray[idx]);idx++
	retInfo.Coordinator = fGetString(indexArray[idx]);idx++
	retInfo.CreatTime = fGetTime(indexArray[idx]);idx++
	if retInfo.CreatTime.Format("20060102") == "00010101"{
		retInfo.CreatTime = time.Now()
	}
	retInfo.PhoneNumber = strings.TrimSpace(retInfo.PhoneNumber)
	return retInfo
}

func (this *TradeApp)insertCsvItem(info *Model.CustomerInformation)  {

	insertRow := this.table_CSV.RowCount()
	this.table_CSV.InsertRow(insertRow)
	this.table_CSV.SetItem(insertRow,0,widgets.NewQTableWidgetItem2(info.CompanyName,0))

	strYearFounded := strconv.Itoa(info.YearFounded)
	this.table_CSV.SetItem(insertRow,1,widgets.NewQTableWidgetItem2(strYearFounded,0))

	strEmployeesNum := strconv.Itoa(info.EmployeesNum)
	this.table_CSV.SetItem(insertRow,2,widgets.NewQTableWidgetItem2(strEmployeesNum,0))

	strAnnualRevenue := strconv.Itoa(info.AnnualRevenue)
	this.table_CSV.SetItem(insertRow,3,widgets.NewQTableWidgetItem2(strAnnualRevenue,0))

	this.table_CSV.SetItem(insertRow,4,widgets.NewQTableWidgetItem2(info.PhoneNumber,0))
	this.table_CSV.SetItem(insertRow,5,widgets.NewQTableWidgetItem2(info.Email,0))
	this.table_CSV.SetItem(insertRow,6,widgets.NewQTableWidgetItem2(info.KeyContact,0))
	this.table_CSV.SetItem(insertRow,7,widgets.NewQTableWidgetItem2(info.HunterData,0))
	this.table_CSV.SetItem(insertRow,8,widgets.NewQTableWidgetItem2(info.CompanyDomain,0))
	this.table_CSV.SetItem(insertRow,9,widgets.NewQTableWidgetItem2(info.CustomerType,0))
	this.table_CSV.SetItem(insertRow,10,widgets.NewQTableWidgetItem2(info.CustomerRating,0))
	this.table_CSV.SetItem(insertRow,11,widgets.NewQTableWidgetItem2(info.FollowupDate.Format("2006-01-02"),0))
	this.table_CSV.SetItem(insertRow,12,widgets.NewQTableWidgetItem2(info.Remarks,0))
	this.table_CSV.SetItem(insertRow,13,widgets.NewQTableWidgetItem2(info.Description,0))
	this.table_CSV.SetItem(insertRow,14,widgets.NewQTableWidgetItem2(info.FacebookPage,0))
	this.table_CSV.SetItem(insertRow,15,widgets.NewQTableWidgetItem2(info.LinkedInPage,0))
	this.table_CSV.SetItem(insertRow,16,widgets.NewQTableWidgetItem2(info.TwitterHandle,0))
	this.table_CSV.SetItem(insertRow,17,widgets.NewQTableWidgetItem2(info.Industry,0))
	this.table_CSV.SetItem(insertRow,18,widgets.NewQTableWidgetItem2(info.Country,0))
	this.table_CSV.SetItem(insertRow,19,widgets.NewQTableWidgetItem2(info.City,0))
	this.table_CSV.SetItem(insertRow,20,widgets.NewQTableWidgetItem2(info.StreetAddress,0))
	this.table_CSV.SetItem(insertRow,21,widgets.NewQTableWidgetItem2(info.PostalCode,0))
	this.table_CSV.SetItem(insertRow,22,widgets.NewQTableWidgetItem2(info.Coordinator,0))
	this.table_CSV.SetItem(insertRow,23,widgets.NewQTableWidgetItem2(info.CreatTime.Format("2006-01-02"),0))
}

func parseXlsx(filePath string)([]Model.CustomerInformation,error)  {

	var retInfoList []Model.CustomerInformation
	hXlsx, err := xlsx.OpenFile(filePath)
	if err != nil{
		return retInfoList,err
	}
	//索引数组,用来确定数据的读取顺序
	indexArray := make([]int,len(gIndexMap))
	for ti:=0;ti<len(indexArray);ti++{
		indexArray[ti] = -1
	}
	for _,sh := range hXlsx.Sheets {
		rowHeader,err := sh.Row(0)
		if err != nil{
			continue
		}
		//初始化数据读取列表
		for iCol:=0;iCol<sh.MaxCol;iCol++{
			tmpCell := rowHeader.GetCell(iCol)
			orderNum,bExists := gIndexMap[tmpCell.Value]
			if bExists == false{
				continue
			}
			indexArray[orderNum] = iCol
		}
		for iRow:=1;iRow<sh.MaxRow;iRow++{
			hRow,err := sh.Row(iRow)
			if err != nil{
				continue
			}
			//读取出数据
			tmpCustomerInfo := getCustomerInformation(hRow,indexArray)
			if tmpCustomerInfo.CompanyDomain == ""{
				continue
			}
			retInfoList = append(retInfoList, tmpCustomerInfo)
		}
	}
	return retInfoList,nil
}

func (this *TradeApp)loadXlsx(filePath string)error  {
	hXlsx, err := xlsx.OpenFile(filePath)
	if err != nil{
		return err
	}
	//索引数组,用来确定数据的读取顺序
	indexArray := make([]int,len(gIndexMap))
	for ti:=0;ti<len(indexArray);ti++{
		indexArray[ti] = -1
	}

	for _,sh := range hXlsx.Sheets {
		rowHeader,err := sh.Row(0)
		if err != nil{
			continue
		}
		//初始化数据读取列表
		for iCol:=0;iCol<sh.MaxCol;iCol++{
			tmpCell := rowHeader.GetCell(iCol)
			orderNum,bExists := gIndexMap[tmpCell.Value]
			if bExists == false{
				continue
			}
			indexArray[orderNum] = iCol
		}
		this.table_CSV.SetSortingEnabled(false)
		for iRow:=1;iRow<sh.MaxRow;iRow++{
			hRow,err := sh.Row(iRow)
			if err != nil{
				continue
			}
			//读取出数据
			tmpCustomerInfo := getCustomerInformation(hRow,indexArray)
			if tmpCustomerInfo.CompanyDomain == ""{
				continue
			}
			this.insertCsvItem(&tmpCustomerInfo)
		}
		this.table_CSV.SetSortingEnabled(true)
	}

	return nil
}

func parseCsv(filePath string)([]Model.CustomerInformation,error) {
	var retInfoList []Model.CustomerInformation
	hFile,err := os.Open(filePath)
	if err != nil{
		return retInfoList,err
	}
	hCsvReader := csv.NewReader(hFile)
	csvHeader,err := hCsvReader.Read()
	if err != nil{
		return retInfoList,err
	}
	//索引数组,用来确定数据的读取顺序
	indexArray := make([]int,len(gIndexMap))
	for ti:=0;ti<len(indexArray);ti++{
		indexArray[ti] = -1
	}
	//初始化数据读取列表
	for iCol:=0;iCol<len(csvHeader);iCol++{
		tmpHeaderName := csvHeader[iCol]
		orderNum,bExists := gIndexMap[tmpHeaderName]
		if bExists == false{
			continue
		}
		indexArray[orderNum] = iCol
	}
	for true{
		vec_Records,err := hCsvReader.Read()
		if err != nil{
			break
		}
		tmpCustomerInfo := getCustomerInformationCsv(vec_Records,indexArray)
		if tmpCustomerInfo.CompanyDomain == ""{
			continue
		}
		retInfoList = append(retInfoList, tmpCustomerInfo)
	}
	return retInfoList,nil
}

func (this *TradeApp)loadCsv(filePath string)error {
	hFile,err := os.Open(filePath)
	if err != nil{
		return err
	}
	hCsvReader := csv.NewReader(hFile)
	csvHeader,err := hCsvReader.Read()
	if err != nil{
		return err
	}
	//索引数组,用来确定数据的读取顺序
	indexArray := make([]int,len(gIndexMap))
	for ti:=0;ti<len(indexArray);ti++{
		indexArray[ti] = -1
	}
	//初始化数据读取列表
	for iCol:=0;iCol<len(csvHeader);iCol++{
		tmpHeaderName := csvHeader[iCol]
		orderNum,bExists := gIndexMap[tmpHeaderName]
		if bExists == false{
			continue
		}
		indexArray[orderNum] = iCol
	}
	this.table_CSV.SetSortingEnabled(false)
	for true{
		vec_Records,err := hCsvReader.Read()
		if err != nil{
			break
		}
		tmpCustomerInfo := getCustomerInformationCsv(vec_Records,indexArray)
		if tmpCustomerInfo.CompanyDomain == ""{
			continue
		}
		this.insertCsvItem(&tmpCustomerInfo)
	}
	this.table_CSV.SetSortingEnabled(true)

	return nil
}