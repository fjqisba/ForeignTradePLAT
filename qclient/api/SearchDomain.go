package api

import (
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"qclient/Config"
	"qclient/Model"
)

func SearchDomain(domain string)(retEmailList []Model.EmailInfo)  {
	resp,err := http.Get("https://api.hunter.io/v2/domain-search?domain=" + domain + "&api_key=" + Config.Instance.GetHunterApKey())
	if err != nil{
		return retEmailList
	}
	defer resp.Body.Close()
	respBytes,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		return retEmailList
	}
	g := gjson.ParseBytes(respBytes)
	eMailList := g.Get("data").Get("emails").Array()
	for _,eEmailInfo := range eMailList{
		var tmpEmailInfo Model.EmailInfo
		tmpEmailInfo.Email = eEmailInfo.Get("value").String()
		tmpEmailInfo.Type = eEmailInfo.Get("type").String()
		tmpEmailInfo.PersonName = eEmailInfo.Get("first_name").String() + " " + eEmailInfo.Get("last_name").String()
		tmpEmailInfo.PhoneNumber = eEmailInfo.Get("phone_number").String()
		tmpEmailInfo.Department = eEmailInfo.Get("department").String()
		tmpEmailInfo.Position = eEmailInfo.Get("position").String()
		tmpEmailInfo.Linkedin = eEmailInfo.Get("linkedin").String()
		tmpEmailInfo.Twitter = eEmailInfo.Get("twitter").String()
		retEmailList = append(retEmailList, tmpEmailInfo)
	}
	return retEmailList
}