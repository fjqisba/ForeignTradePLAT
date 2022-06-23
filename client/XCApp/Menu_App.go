package XCApp

import (
	"github.com/twgh/xcgui/wapi"
)

func (this *XCApp)ClearCSVList()int  {
	if wapi.MessageBoxW(this.wnd_Main.GetHWND(),
		"是否要清空列表数据?","请确认",wapi.MB_IconAsterisk|wapi.MB_YesNo) == wapi.ID_NO{
		return 0
	}
	this.list_Csv.DeleteItemAll()
	return 0
}

//菜单_程序被选择

func (this *XCApp)OnMenuApp_Selected(nID int, pbHandled *bool)int  {

	//导入CSV
	if nID == 101{
		return this.importClientData()
	}
	//保存CSV
	if nID == 102{
		return this.ExportCSV()
	}
	//清空CSV
	if nID == 103{
		return this.ClearCSVList()
	}
	//退出
	if nID == 104{
		this.wnd_Main.CloseWindow()
	}
	return 0
}