package TradeApp

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"os"
)

type TradeApp struct {
	app *widgets.QApplication
	wnd_Main *widgets.QMainWindow
	table_CSV *widgets.QTableWidget
	//统计页数
	label_CsvCount	*widgets.QLabel
	//上一页按钮
	btn_GoToLastPage *widgets.QPushButton
	//下一页按钮
	btn_GoToNextPage *widgets.QPushButton
}

func NewTradeApp()*TradeApp  {
	return &TradeApp{}
}

func (this *TradeApp)onMenu_ImportCSV(checked bool)  {
	filePath := widgets.QFileDialog_GetOpenFileName(this.wnd_Main,"导入客户数据","","表格文件(*.csv *.xlsx)","",widgets.QFileDialog__ReadOnly)
	if filePath == ""{
		return
	}
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
	this.table_CSV.SetShowGrid(false)
	this.table_CSV.SetSortingEnabled(false)

	gridLayout.AddWidget3(this.table_CSV,0,0,1,4,0)

	//添加空白项
	horizontalSpacer := widgets.NewQSpacerItem(200,20,widgets.QSizePolicy__Expanding,widgets.QSizePolicy__Minimum)
	gridLayout.AddItem(horizontalSpacer,1,0,1,1,0)

	this.label_CsvCount = widgets.NewQLabel2("共1页", centerWidget,0)
	gridLayout.AddWidget3(this.label_CsvCount,1,1,1,1,0)

	this.btn_GoToLastPage = widgets.NewQPushButton2("上一页",centerWidget)
	gridLayout.AddWidget3(this.btn_GoToLastPage,1,2,1,1,0)

	this.btn_GoToNextPage = widgets.NewQPushButton2("下一页",centerWidget)
	gridLayout.AddWidget3(this.btn_GoToNextPage,1,3,1,1,0)

	this.wnd_Main.SetCentralWidget(centerWidget)
}

func (this *TradeApp)initMainWindow()  {

	this.setupUi()

	// 设置窗口最小尺寸
	this.wnd_Main.SetMinimumSize2(800, 600)
	// 设置窗口默认全屏
	this.wnd_Main.SetWindowState(core.Qt__WindowMaximized)
	// 设置标题
	this.wnd_Main.SetWindowTitle("外贸客户管理平台")

	//添加菜单栏
	menuBar := this.wnd_Main.MenuBar()
	menu_App := menuBar.AddMenu2("程序")
	menu_Admin := menuBar.AddMenu2("管理员")
	menu_About := menuBar.AddMenu2("关于")

	menu_App.AddAction("清空当前列表").ConnectTriggered(this.onMenu_ClearCSV)
	menu_App.AddAction("退出程序").ConnectTriggered(this.onMenu_ExitApp)
	menu_Admin.AddAction("导入本地数据").ConnectTriggered(this.onMenu_ImportCSV)
	menu_Admin.AddAction("导出本地数据").ConnectTriggered(this.onMenu_ExportCSV)
	menu_About.AddAction("关于程序").ConnectTriggered(this.onMenu_AboutApp)

	this.table_CSV.SetColumnCount(24)
	this.table_CSV.SetHorizontalHeaderLabels([]string{
		"公司名称","成立年份","员工数量","年营业额(美元)","联系电话","客户邮箱",
		"联系人名字","Hunter数据","客户网址","客户类型","客户等级","最后跟进日期","备注","客户介绍",
		"Facebook主页","领英主页","推特账号","行业","国家","城市","详细地址","邮编","跟进人","导入日期"})

}

func (this *TradeApp)InitTradeApp()  {
	// 创建应用程序
	this.app = widgets.NewQApplication(len(os.Args), os.Args)

	this.initMainWindow()



}

func (this *TradeApp)Run()  {
	// 显示窗口
	this.wnd_Main.Show()
	// 进入消息循环
	this.app.Exec()
}