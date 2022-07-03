package Config

import (
	"gopkg.in/ini.v1"
	"log"
	"os"
	"qclient/Utils"
	"strconv"
)

var(
	Instance AppConfig
)

type AppConfig struct {
	ini *ini.File
	filePath string
}

//保存列宽度

func (this *AppConfig)SaveItemWidth(iItem int, nWidth int)  {
	this.ini.Section("MainWindow").Key("Item"+strconv.Itoa(iItem)).SetValue(strconv.Itoa(nWidth))
	err := this.ini.SaveTo(this.filePath)
	if err != nil{
		log.Println("[SaveItemWidth]:",err)
	}
}

//获取列宽度

func (this *AppConfig)GetItemWidth(iItem int)int {
	i, err := strconv.Atoi(this.ini.Section("MainWindow").Key("Item"+strconv.Itoa(iItem)).Value())
	if err != nil{
		return 100
	}
	return i
}


//获取打开文件的路径

func (this *AppConfig)GetOpenFileDir()string  {
	return this.ini.Section("MainWindow").Key("OpenFileDir").Value()
}

//设置打开文件的路径

func (this *AppConfig)SetOpenFileDir(filePath string)  {
	this.ini.Section("MainWindow").Key("OpenFileDir").SetValue(filePath)
	this.ini.SaveTo(this.filePath)
}

func (this *AppConfig)SetUserName(UserName string)  {
	this.ini.Section("MainWindow").Key("UserName").SetValue(UserName)
	this.ini.SaveTo(this.filePath)
}

func (this *AppConfig)GetUserName()string  {
	return this.ini.Section("MainWindow").Key("UserName").Value()
}


//获取保存文件的路径

func (this *AppConfig)GetCreateFileDir()string  {
	return this.ini.Section("MainWindow").Key("CreateFileDir").Value()
}

//设置保存文件的路径

func (this *AppConfig)SetCreateFileDir(filePath string)  {
	this.ini.Section("MainWindow").Key("CreateFileDir").SetValue(filePath)
	this.ini.SaveTo(this.filePath)
}


func (this *AppConfig)GetHunterApKey()string  {
	return this.ini.Section("Main").Key("HunterApiKey").Value()
}

func (this *AppConfig)initConfigManager(settingPath string)error  {
	var err error
	this.ini,err = ini.Load(settingPath)
	if err != nil{
		return err
	}
	this.filePath = settingPath
	return nil
}

func init()  {
	settingPath := "./setting.ini"
	if Utils.IsPathExists(settingPath) == false{
		hFile,_ := os.Create(settingPath)
		if hFile != nil{
			hFile.Close()
		}
	}
	err := Instance.initConfigManager(settingPath)
	if err != nil{
		log.Panicln(err)
	}
}

func (this *AppConfig)SaveViewHunterItemWidth(iItem int, nWidth int)  {
	this.ini.Section("ViewHunter").Key("Item"+strconv.Itoa(iItem)).SetValue(strconv.Itoa(nWidth))
	err := this.ini.SaveTo(this.filePath)
	if err != nil{
		log.Println("[SaveItemWidth]:",err)
	}
}

//获取列宽度

func (this *AppConfig)GetViewHunterItemWidth(iItem int)int {
	i, err := strconv.Atoi(this.ini.Section("ViewHunter").Key("Item"+strconv.Itoa(iItem)).Value())
	if err != nil{
		return 100
	}
	return i
}