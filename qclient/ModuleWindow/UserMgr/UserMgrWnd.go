package UserMgr

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"qclient/Utils"
	"qclient/api"
)

type UserMgrWnd struct {
	Dlg *widgets.QDialog
	edit_UserName *widgets.QLineEdit
	list_UserName *widgets.QListWidget
	menu_UserName *widgets.QMenu
}

func NewUserMgrWnd()(retWnd *UserMgrWnd)  {
	retWnd = &UserMgrWnd{}
	retWnd.Dlg = widgets.NewQDialog(nil,0)
	retWnd.Dlg.SetWindowTitle("用户管理")
	retWnd.Dlg.SetWindowIcon(gui.NewQIcon5(":/Trade/trade.png"))
	retWnd.Dlg.SetWindowFlags(core.Qt__CustomizeWindowHint|core.Qt__WindowCloseButtonHint)
	retWnd.Dlg.Resize2(300,400)
	retWnd.Dlg.SetMaximumSize2(600,800)
	retWnd.Dlg.SetWindowModality(core.Qt__ApplicationModal)
	return retWnd
}

func (this *UserMgrWnd)onBtn_AddUserClicked(checked bool)  {
	userName := this.edit_UserName.Text()
	err := api.AddUser(userName)
	if err != nil{
		Utils.MessageBox_Error(this.Dlg,"添加用户","出错:"+err.Error())
		return
	}
	Utils.MessageBox_Info(this.Dlg,"添加用户","添加用户成功")
	this.list_UserName.AddItem(userName)
}

func (this *UserMgrWnd)onUserMenu_Clicked(pos *core.QPoint)  {
	//确保只有在选中项目的时候,才能弹出菜单
	currentItem := this.list_UserName.ItemAt(pos)
	if currentItem.Pointer() == nil{
		return
	}
	this.menu_UserName.Exec2(gui.QCursor_Pos(),nil)
}

func (this *UserMgrWnd)onUserMenu_DeleteUser(checked bool)  {
	rowList := this.list_UserName.SelectedItems()
	for _,eRow := range rowList{
		err := api.DeleteUser(eRow.Text())
		if err != nil{
			Utils.MessageBox_Error(this.Dlg,"删除用户","错误:"+err.Error())
			return
		}
	}
	Utils.MessageBox_Info(this.Dlg,"删除用户","删除用户成功")
}



func (this *UserMgrWnd)Run()  {


	//建立布局
	gridLayout := widgets.NewQGridLayout(this.Dlg)
	gridLayout.SetSpacing(6)
	gridLayout.SetContentsMargins(11,11,11,11)

	label_UserName := widgets.NewQLabel2("用户名:",this.Dlg,0)
	gridLayout.AddWidget3(label_UserName,0,0,1,1,0)

	this.edit_UserName = widgets.NewQLineEdit(this.Dlg)
	gridLayout.AddWidget3(this.edit_UserName,0,1,1,1,0)
	btn_AddUser := widgets.NewQPushButton2("新增用户",this.Dlg)
	btn_AddUser.ConnectClicked(this.onBtn_AddUserClicked)
	gridLayout.AddWidget3(btn_AddUser,0,2,1,1,0)

	this.list_UserName = widgets.NewQListWidget(this.Dlg)
	this.list_UserName.SetContextMenuPolicy(core.Qt__CustomContextMenu)
	this.list_UserName.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
	this.list_UserName.ConnectCustomContextMenuRequested(this.onUserMenu_Clicked)
	gridLayout.AddWidget3(this.list_UserName,1,0,1,3,0)

	//创建菜单
	this.menu_UserName = widgets.NewQMenu(this.Dlg)
	this.menu_UserName.AddAction("删除用户").ConnectTriggered(this.onUserMenu_DeleteUser)

	//初始化信息
	userList,err := api.GetUserList()
	if err != nil{
		Utils.MessageBox_Error(this.Dlg,"用户管理","获取用户列表失败:" + err.Error())
	}
	for _,eUserName := range userList{
		this.list_UserName.AddItem(eUserName)
	}
	this.Dlg.Show()
}