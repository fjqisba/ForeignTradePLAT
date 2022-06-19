package main

import (
	"client/XCApp"
	"syscall"
	"unsafe"
)

func callMsgBox()  {
	hUser32,_ := syscall.LoadDLL("user32.dll")
	callMessageBox, _ := hUser32.FindProc("MessageBoxW")
	text := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("我是内容")))
	title := uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr("我是标题")))
	callMessageBox.Call(0,text,title,0)
}

func main()  {
	xcApp := XCApp.NewXCApp()
	xcApp.InitXCApp()
	xcApp.Run()
}