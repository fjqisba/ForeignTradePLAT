package api

import (
	"errors"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"qclient/Model"
	"qclient/global"
	"strings"
)



func GetUserLevel(userName string)(Model.UserLevel,error)  {
	resp,err := http.Post(global.GServerUrl + "/getUserLevel","plain/text",strings.NewReader(userName))
	if err != nil{
		return Model.USER_INVALID,err
	}
	defer resp.Body.Close()
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return Model.USER_INVALID,err
	}
	g := gjson.ParseBytes(respBytes)
	if g.Get("code").Int() != 0{
		return Model.USER_INVALID,errors.New(g.Get("msg").String())
	}
	level := int(g.Get("userLevel").Int())
	return Model.UserLevel(level),nil
}