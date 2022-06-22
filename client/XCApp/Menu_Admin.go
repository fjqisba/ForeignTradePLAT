package XCApp

import (
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/xcc"
)

//菜单_管理员被选择

func (this *XCApp)OnMenuAdmin_Selected(nID int, pbHandled *bool)int  {

	//获取Hunter数据
	if nID == 201{
		if wapi.MessageBoxW(this.wnd_Main.GetHWND(),
			"是否要开始批量获取Hunter数据?","请确认",wapi.MB_IconAsterisk|wapi.MB_YesNo) == wapi.ID_NO{
			return 0
		}
		widget.NewMenuByHandle(this.menu_Admin.GetMenu(0)).SetItemFlags(nID,int(xcc.Menu_Item_Flag_Disable))
		this.executeHunterApi()
		widget.NewMenuByHandle(this.menu_Admin.GetMenu(0)).SetItemFlags(nID,int(xcc.Menu_Item_Flag_Normal))
	}

	return 0
}