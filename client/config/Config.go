package config

import (
	"client/Utils"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"strconv"
)

var(
	Instance AppConfig
)

type AppConfig struct {
	ini *ini.File
	filePath string
}

func (this *AppConfig)SaveItemWidth(iItem int, nWidth int)  {
	this.ini.Section("MainWindow").Key("Item"+strconv.Itoa(iItem)).SetValue(strconv.Itoa(nWidth))
	this.ini.SaveTo(this.filePath)
}

func (this *AppConfig)GetItemWidth(iItem int)int {
	i, err :=strconv.Atoi(this.ini.Section("MainWindow").Key("Item"+strconv.Itoa(iItem)).Value())
	if err != nil{
		return 0
	}
	return i
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
