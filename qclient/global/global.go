package global

import (
	"qclient/Config"
	"qclient/Model"
)

//存储全局变量

var(
	GServerUrl string
	//用户的等级
	GUserLevel Model.UserLevel
	//用户名称
	GUserName string
	//排序依据
	GOrderBy = "followup_date"
	//当前页面
	GCurrentPageIndex = 1
	//总页数
	GTotalPageCount = 1
	//当前用户列表
	GUserList []string
)

func init()  {
	GServerUrl = Config.Instance.GetServerAddr()
}