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

func GetAllCustomerData(page int,count int,orderBy string)(Model.CustomerDataRsp,error){
	var retRsp Model.CustomerDataRsp
	type getAllCustomerDataReq struct {
		PageIndex int `json:"page_index"`
		PageCount int `json:"page_count"`
		OrderBy     string `json:"order_by"`
	}
	reqBytes,err := json.Marshal(getAllCustomerDataReq{PageIndex: page,PageCount: count,OrderBy:orderBy})
	if err != nil{
		return retRsp,err
	}
	hReq,err := http.NewRequest("POST",global.GServerUrl + "/getAllCustomerData",bytes.NewReader(reqBytes))
	if err != nil{
		return retRsp,err
	}
	hReq.Header.Set("AUTH","123456789")
	resp,err := http.DefaultClient.Do(hReq)
	if err != nil{
		return retRsp,err
	}
	defer resp.Body.Close()
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return retRsp,err
	}
	err = json.Unmarshal(respBytes,&retRsp)
	if err != nil{
		return retRsp,err
	}
	if retRsp.Code != 0{
		return retRsp,errors.New(retRsp.Msg)
	}
	return retRsp,nil
}

func GetCustomerData(page int,count int,orderBy string,userName string)(Model.CustomerDataRsp,error){
	var retRsp Model.CustomerDataRsp
	type getAllCustomerDataReq struct {
		PageIndex int `json:"page_index"`
		PageCount int `json:"page_count"`
		OrderBy     string `json:"order_by"`
		UserName	string	`json:"user_name"`
	}
	reqBytes,err := json.Marshal(getAllCustomerDataReq{PageIndex: page,PageCount: count,OrderBy:orderBy,UserName: userName})
	if err != nil{
		return retRsp,err
	}
	hReq,err := http.NewRequest("POST",global.GServerUrl + "/getCustomerData",bytes.NewReader(reqBytes))
	if err != nil{
		return retRsp,err
	}
	resp,err := http.DefaultClient.Do(hReq)
	if err != nil{
		return retRsp,err
	}
	defer resp.Body.Close()
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return retRsp,err
	}
	err = json.Unmarshal(respBytes,&retRsp)
	if err != nil{
		return retRsp,err
	}
	if retRsp.Code != 0{
		return retRsp,errors.New(retRsp.Msg)
	}
	return retRsp,nil
}


func UpdateCustomerData(updateList []Model.CustomerInformation)error{
	reqBytes,err := json.Marshal(updateList)
	if err != nil{
		return err
	}
	hReq,err := http.NewRequest("POST",global.GServerUrl + "/updateCustomerData",bytes.NewReader(reqBytes))
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