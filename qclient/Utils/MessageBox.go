package Utils

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func MessageBox_Info(parent widgets.QWidget_ITF,title string,content string)  {
	msgBox := widgets.NewQMessageBox2(widgets.QMessageBox__Information,
		title,content,widgets.QMessageBox__Yes,
		parent,core.Qt__Dialog)
	msgBox.Button(widgets.QMessageBox__Yes).SetText("好的")
	msgBox.Exec()
}

func MessageBox_Error(parent widgets.QWidget_ITF,title string,content string)  {
	msgBox := widgets.NewQMessageBox2(widgets.QMessageBox__Critical,
		title,content,widgets.QMessageBox__Yes,
		parent,core.Qt__Dialog)
	msgBox.Button(widgets.QMessageBox__Yes).SetText("好的")
	msgBox.Exec()
}