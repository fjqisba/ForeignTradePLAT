package XCApp

import (
	"client/api"
	"encoding/json"
	"github.com/twgh/xcgui/widget"
)

func (this *XCApp)executeHunterApi()  {
	rowCount := this.list_Csv.GetCount_AD()
	for i:=0;i<rowCount;i++{
		var records []string
		for j:=0;j<24;j++{
			records = append(records,this.list_Csv.GetItemText(i,j))
		}
		domain := records[8]
		hunterData := records[7]
		if hunterData != ""{
			continue
		}
		emailList := api.SearchDomain(domain)
		emailStr, _ := json.Marshal(emailList)
		hShapeTextHandle := this.list_Csv.GetTemplateObject(i,7,8)
		if hShapeTextHandle == 0{
			continue
		}
		this.list_Csv.SetItemText(i,7,string(emailStr))
		widget.NewShapeTextByHandle(hShapeTextHandle).SetText(string(emailStr))
	}
}
