package TradeApp

import (
	"github.com/therecipe/qt/widgets"
	"qclient/Model"
	"qclient/Utils"

	"qclient/api"
	"qclient/global"
)

type selectAction struct {
	act *widgets.QAction
	app *TradeApp
}

func (this *selectAction)onAction_AssignTaskCilcked(checked bool)  {
	target := this.act.Text()
	tableCsv :=  this.app.table_CSV
	rowList := tableCsv.SelectionModel().SelectedRows(0)
	assignTask := Model.AssignTaskReq{Target:target}
	for _,eRow := range rowList{
		tableCsv.Item(eRow.Row(),COLUMN_Coordinator).SetText(target)
		assignTask.Domain = append(assignTask.Domain, tableCsv.Item(eRow.Row(),COLUMN_CompanyDomain).Text())
	}
	err := api.AssignTask(assignTask)
	if err != nil{
		Utils.MessageBox_Error(this.app.wnd_Main,"分配任务","出错:"+err.Error())
		return
	}
	Utils.MessageBox_Info(this.app.wnd_Main,"分配任务","分配任务成功")
}

func (this *TradeApp)UpdateUserList()  {
	if global.GUserLevel != Model.USER_ADMIN{
		return
	}
	userList,err := api.GetUserList()
	if err != nil{
		return
	}
	for _,eAction := range this.tmpActionList{
		eAction.act.DeleteLater()
		this.menu_AssignTask.RemoveAction(eAction.act)
	}
	this.tmpActionList = []selectAction{}
	for _,eUser := range userList{
		tmpAct := selectAction{this.menu_AssignTask.AddAction(eUser),this}
		tmpAct.act.ConnectTriggered(tmpAct.onAction_AssignTaskCilcked)
		this.tmpActionList = append(this.tmpActionList, tmpAct)
	}
	global.GUserList = userList
}
