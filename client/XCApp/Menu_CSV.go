package XCApp

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

//列表右键单击

func (this *XCApp)on_CSVRightClicked(nFlags int, pPt *xc.POINT, pbHandled *bool)int  {

	var item,subItem int
	if false == this.list_Csv.HitTestOffset(pPt,&item,&subItem){
		return 0
	}

	//if this.list_Csv.GetSelectItemCount() == 0{
	//	this.list_Csv.SetSelectItem(item)
	//	this.list_Csv.Redraw(true)
	//}

	//计算出弹出菜单的屏幕坐标
	var offRect xc.RECT
	this.list_Csv.GetWndClientRect(&offRect)
	pPt.X = pPt.X + offRect.Left
	pPt.Y = pPt.Y + offRect.Top
	xc.ClientToScreen(this.wnd_Main.Handle,pPt)

	//创建列表菜单
	this.menu_Csv = widget.NewMenu()
	//this.menu_Csv.AddItem(1,"删除整行",0,xcc.Menu_Item_Flag_Normal)
	this.menu_Csv.AddItem(2,"测试菜单",0,0)
	if this.menu_Csv.Popup(0,int(pPt.X),int(pPt.Y),this.list_Csv.Handle,xcc.Menu_Popup_Position_Left_Top) == false{
		return 0
	}

	return 0
}


//菜单被选择
func (this *XCApp)on_CsvMenuSelected(nID int, pbHandled *bool)int  {

	//删除选择的行
	//if nID == 1{
	//	pSelectBuffer := make([]int32,100)
	//	nSelectCount := this.list_Csv.GetSelectAll(&pSelectBuffer,100)
	//	log.Println(nSelectCount)
	//	for i:=nSelectCount-1;i>=0;i--{
	//		this.list_Csv.DeleteItem(int(pSelectBuffer[i]))
	//	}
	//	this.list_Csv.Redraw(false)
	//	return 0
	//}

	return 0
}