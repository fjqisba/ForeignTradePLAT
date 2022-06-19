package XCApp

import (
	"github.com/twgh/xcgui/listitemtemplate"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

//加载功能窗口
//bAdmin为true表示为管理员模式

func (this *XCApp)loadMainWindow(bAdmin bool) {

	//加载主窗口
	this.wnd_Main = window.NewWindowByLayoutZipMem(TradeUI, "main.xml", "", 0, 0)

	布局_菜单 := widget.NewLayoutEleByName("布局_菜单")
	this.menu_App = widget.NewMenuBar(0,0,100,20,布局_菜单.Handle)
	this.menu_App.LayoutItem_SetWidth(xcc.Layout_Size_Fixed,55)
	this.menu_App.AddButton(" 程序    ")

	widget.NewMenuByHandle(this.menu_App.GetMenu(0)).AddItem(0,"导入CSV",0,xcc.Menu_Item_Flag_Normal)
	widget.NewMenuByHandle(this.menu_App.GetMenu(0)).AddItem(1,"退出",0,xcc.Menu_Item_Flag_Normal)
	this.menu_App.Event_MENU_SELECT(this.OnMenuApp_Selected)

	//管理员模式才有的功能
	if bAdmin == true{
		this.menu_Admin = widget.NewMenuBar(0,0,100,20,布局_菜单.Handle)
		this.menu_Admin.LayoutItem_SetWidth(xcc.Layout_Size_Fixed,70)
		this.menu_Admin.AddButton(" 管理员    ")
	}

	this.menu_About = widget.NewMenuBar(0,0,100,20,布局_菜单.Handle)
	this.menu_About.LayoutItem_SetWidth(xcc.Layout_Size_Fixed,57)
	this.menu_About.AddButton(" 关于    ")
	widget.NewMenuByHandle(this.menu_About.GetMenu(0)).AddItem(0,"关于软件",0,xcc.Menu_Item_Flag_Normal)
	this.menu_About.Event_MENU_SELECT(this.OnMenuAbout_Selected)

	//列表数据
	this.list_Csv = widget.NewListByName("列表_数据")//widget.NewList(100,200,30,30,this.wnd_Main.Handle)
	csvHeadTemplate := listitemtemplate.NewListItemTemplate_LoadZipMem(
		xcc.ListItemTemp_Type_List_Head ,TradeUI,"listItem.xml","")
	csvItemTemplate := listitemtemplate.NewListItemTemplate_LoadZipMem(
		xcc.ListItemTemp_Type_List_Item,TradeUI,"listItem.xml","")
	this.list_Csv.SetItemTemplate(csvItemTemplate.Handle)
	this.list_Csv.SetItemTemplate(csvHeadTemplate.Handle)

	this.list_Csv.CreateAdapterHeader()
	this.list_Csv.CreateAdapter()
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)
	this.list_Csv.AddColumn(100);this.list_Csv.AddColumn(100)

	//添加事件
	this.list_Csv.Event_LBUTTONDBCLICK(this.on_EditCSV)

	this.wnd_Main.AdjustLayout()
	this.wnd_Main.ShowWindow(xcc.SW_SHOW)
}