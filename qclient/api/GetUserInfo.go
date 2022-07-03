package api

import (
	"errors"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"qclient/Const"
	"strings"
)

type UserLevel int

const(
	//无效登录
	USER_INVALID UserLevel = 0
	//管理员登录
	USER_ADMIN UserLevel = 1
	//普通用户
	USER_USER UserLevel = 2
)

func GetUserLevel(userName string)(UserLevel,error)  {
	resp,err := http.Post(Const.ServerUrl + "/getUserLevel","plain/text",strings.NewReader(userName))
	if err != nil{
		return USER_INVALID,err
	}
	defer resp.Body.Close()
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return USER_INVALID,err
	}
	g := gjson.ParseBytes(respBytes)
	if g.Get("code").Int() != 0{
		return USER_INVALID,errors.New(g.Get("msg").String())
	}
	level := int(g.Get("userLevel").Int())
	return UserLevel(level),nil
}