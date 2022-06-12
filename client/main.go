package main

import (
	_ "embed"
	"github.com/twgh/xcgui/app"
	"github.com/twgh/xcgui/widget"
	"github.com/twgh/xcgui/window"
	"github.com/twgh/xcgui/xcc"
	"log"
)


type Button struct {
	widget.Element
}

//go:embed res/TradeLogin.zip
var LoginUI []byte

var(
	Button_Login *widget.Button
	Edit_UserName *widget.Edit
	Edit_PassWord *widget.Edit
)

func ButtonLogin_Clicked(pbHandled *bool)int  {

	var userName string
	var passWord string

	nLen := Edit_UserName.GetText(&userName,256)
	if nLen == 0{
		return 0
	}
	nLen = Edit_PassWord.GetText(&passWord,256)
	if nLen == 0{
		return 0
	}



	return 0
}

func main()  {
	a := app.New(true)
	// 从内存zip中加载资源文件
	a.LoadResourceZipMem(LoginUI, "resource.res", "")
	// 从内存zip中加载布局文件, 创建窗口对象
	w := window.NewWindowByLayoutZipMem(LoginUI, "main.xml", "", 0, 0)
	// 调整布局
	w.AdjustLayout()

	Edit_UserName = widget.NewEditByName("编辑框_用户名")
	Edit_PassWord = widget.NewEditByName("编辑框_密码")
	Button_Login = widget.NewButtonByName("按钮_登录")
	Button_Login.Event_BnClick(ButtonLogin_Clicked)

	// 显示窗口
	w.ShowWindow(xcc.SW_SHOW)
	a.Run()

	log.Println("123123")
	a.Exit()
}