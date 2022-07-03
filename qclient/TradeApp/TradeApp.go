package TradeApp

import (
	"encoding/json"
	"fmt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/widgets"
	"os"
	"path/filepath"
	"qclient/Config"
	"qclient/Const"
	"qclient/Model"
	"qclient/ModuleWindow/UserMgr"
	"qclient/ModuleWindow/ViewHunterDlg"
	"qclient/Utils"
	"qclient/api"
	"qclient/global"
	"strconv"
	"strings"
)

const(
	COLUMN_CompanyName = 0
	COLUMN_YearFounded = 1
	COLUMN_EmployeesNum = 2
	COLUMN_AnnualRevenue = 3
	COLUMN_PhoneNumber = 4
	COLUMN_Email = 5
	COLUMN_KeyContact = 6
	COLUMN_HunterData = 7
	COLUMN_CompanyDomain = 8
	COLUMN_CustomerType = 9
	COLUMN_CustomerRating = 10
	COLUMN_FollowupDate = 11
	COLUMN_Remarks = 12
	COLUMN_Description = 13
	COLUMN_FacebookPage = 14
	COLUMN_LinkedInPage = 15
	COLUMN_TwitterHandle = 16
	COLUMN_Industry = 17
	COLUMN_Country = 18
	COLUMN_City = 19
	COLUMN_StreetAddress = 20
	COLUMN_PostalCode = 21
	COLUMN_Coordinator = 22
	COLUMN_CreatTime = 23
	WINDOWS_TITLE = "外贸客户管理平台"
)


type TradeApp struct {
	core.QObject
	app *widgets.QApplication

	loginDlg *widgets.QDialog
	lineEdit_UserName *widgets.QLineEdit

	wnd_Main *widgets.QMainWindow
	table_CSV *widgets.QTableWidget
	menu_CSV *widgets.QMenu
	//统计页数
	label_CsvCount	*widgets.QLabel

	lineEdit_JumpPage	*widgets.QLineEdit
	//上一页按钮
	btn_GoToLastPage *widgets.QPushButton
	//下一页按钮
	btn_GoToNextPage *widgets.QPushButton

	//分配任务
	menu_AssignTask *widgets.QMenu
	tmpActionList []selectAction
	//清空列表
	act_ClearCSV	*widgets.QAction
}

func NewTradeAppClient()*TradeApp  {
	return &TradeApp{}
}

func (this *TradeApp)onUserMgrWnd_Closed(event *gui.QCloseEvent)  {
	this.UpdateUserList()
	return
}

func (this *TradeApp)onMenu_UserManager(checked bool) {
	userMgrDlg := UserMgr.NewUserMgrWnd()
	userMgrDlg.Dlg.ConnectCloseEvent(this.onUserMgrWnd_Closed)
	userMgrDlg.Run()
}

func (this *TradeApp)onMenu_UpdateCSV(checked bool) {
	rowCount := this.table_CSV.RowCount()
	var customerDataList []Model.CustomerInformation
	for i:=0;i<rowCount;i++{
		customerDataList = append(customerDataList, this.getCustomerData(i))
	}
	err := api.UpdateCustomerData(customerDataList)
	if err != nil{
		Utils.MessageBox_Error(this.wnd_Main,"更新数据","出错:"+err.Error())
		return
	}
	Utils.MessageBox_Info(this.wnd_Main,"更新数据","更新数据成功")
}


func (this *TradeApp)onMenu_UploadCSV(checked bool) {
	filePath := widgets.QFileDialog_GetOpenFileName(this.wnd_Main,"上传客户数据",Config.Instance.GetOpenFileDir(),"表格文件(*.csv *.xlsx)","",widgets.QFileDialog__ReadOnly)
	if filePath == ""{
		return
	}
	Config.Instance.SetOpenFileDir(filepath.Dir(filePath))
	customerInfoList,err := parseXlsx(filePath)
	if err != nil{
		customerInfoList,err = parseCsv(filePath)
	}
	if err != nil{
		Utils.MessageBox_Error(this.wnd_Main,"上传客户数据","出错:"+err.Error())
		return
	}
	err = api.UpdateCustomerData(customerInfoList)
	if err != nil{
		Utils.MessageBox_Error(this.wnd_Main,"上传客户数据","出错:"+err.Error())
		return
	}
	Utils.MessageBox_Info(this.wnd_Main,"上传客户数据","上传客户数据成功")
	return
}

func (this *TradeApp)onMenu_ImportCSV(checked bool)  {
	filePath := widgets.QFileDialog_GetOpenFileName(this.wnd_Main,"导入客户数据",Config.Instance.GetOpenFileDir(),"表格文件(*.csv *.xlsx)","",widgets.QFileDialog__ReadOnly)
	if filePath == ""{
		return
	}
	Config.Instance.SetOpenFileDir(filepath.Dir(filePath))
	err := this.loadXlsx(filePath)
	if err == nil{
		return
	}
	this.loadCsv(filePath)
}

func (this *TradeApp)onMenu_ExitApp(checked bool)  {
	this.wnd_Main.Close()
}

func (this *TradeApp)onMenu_AboutApp(checked bool)  {
	msgBox := widgets.NewQMessageBox2(widgets.QMessageBox__Information,
		"关于外贸管理平台软件","企业内部定制版,禁止外部分享",widgets.QMessageBox__Yes,
		this.wnd_Main,core.Qt__Dialog)
	msgBox.Button(widgets.QMessageBox__Yes).SetText("好的")
	msgBox.Exec()
}

func (this *TradeApp)onKeyDown_JumpPage(event *gui.QKeyEvent)  {
	this.lineEdit_JumpPage.KeyPressEventDefault(event)
	page,err := strconv.Atoi(this.lineEdit_JumpPage.Text())
	if err != nil{
		return
	}
	if page <= 0{
		return
	}
	if event.Key() == int(core.Qt__Key_Enter){
		this.UpdatePage(page,Const.PageNum,global.GOrderBy,global.GUserName)
	}
}

//上一页被点击

func (this *TradeApp)onBtn_LastPageClicked(checked bool)  {
	pageIndex := global.GCurrentPageIndex - 1
	if pageIndex <= 0{
		return
	}
	this.UpdatePage(pageIndex,Const.PageNum,global.GOrderBy,global.GUserName)
}

//下一页被点击

func (this *TradeApp)onBtn_NextPageClicked(checked bool)  {
	pageIndex := global.GCurrentPageIndex + 1
	if pageIndex > global.GTotalPageCount{
		return
	}
	this.UpdatePage(pageIndex,Const.PageNum,global.GOrderBy,global.GUserName)
}

func (this *TradeApp)setupUi()  {

	// 创建主窗口
	this.wnd_Main = widgets.NewQMainWindow(nil, 0)

	//初始化核心布局
	centerWidget := widgets.NewQWidget(this.wnd_Main,core.Qt__Widget)
	gridLayout := widgets.NewQGridLayout(centerWidget)
	gridLayout.SetSpacing(6)
	centerWidget.SetContentsMargins(11,11,11,11)

	//添加表格控件
	this.table_CSV = widgets.NewQTableWidget(centerWidget)
	this.table_CSV.SetMinimumSize2(0,500)
	this.table_CSV.SetAlternatingRowColors(false)
	this.table_CSV.SetSelectionBehavior(widgets.QAbstractItemView__SelectRows)
	this.table_CSV.SetSelectionMode(widgets.QAbstractItemView__ExtendedSelection)
	this.table_CSV.SetShowGrid(false)
	this.table_CSV.SetSortingEnabled(false)
	this.table_CSV.SetContextMenuPolicy(core.Qt__CustomContextMenu)
	gridLayout.AddWidget3(this.table_CSV,0,0,1,5,0)

	//添加空白项
	horizontalSpacer := widgets.NewQSpacerItem(200,20,widgets.QSizePolicy__Expanding,widgets.QSizePolicy__Minimum)
	gridLayout.AddItem(horizontalSpacer,1,0,1,1,0)

	this.label_CsvCount = widgets.NewQLabel2("共1页,跳转到:", centerWidget,0)
	gridLayout.AddWidget3(this.label_CsvCount,1,1,1,1,0)

	this.lineEdit_JumpPage = widgets.NewQLineEdit2("1",centerWidget)
	this.lineEdit_JumpPage.SetMaximumWidth(60)
	this.lineEdit_JumpPage.SetValidator(gui.NewQIntValidator2(1,99999,this.lineEdit_JumpPage))
	this.lineEdit_JumpPage.ConnectKeyPressEvent(this.onKeyDown_JumpPage)

	gridLayout.AddWidget3(this.lineEdit_JumpPage,1,2,1,1,0)

	this.btn_GoToLastPage = widgets.NewQPushButton2("上一页",centerWidget)
	this.btn_GoToLastPage.ConnectClicked(this.onBtn_LastPageClicked)
	gridLayout.AddWidget3(this.btn_GoToLastPage,1,3,1,1,0)

	this.btn_GoToNextPage = widgets.NewQPushButton2("下一页",centerWidget)
	this.btn_GoToNextPage.ConnectClicked(this.onBtn_NextPageClicked)
	gridLayout.AddWidget3(this.btn_GoToNextPage,1,4,1,1,0)

	this.wnd_Main.SetCentralWidget(centerWidget)
}

func (this *TradeApp)onCsvMenu_EmailContact(checked bool)  {

	rowList := this.table_CSV.SelectionModel().SelectedRows(0)
	openCount := 0
	for _,eRow := range rowList{
		emailAddr := this.table_CSV.Item(eRow.Row(),COLUMN_Email).Text()
		emailAddr = strings.TrimSpace(emailAddr)
		if emailAddr == ""{
			continue
		}
		contactUrl := fmt.Sprintf("mailto://%s",emailAddr)
		gui.QDesktopServices_OpenUrl(core.NewQUrl3(contactUrl,core.QUrl__TolerantMode))
		openCount = openCount + 1
		//禁止一次性打开太多
		if openCount >= 10{
			break
		}
	}
	return
}

//获取Hunter数据
func (this *TradeApp)onCsvMenu_GetHunterData(checked bool) {
	if isExecute.Load() != nil && isExecute.Load() == true{
		msgBox := widgets.NewQMessageBox2(widgets.QMessageBox__Critical,
			"获取Hunter数据","已有任务正在运行!",widgets.QMessageBox__Yes,
			this.wnd_Main,core.Qt__Dialog)
		msgBox.Button(widgets.QMessageBox__Yes).SetText("好的")
		msgBox.Exec()
		return
	}
	rowList := this.table_CSV.SelectionModel().SelectedRows(0)
	var vecRow []int
	for _,eRow := range rowList{
		vecRow = append(vecRow, eRow.Row())
	}
	if len(vecRow) == 0{
		return
	}
	isExecute.Store(true)
	this.table_CSV.SetSortingEnabled(false)
	this.wnd_Main.SetWindowTitle(WINDOWS_TITLE + "(获取Hunter数据...)")
	this.act_ClearCSV.SetDisabled(true)
	go this.executeHunterApi(vecRow)
}

//查看Hunter数据

func (this *TradeApp)onCsvMenu_ViewHunter(checked bool)  {

	rowList := this.table_CSV.SelectionModel().SelectedRows(0)
	if len(rowList) == 0{
		return
	}
	hunterData := this.table_CSV.Item(rowList[0].Row(),COLUMN_HunterData).Text()
	//hunterData := `[{"email":"lois@illumco.com","type":"personal","personName":"Lois Hughs","phoneNumber":"+1 404 876 1064","position":"Sales Agent"},{"email":"sales@illumco.com","type":"generic","personName":" ","department":"sales"},{"email":"sharon@illumco.com","type":"personal","personName":" "},{"email":"info@illumco.com","type":"generic","personName":" ","department":"support"},{"email":"b4b@illumco.com","type":"generic","personName":" "},{"email":"websales@illumco.com","type":"generic","personName":" ","department":"sales"}]`
	var emailList []Model.EmailInfo
	err := json.Unmarshal([]byte(hunterData),&emailList)
	if err != nil{
		return
	}
	dlg := ViewHunterDlg.NewViewHunterWnd()
	dlg.Run(emailList)
}

//电话联系按钮被点击
func (this *TradeApp)onCsvMenu_PhoneContact(checked bool)  {

	rowList := this.table_CSV.SelectionModel().SelectedRows(0)
	openCount := 0
	for _,eRow := range rowList{
		phoneNumber := this.table_CSV.Item(eRow.Row(),COLUMN_PhoneNumber).Text()
		phoneNumber = strings.TrimSpace(phoneNumber)
		phoneNumber = strings.TrimPrefix(phoneNumber,"+")
		if phoneNumber == ""{
			continue
		}
		contactUrl := fmt.Sprintf("https://wa.me/%s?text=Hi",phoneNumber)
		gui.QDesktopServices_OpenUrl(core.NewQUrl3(contactUrl,core.QUrl__TolerantMode))
		openCount = openCount + 1
		//禁止一次性打开太多
		if openCount >= 10{
			break
		}
	}
}

func (this *TradeApp)onRightClicked_ShowMenu(pos *core.QPoint)  {
	//确保只有在选中项目的时候,才能弹出菜单
	currentItem := this.table_CSV.ItemAt(pos)
	if currentItem.Pointer() == nil{
		return
	}
	this.menu_CSV.Exec2(gui.QCursor_Pos(),nil)
}

func onTableCSV_HeaderResized(logicalIndex int, oldSize int, newSize int)  {
	Config.Instance.SaveItemWidth(logicalIndex,newSize)
}

func (this *TradeApp)initTableCsvData()  {
	this.UpdatePage(1,Const.PageNum,global.GOrderBy,global.GUserName)
	this.UpdateUserList()
}

func (this *TradeApp)initMainWindow()  {

	this.setupUi()

	this.wnd_Main.SetWindowIcon(gui.NewQIcon5(":/Trade/trade.png"))
	// 设置窗口最小尺寸
	this.wnd_Main.SetMinimumSize2(800, 600)
	// 设置窗口默认全屏
	this.wnd_Main.SetWindowState(core.Qt__WindowMaximized)

	// 设置标题
	if global.GUserLevel == Model.USER_ADMIN{
		this.wnd_Main.SetWindowTitle("外贸客户管理平台(管理员)")
	}else{
		this.wnd_Main.SetWindowTitle("外贸客户管理平台")
	}

	//添加窗口菜单栏
	menuBar := this.wnd_Main.MenuBar()
	menu_App := menuBar.AddMenu2("程序")

	if global.GUserLevel == Model.USER_ADMIN{
		menu_Admin := menuBar.AddMenu2("管理员")
		menu_Admin.AddAction("上传数据").ConnectTriggered(this.onMenu_UploadCSV)
		menu_Admin.AddAction("提交数据").ConnectTriggered(this.onMenu_UpdateCSV)
		menu_Admin.AddAction("用户管理").ConnectTriggered(this.onMenu_UserManager)
	}

	menu_About := menuBar.AddMenu2("关于")
	//this.act_ClearCSV = menu_App.AddAction("清空当前列表")
	//this.act_ClearCSV.ConnectTriggered(this.onMenu_ClearCSV)
	menu_App.AddAction("导出数据").ConnectTriggered(this.onMenu_ExportCSV)
	menu_App.AddAction("退出程序").ConnectTriggered(this.onMenu_ExitApp)
	menu_About.AddAction("关于程序").ConnectTriggered(this.onMenu_AboutApp)
	this.table_CSV.SetColumnCount(24)
	this.table_CSV.SetHorizontalHeaderLabels([]string{
		"公司名称","成立年份","员工数量","年营业额(美元)","联系电话","客户邮箱",
		"联系人名字","Hunter数据","客户网址","客户类型","客户等级","最后跟进日期","备注","客户介绍",
		"Facebook主页","领英主页","推特账号","行业","国家","城市","详细地址","邮编","跟进人","导入日期"})

	//设置每一列的长度
	for col:=0;col<24;col++{
		this.table_CSV.SetColumnWidth(col,Config.Instance.GetItemWidth(col))
	}

	//添加表格菜单
	this.menu_CSV = widgets.NewQMenu(this.wnd_Main)
	this.menu_CSV.AddAction("查看Hunter").ConnectTriggered(this.onCsvMenu_ViewHunter)
	this.menu_CSV.AddAction("电话联系").ConnectTriggered(this.onCsvMenu_PhoneContact)
	this.menu_CSV.AddAction("邮箱联系").ConnectTriggered(this.onCsvMenu_EmailContact)
	if global.GUserLevel == Model.USER_ADMIN{
		this.menu_AssignTask = this.menu_CSV.AddMenu2("分配任务至")
		this.menu_CSV.AddAction("获取Hunter").ConnectTriggered(this.onCsvMenu_GetHunterData)
	}
	this.table_CSV.ConnectCustomContextMenuRequested(this.onRightClicked_ShowMenu)
	this.table_CSV.HorizontalHeader().ConnectSectionResized(onTableCSV_HeaderResized)
	this.initTableCsvData()
}

func (this *TradeApp)InitTradeApp()  {
	// 创建应用程序
	this.app = widgets.NewQApplication(len(os.Args), os.Args)
}

func (this *TradeApp)on_LoginBtnClicked(checked bool)  {

	userName := this.lineEdit_UserName.Text()
	Config.Instance.SetUserName(userName)
	userLevel,err := api.GetUserLevel(userName)
	if err != nil{
		Utils.MessageBox_Error(this.loginDlg,"登录","异常:"+err.Error())
		return
	}
	if userLevel == Model.USER_INVALID{
		Utils.MessageBox_Error(this.loginDlg,"登录","无效的用户名")
		return
	}
	//通过用户名检查
	global.GUserName = userName
	global.GUserLevel = userLevel
	this.initMainWindow()
	this.wnd_Main.Show()
	this.loginDlg.DestroyQDialog()
}

func (this *TradeApp)Run()  {

	this.loginDlg = widgets.NewQDialog(nil,0)
	this.loginDlg.SetWindowTitle(WINDOWS_TITLE + "(登录)")
	this.loginDlg.SetWindowFlags(core.Qt__CustomizeWindowHint|core.Qt__WindowCloseButtonHint)
	this.loginDlg.Resize2(400,300)
	this.loginDlg.SetFixedSize(this.loginDlg.Geometry().Size())
	labelName := widgets.NewQLabel2("用户名:",this.loginDlg,0)
	labelName.SetGeometry2(110,107,50,20)

	this.lineEdit_UserName = widgets.NewQLineEdit(this.loginDlg)
	this.lineEdit_UserName.SetGeometry2(190,100,120,30)
	this.lineEdit_UserName.SetText(Config.Instance.GetUserName())

	loginBtn := widgets.NewQPushButton2("登录",this.loginDlg)
	loginBtn.SetGeometry2(130,190,140,30)
	loginBtn.ConnectClicked(this.on_LoginBtnClicked)
	this.loginDlg.Show()

	// 进入消息循环
	this.app.Exec()
}