package XCApp

import (
	"client/Model"
	"github.com/tealeg/xlsx/v3"
	"github.com/twgh/xcgui/common"
	"github.com/twgh/xcgui/wapi"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
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
		"Region":19,
		"City":20,
		"Street Address":21,
		"Postal Code":22,
		"LinkedIn Bio":23,
		"Coordinator":24,
		"Creation Date":25,
	}
)


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
	retInfo.Region = fGetString(indexArray[idx]);idx++
	retInfo.City = fGetString(indexArray[idx]);idx++
	retInfo.StreetAddress = fGetString(indexArray[idx]);idx++
	retInfo.PostalCode = fGetString(indexArray[idx]);idx++
	retInfo.LinkedInBio = fGetString(indexArray[idx]);idx++
	retInfo.Coordinator = fGetString(indexArray[idx]);idx++
	retInfo.CreatTime = fGetTime(indexArray[idx]);idx++
	if retInfo.CreatTime.Format("20060102") == "00010101"{
		retInfo.CreatTime = time.Now()
	}
	return retInfo
}

func (this *XCApp)loadXlsx(hFile *xlsx.File)int  {

	//索引数组,用来确定数据的读取顺序
	indexArray := make([]int,len(gIndexMap))
	for ti:=0;ti<len(indexArray);ti++{
		indexArray[ti] = -1
	}

	for _,sh := range hFile.Sheets {
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
			itemIndex := this.list_Csv.AddItemText(tmpCustomerInfo.CompanyName)
			this.list_Csv.SetItemText(itemIndex,1,strconv.Itoa(tmpCustomerInfo.YearFounded))
			this.list_Csv.SetItemText(itemIndex,2,strconv.Itoa(tmpCustomerInfo.EmployeesNum))
			this.list_Csv.SetItemText(itemIndex,3,strconv.Itoa(tmpCustomerInfo.AnnualRevenue))
			this.list_Csv.SetItemText(itemIndex,4,strconv.Itoa(tmpCustomerInfo.AnnualRevenue))
			this.list_Csv.SetItemText(itemIndex,5,tmpCustomerInfo.Email)
			this.list_Csv.SetItemText(itemIndex,6,tmpCustomerInfo.KeyContact)
			this.list_Csv.SetItemText(itemIndex,7,tmpCustomerInfo.HunterData)
			this.list_Csv.SetItemText(itemIndex,8,tmpCustomerInfo.CompanyDomain)
			this.list_Csv.SetItemText(itemIndex,9,tmpCustomerInfo.CustomerType)
			this.list_Csv.SetItemText(itemIndex,10,tmpCustomerInfo.CustomerRating)
			this.list_Csv.SetItemText(itemIndex,11,tmpCustomerInfo.FollowupDate.Format("2006-01-02"))
			this.list_Csv.SetItemText(itemIndex,12,tmpCustomerInfo.Remarks)
			this.list_Csv.SetItemText(itemIndex,13,tmpCustomerInfo.Description)
			this.list_Csv.SetItemText(itemIndex,14,tmpCustomerInfo.FacebookPage)
			this.list_Csv.SetItemText(itemIndex,15,tmpCustomerInfo.LinkedInPage)
			this.list_Csv.SetItemText(itemIndex,16,tmpCustomerInfo.TwitterHandle)
			this.list_Csv.SetItemText(itemIndex,17,tmpCustomerInfo.Industry)
			this.list_Csv.SetItemText(itemIndex,18,tmpCustomerInfo.Country)
			this.list_Csv.SetItemText(itemIndex,19,tmpCustomerInfo.Region)
			this.list_Csv.SetItemText(itemIndex,20,tmpCustomerInfo.City)
			this.list_Csv.SetItemText(itemIndex,21,tmpCustomerInfo.StreetAddress)
			this.list_Csv.SetItemText(itemIndex,22,tmpCustomerInfo.PostalCode)
			this.list_Csv.SetItemText(itemIndex,23,tmpCustomerInfo.LinkedInBio)
			this.list_Csv.SetItemText(itemIndex,24,tmpCustomerInfo.Coordinator)
			this.list_Csv.SetItemText(itemIndex,25,tmpCustomerInfo.CreatTime.Format("2006-01-02"))
		}
	}
	return 0
}

func (this *XCApp)importClientData()int  {

	c := "\x00"
	lpstrFilter := strings.Join([]string{"表格(csv、xlsx)", "*.csv;*.xlsx", "任意文件(*.*)", "*.*"}, c) + c + c
	lpstrFile := make([]uint16, 260)
	lpstrFileTitle := make([]uint16, 260)
	ofn := wapi.OpenFileNameW{
		LStructSize:       76,
		HwndOwner:         0,
		HInstance:         0,
		LpstrFilter:       common.StringToUint16Ptr(lpstrFilter),
		LpstrCustomFilter: nil,
		NMaxCustFilter:    0,
		NFilterIndex:      1,
		LpstrFile:         &lpstrFile[0],
		NMaxFile:          260,
		LpstrFileTitle:    &lpstrFileTitle[0],
		NMaxFileTitle:     260,
		LpstrInitialDir:   common.StrPtr("C:"),
		LpstrTitle:        common.StrPtr("打开文件"),
		Flags:             wapi.OFN_PATHMUTEXIST, // 用户只能键入有效的路径和文件名
		NFileOffset:       0,
		NFileExtension:    0,
		LpstrDefExt:       0,
		LCustData:         0,
		LpfnHook:          0,
		LpTemplateName:    0,
	}
	ofn.LStructSize = uint32(unsafe.Sizeof(ofn))
	if wapi.GetOpenFileNameW(&ofn) == false{
		return 0
	}
	filePath := syscall.UTF16ToString(lpstrFile)
	hXlsx, err := xlsx.OpenFile(filePath)
	if err == nil{
		return this.loadXlsx(hXlsx)
	}
	return 0
}