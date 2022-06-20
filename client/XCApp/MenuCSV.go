package XCApp

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

//列表右键单击

func (this *XCApp)on_CSVRightClicked(nFlags int, pPt *xc.POINT, pbHandled *bool)int  {

	var item,subItem int
	if false == this.list_Csv.HitTest(pPt,&item,&subItem){
		return 0
	}
	this.list_Csv.SetSelectItem(item)
	this.list_Csv.Redraw(true)

	//计算出弹出菜单的屏幕坐标
	var offRect xc.RECT
	this.list_Csv.GetWndClientRect(&offRect)
	pPt.X = pPt.X + offRect.Left
	pPt.Y = pPt.Y + offRect.Top
	xc.ClientToScreen(this.wnd_Main.Handle,pPt)

	//创建列表菜单
	this.menu_Csv = widget.NewMenu()
	this.menu_Csv.AddItem(1,"测试",0,0)
	if this.menu_Csv.Popup(this.list_Csv.Handle,int(pPt.X),int(pPt.Y),0,xcc.Menu_Popup_Position_Left_Top) == false{
		return 0
	}

	return 0
}