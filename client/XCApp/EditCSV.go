package XCApp

import (
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/xc"
	"github.com/twgh/xcgui/xcc"
	"log"
)

func (this *XCApp)on_EditKeyDown(wParam int, lParam int, pbHandled *bool)int  {

	if wParam == 13{

	}
	log.Println(wParam)

	return 0
}

func (this *XCApp)on_EditCSV(nFlags int, pPt *xc.POINT, pbHandled *bool)int  {
	var itemIndex,subItemIndex int
	if this.list_Csv.HitTest(pPt,&itemIndex,&subItemIndex) == false{
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
		tmpEdit := widget.NewEdit(int(rect.Left),int(rect.Top),
			int(rect.Right-rect.Left),int(rect.Bottom-rect.Top),this.list_Csv.Handle)
		this.wnd_Main.SetFocusEle(tmpEdit.Handle)
		oContent := hShapeText.GetText()
		tmpEdit.SetText(oContent)
		tmpEdit.SetUserData(itemIndex)
		tmpEdit.Event_KEYDOWN(this.on_EditKeyDown)
	}


	return 0
}