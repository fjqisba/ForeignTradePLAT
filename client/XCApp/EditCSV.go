package XCApp

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
)

func (this *XCApp)updateEditBoxData()int  {

	if xc.XC_GetObjectType(this.tmpEditBox.Handle) != xcc.XC_EDIT{
		return 0
	}
	var newTxt string
	this.tmpEditBox.GetText(&newTxt,1048576)
	originHandle := this.tmpEditBox.GetUserData()
	if xc.XC_GetObjectType(originHandle) == xcc.XC_SHAPE_TEXT{
		hShapeBox := widget.NewShapeTextByHandle(originHandle)
		hShapeBox.SetText(newTxt)
	}
	this.tmpEditBox.Destroy()
	this.list_Csv.Redraw(false)
	return 0
}

func (this *XCApp)on_FinishEditBox(hEle int, pbHandled *bool)int  {
	if xc.XC_GetObjectType(hEle) == xcc.XC_EDIT{
		this.updateEditBoxData()
	}
	return 0
}

func (this *XCApp)on_EditKeyDown(wParam int, lParam int, pbHandled *bool)int  {
	if wParam == 13{
		return this.updateEditBoxData()
	}
	return 0
}

func (this *XCApp)on_EditCSV(nFlags int, pPt *xc.POINT, pbHandled *bool)int  {
	var itemIndex,subItemIndex int
	if this.list_Csv.HitTestOffset(pPt,&itemIndex,&subItemIndex) == false{
		return 0
	}
	hShapeTextHandle := this.list_Csv.GetTemplateObject(itemIndex,subItemIndex,subItemIndex+1)
	if hShapeTextHandle == 0{
		return 0
	}
	hShapeText := widget.NewShapeTextByHandle(hShapeTextHandle)
	hLayOutHandle := hShapeText.GetParentEle()
	if xc.XC_GetObjectType(hLayOutHandle) == xcc.XC_ELE_LAYOUT{
		hElement := widget.NewLayoutEleByHandle(hLayOutHandle)
		var rect xc.RECT
		hElement.GetRect(&rect)
		this.tmpEditBox = widget.NewEdit(int(rect.Left),int(rect.Top),
			int(rect.Right-rect.Left),int(rect.Bottom-rect.Top),this.list_Csv.Handle)
		this.wnd_Main.SetFocusEle(this.tmpEditBox.Handle)
		oContent := hShapeText.GetText()
		this.tmpEditBox.SetText(oContent)
		this.tmpEditBox.SetUserData(hShapeText.Handle)
		this.tmpEditBox.Event_KEYDOWN(this.on_EditKeyDown)
		this.tmpEditBox.Event_KILLFOCUS1(this.on_FinishEditBox)
	}

	return 0
}