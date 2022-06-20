package XCApp

import (
	"client/Model"
	"client/api"
	"client/global"
	_ "embed"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/wapi"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
)

//go:embed res/Trade.zip
var TradeUI []byte

type XCApp struct {
	app *app.App
	//登录窗口
	wnd_Login *window.Window
	button_Login *widget.Button
	edit_UserName *widget.Edit
	edit_PassWord *widget.Edit
	//主窗口
	wnd_Main *window.Window
	//列表数据
	list_Csv *widget.List
	//菜单_程序
	menu_App *widget.MenuBar
	//菜单_关于
	menu_About *widget.MenuBar
	//菜单_管理员
	menu_Admin *widget.MenuBar
	//临时编辑框
	tmpEditBox *widget.Edit

	//列表菜单
	menu_Csv *widget.Menu
}

func NewXCApp()*XCApp  {
	return &XCApp{}
}

func (this *XCApp)on_buttonLogin_Clicked(pbHandled *bool)int  {

	var userName string
	var passWord string

	nLen := this.edit_UserName.GetText(&userName,256)
	if nLen == 0{
		return 0
	}
	nLen = this.edit_PassWord.GetText(&passWord,256)
	if nLen == 0{
		return 0
	}

	userRole := api.Login(userName,passWord)
	if userRole == Model.LOGIN_INVALID{
		wapi.MessageBoxW(this.wnd_Login.GetHWND(),"用户名密码错误","错误",wapi.MB_IconError)
		return 0
	}

	if userRole == Model.LOGIN_ADMIN{
		global.GLOBAL_USERROLE = Model.LOGIN_ADMIN
		this.loadMainWindow(true)
		this.wnd_Login.CloseWindow()
		return 0
	}else if userRole == Model.LOGIN_USER{
		global.GLOBAL_USERROLE = Model.LOGIN_USER
		this.loadMainWindow(false)
		this.wnd_Login.CloseWindow()
		return 0
	}

	return 0
}



//关于菜单

func (this *XCApp)OnMenuAbout_Selected(nID int, pbHandled *bool)int  {

	if nID == 0{
		wapi.MessageBoxW(this.wnd_Main.GetHWND(),"企业内部定制版,禁止外部分享","关于外贸管理平台软件",wapi.MB_IconInformation)
		return 0
	}

	return 0
}



//加载登录窗口

func (this *XCApp)loadLoginWindow()  {

	//加载Login窗口
	this.wnd_Login = window.NewWindowByLayoutZipMem(TradeUI, "login.xml", "", 0, 0)
	this.wnd_Login.AdjustLayout()

	//设置登录窗口控件
	this.edit_UserName = widget.NewEditByName("编辑框_用户名")
	this.edit_PassWord = widget.NewEditByName("编辑框_密码")
	this.button_Login = widget.NewButtonByName("按钮_登录")
	this.button_Login.Event_BnClick(this.on_buttonLogin_Clicked)

	this.wnd_Login.ShowWindow(xcc.SW_SHOW)
}

func (this *XCApp)Run()  {
	//加载登录窗口
	this.loadMainWindow(true)
	this.app.Run()
	//所有的窗口退出,程序就退出
	this.app.Exit()
}

func (this *XCApp)InitXCApp()error  {

	this.app = app.New(true)
	// 从内存zip中加载资源文件
	this.app.LoadResourceZipMem(TradeUI, "resource.res", "")


	return nil
}