package api

import (
	"errors"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"qclient/Const"
	"strings"
)

func AddUser(userName string)error  {
	hReq,err := http.NewRequest("POST",Const.ServerUrl +"/addUser",strings.NewReader(userName))
	hReq.Header.Set("AUTH","123456789")
	resp,err := http.DefaultClient.Do(hReq)
	if err != nil{
		return err
	}
	defer resp.Body.Close()
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}
	g := gjson.ParseBytes(respBytes)
	if g.Get("code").Int() != 0{
		return errors.New(g.Get("msg").String())
	}
	return nil
}

func DeleteUser(userName string)error  {
	hReq,err := http.NewRequest("POST",Const.ServerUrl +"/deleteUser",strings.NewReader(userName))
	hReq.Header.Set("AUTH","123456789")
	resp,err := http.DefaultClient.Do(hReq)
	if err != nil{
		return err
	}
	defer resp.Body.Close()
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return err
	}
	if resp.StatusCode != 200{
		return errors.New(string(respBytes))
	}
	g := gjson.ParseBytes(respBytes)
	if g.Get("code").Int() != 0{
		return errors.New(g.Get("msg").String())
	}
	return nil
}

//获取普通用户列表

func GetUserList()([]string,error)  {
	hReq,err := http.NewRequest("GET",Const.ServerUrl + "/getUserList",nil)
	if err != nil{
		return nil,err
	}
	hReq.Header.Set("AUTH","123456789")
	resp,err := http.DefaultClient.Do(hReq)
	if err != nil{
		return nil,err
	}
	defer resp.Body.Close()
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return nil,err
	}
	if resp.StatusCode != 200{
		return nil,errors.New(string(respBytes))
	}
	g := gjson.ParseBytes(respBytes)
	if g.Get("code").Int() != 0{
		return nil,errors.New(g.Get("msg").String())
	}
	userList := g.Get("userList").Array()
	var retUserList []string
	for _,eUser := range userList{
		retUserList = append(retUserList, eUser.String())
	}
	return retUserList,nil
}