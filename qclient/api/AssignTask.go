package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"qclient/Model"
	"qclient/global"
)

func AssignTask(assignInfo Model.AssignTaskReq)error  {
	reqBytes,err := json.Marshal(assignInfo)
	if err != nil{
		return err
	}
	hReq,err := http.NewRequest("POST",global.GServerUrl + "/assignTask",bytes.NewReader(reqBytes))
	if err != nil{
		return err
	}
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
