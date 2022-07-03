package TradeApp

import (
	"bytes"
	"encoding/json"
	"github.com/tealeg/xlsx/v3"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"qclient/Model"
	"qclient/Utils"
	"qclient/global"
)

func (this *TradeApp)uploadXlsx(filePath string)error  {
	hXlsx, err := xlsx.OpenFile(filePath)
	if err != nil{
		return err
	}
	//索引数组,用来确定数据的读取顺序
	indexArray := make([]int,len(gIndexMap))
	for ti:=0;ti<len(indexArray);ti++{
		indexArray[ti] = -1
	}

	var uploadList []Model.CustomerInformation
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
			uploadList = append(uploadList, tmpCustomerInfo)
		}
	}

	uploadBytes,_ := json.Marshal(uploadList)
	resp,err := http.Post(global.GServerUrl + "/upload","application/json",bytes.NewReader(uploadBytes))
	if err != nil{
		return nil
	}
	defer resp.Body.Close()
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil
	}
	g := gjson.ParseBytes(respBytes)
	if g.Get("code").Int() != 0{
		Utils.MessageBox_Error(this.wnd_Main,"上传数据","上传数据错误:" + g.Get("msg").String())
		return nil
	}
	Utils.MessageBox_Info(this.wnd_Main,"上传数据","上传客户数据成功")
	return nil
}