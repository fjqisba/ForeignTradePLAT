package TradeApp

import (
	"encoding/json"
	"qclient/api"
	"sync/atomic"
)

var(
	isExecute atomic.Value
)

func (this *TradeApp)executeHunterApi(rowList []int)  {

	defer func() {
		isExecute.Store(false)
		this.table_CSV.SetSortingEnabled(true)
		this.wnd_Main.SetWindowTitle(WINDOWS_TITLE)
		this.act_ClearCSV.SetEnabled(true)
	}()

	for _,eRow := range rowList{
		if isExecute.Load() == false{
			return
		}
		//没钱,不爬已有数据的
		if this.table_CSV.Item(eRow,COLUMN_HunterData).Text() != ""{
			continue
		}
		domain := this.table_CSV.Item(eRow,COLUMN_CompanyDomain).Text()
		if domain == ""{
			continue
		}
		emailList := api.SearchDomain(domain)
		emailStr, _ := json.Marshal(emailList)
		this.table_CSV.Item(eRow,COLUMN_HunterData).SetText(string(emailStr))
	}
}
