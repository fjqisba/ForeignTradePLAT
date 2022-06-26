package TradeApp

import (
	"encoding/csv"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
	"log"
	"os"
)

func (this *TradeApp)onMenu_ClearCSV(checked bool)  {

	msgBox := widgets.NewQMessageBox2(widgets.QMessageBox__Question,
		"清空列表","您确定要清空当前列表中的数据吗?",widgets.QMessageBox__Yes|widgets.QMessageBox__No,
		this.wnd_Main,core.Qt__Dialog)
	msgBox.Button(widgets.QMessageBox__Yes).SetText("是")
	msgBox.Button(widgets.QMessageBox__No).SetText("否")
	if msgBox.Exec() == int(widgets.QMessageBox__Yes){
		this.table_CSV.SetRowCount(0)
		this.table_CSV.ClearContents()
	}
	return
}

func (this *TradeApp)onMenu_ExportCSV(checked bool)  {

	rowCount := this.table_CSV.RowCount()
	if rowCount == 0{
		return
	}
	filePath := widgets.QFileDialog_GetSaveFileName(this.wnd_Main,"导入客户数据","","表格文件(*.csv *.xlsx)","",widgets.QFileDialog__ReadOnly)
	if filePath == ""{
		return
	}
	//开始写出Csv
	hFile,err := os.OpenFile(filePath,os.O_WRONLY|os.O_CREATE|os.O_TRUNC,0666)
	if err != nil{
		return
	}
	defer hFile.Close()
	hCsvWriter := csv.NewWriter(hFile)
	hCsvWriter.Write([]string{"Company Name","Year Founded","Number of Employees","Annual Revenue",
		"Phone Number","Email", "Key Contact","Hunter Data","Company Domain Name","Customer Type",
		"Customer Rating","Followup Date","Remarks","Description","Facebook Page", "LinkedIn Company Page",
		"Twitter Handle","Industry","Country","City","Street Address","Postal Code","Coordinator","Creation Date"})

	for i:=0;i<rowCount;i++{
		log.Println(this.table_CSV.Item(i,0).Text())
		hCsvWriter.Write([]string{
			this.table_CSV.Item(i,0).Text(),this.table_CSV.Item(i,1).Text(),
			this.table_CSV.Item(i,2).Text(),this.table_CSV.Item(i,3).Text(),
			this.table_CSV.Item(i,4).Text(),this.table_CSV.Item(i,5).Text(),
			this.table_CSV.Item(i,6).Text(),this.table_CSV.Item(i,7).Text(),
			this.table_CSV.Item(i,8).Text(),this.table_CSV.Item(i,9).Text(),
			this.table_CSV.Item(i,10).Text(),this.table_CSV.Item(i,11).Text(),
			this.table_CSV.Item(i,12).Text(),this.table_CSV.Item(i,13).Text(),
			this.table_CSV.Item(i,14).Text(),this.table_CSV.Item(i,15).Text(),
			this.table_CSV.Item(i,16).Text(),this.table_CSV.Item(i,17).Text(),
			this.table_CSV.Item(i,18).Text(),this.table_CSV.Item(i,19).Text(),
			this.table_CSV.Item(i,20).Text(),this.table_CSV.Item(i,21).Text(),
			this.table_CSV.Item(i,22).Text(),this.table_CSV.Item(i,23).Text(),
		})
	}
	hCsvWriter.Flush()
}