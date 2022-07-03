package ViewHunterDlg

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"qclient/Config"
	"qclient/Model"
)

type ViewHunterWnd struct {
	dlg *widgets.QDialog
	table_HunterData *widgets.QTableWidget
}

func NewViewHunterWnd()*ViewHunterWnd  {
	return &ViewHunterWnd{}
}

func (this *ViewHunterWnd)insertCsvItem(info *Model.EmailInfo) {

	insertRow := this.table_HunterData.RowCount()
	this.table_HunterData.InsertRow(insertRow)
	this.table_HunterData.SetItem(insertRow, 0, widgets.NewQTableWidgetItem2(info.Email, 0))
	this.table_HunterData.SetItem(insertRow, 1, widgets.NewQTableWidgetItem2(info.PersonName, 0))
	this.table_HunterData.SetItem(insertRow, 2, widgets.NewQTableWidgetItem2(info.PhoneNumber, 0))
	this.table_HunterData.SetItem(insertRow, 3, widgets.NewQTableWidgetItem2(info.Position, 0))
	this.table_HunterData.SetItem(insertRow, 4, widgets.NewQTableWidgetItem2(info.Department, 0))
	this.table_HunterData.SetItem(insertRow, 5, widgets.NewQTableWidgetItem2(info.Type, 0))
	this.table_HunterData.SetItem(insertRow, 6, widgets.NewQTableWidgetItem2(info.Linkedin, 0))
	this.table_HunterData.SetItem(insertRow, 7, widgets.NewQTableWidgetItem2(info.Twitter, 0))
}

func onViewHunterCsv_HeaderResized(logicalIndex int, oldSize int, newSize int)  {
	Config.Instance.SaveViewHunterItemWidth(logicalIndex,newSize)
}

func (this *ViewHunterWnd)Run(emailList []Model.EmailInfo)  {

	this.dlg = widgets.NewQDialog(nil,0)
	this.dlg.SetWindowTitle("Hunter数据查看")
	this.dlg.SetWindowIcon(gui.NewQIcon5(":/Trade/trade.png"))
	this.dlg.SetWindowFlags(core.Qt__CustomizeWindowHint|core.Qt__WindowCloseButtonHint)
	this.dlg.Resize2(800,600)
	this.dlg.SetMaximumSize2(1600,1200)
	this.dlg.SetWindowModality(core.Qt__ApplicationModal)

	//建立布局
	gridLayout := widgets.NewQGridLayout(this.dlg)
	gridLayout.SetSpacing(6)
	gridLayout.SetContentsMargins(11,11,11,11)

	this.table_HunterData = widgets.NewQTableWidget(this.dlg)
	this.table_HunterData.SetMinimumSize2(0,500)
	this.table_HunterData.SetAlternatingRowColors(false)
	this.table_HunterData.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	this.table_HunterData.SetSelectionMode(widgets.QAbstractItemView__ExtendedSelection)
	this.table_HunterData.SetShowGrid(false)
	this.table_HunterData.SetSortingEnabled(true)
	this.table_HunterData.HorizontalHeader().ConnectSectionResized(onViewHunterCsv_HeaderResized)
	gridLayout.AddWidget3(this.table_HunterData,0,0,1,1,0)

	this.table_HunterData.SetColumnCount(8)
	this.table_HunterData.SetHorizontalHeaderLabels([]string{
		"邮箱","姓名","电话","职位","部门","邮箱类型","领英","推特"})
	//设置每一列的长度
	for col:=0;col<8;col++{
		this.table_HunterData.SetColumnWidth(col,Config.Instance.GetViewHunterItemWidth(col))
	}
	for _,eInfo := range emailList{
		this.insertCsvItem(&eInfo)
	}
	this.table_HunterData.Show()

	//this.edit_UserName = widgets.NewQLineEdit(this.dlg)
	//gridLayout.AddWidget3(this.edit_UserName,0,1,1,1,0)
	//btn_AddUser := widgets.NewQPushButton2("新增用户",this.dlg)
	//btn_AddUser.ConnectClicked(this.onBtn_AddUserClicked)
	//gridLayout.AddWidget3(btn_AddUser,0,2,1,1,0)
	//
	//this.list_UserName = widgets.NewQListWidget(this.dlg)
	//this.list_UserName.SetContextMenuPolicy(core.Qt__CustomContextMenu)
	//this.list_UserName.SetSelectionMode(widgets.QAbstractItemView__SingleSelection)
	//this.list_UserName.ConnectCustomContextMenuRequested(this.onUserMenu_Clicked)
	//gridLayout.AddWidget3(this.list_UserName,1,0,1,3,0)
	//
	////创建菜单
	//this.menu_UserName = widgets.NewQMenu(this.dlg)
	//this.menu_UserName.AddAction("删除用户").ConnectTriggered(this.onUserMenu_DeleteUser)

	this.dlg.Show()
}

