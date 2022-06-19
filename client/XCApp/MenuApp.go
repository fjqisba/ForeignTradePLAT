package XCApp



//菜单_程序被选择

func (this *XCApp)OnMenuApp_Selected(nID int, pbHandled *bool)int  {

	//导入CSV
	if nID == 0{
		return this.importClientData()
	}

	if nID == 3{
		this.wnd_Main.CloseWindow()
	}

	return 0
}