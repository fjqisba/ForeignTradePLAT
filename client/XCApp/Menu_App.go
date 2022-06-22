package XCApp



//菜单_程序被选择

func (this *XCApp)OnMenuApp_Selected(nID int, pbHandled *bool)int  {

	//导入CSV
	if nID == 0{
		return this.importClientData()
	}
	//保存CSV
	if nID == 1{
		return this.ExportCSV()
	}
	//退出
	if nID == 2{
		this.wnd_Main.CloseWindow()
	}
	return 0
}