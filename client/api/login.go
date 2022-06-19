package api

import (
	"bytes"
	"client/Model"
	"client/config"
	"encoding/json"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
)



func Login(userName string,passWord string)Model.LoginRet  {
	type loginRequest struct {
		UserName string `json:"userName"`
		PassWord string `json:"passWord"`
	}
	loginReq,err := json.Marshal(loginRequest{userName,passWord})
	if err != nil{
		return Model.LOGIN_INVALID
	}
	resp,err := http.Post(config.APIURL + "/login","application/json",bytes.NewReader(loginReq))
	if err != nil{
		return Model.LOGIN_INVALID
	}
	defer resp.Body.Close()
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return Model.LOGIN_INVALID
	}
	g := gjson.ParseBytes(respBytes)
	code := g.Get("code").Int()
	userRole := g.Get("userRole").Int()

	if code != 0{
		return Model.LOGIN_INVALID
	}
	return Model.LoginRet(userRole)
}