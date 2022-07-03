package ServerApi

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"server/DataBase"
	"server/Model"
)

//不分用户,取全部数据

func API_GetAllCustomerData(c *gin.Context)  {

	retMsg := make(map[string]interface{})
	respBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		retMsg["code"] = 401
		retMsg["msg"] = "invalid request"
		c.JSON(http.StatusOK, retMsg)
		return
	}
	if c.Request.Header.Get("AUTH") != "123456789"{
		retMsg["code"] = 402
		retMsg["msg"] = "invalid auth"
		c.JSON(http.StatusOK, retMsg)
		return
	}
	type getAllCustomerDataReq struct {
		PageIndex int `json:"page_index"`
		PageCount int `json:"page_count"`
		OrderBy     string `json:"order_by"`
	}
	var reqData getAllCustomerDataReq
	err = json.Unmarshal(respBody,&reqData)
	if err != nil{
		retMsg["code"] = 403
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	from := (reqData.PageIndex-1) * reqData.PageCount
	customList,err := DataBase.Instance.QueryAllCustomerData(from,reqData.PageCount,reqData.OrderBy)
	if err != nil{
		retMsg["code"] = 404
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	rowsCount ,err := DataBase.Instance.GetAllDataCount()
	if err != nil{
		retMsg["code"] = 405
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	retMsg["code"] = 0
	retMsg["data"] = customList
	retMsg["count"] = rowsCount
	c.JSON(http.StatusOK, retMsg)
}

//只取分配的用户数据

func API_GetCustomerData(c *gin.Context)  {

	retMsg := make(map[string]interface{})
	respBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		retMsg["code"] = 401
		retMsg["msg"] = "invalid request"
		c.JSON(http.StatusOK, retMsg)
		return
	}
	type getCustomerDataReq struct {
		PageIndex int `json:"page_index"`
		PageCount int `json:"page_count"`
		OrderBy     string `json:"order_by"`
		UserName  string	`json:"user_name"`
	}
	var reqData getCustomerDataReq
	err = json.Unmarshal(respBody,&reqData)
	if err != nil{
		retMsg["code"] = 403
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	from := (reqData.PageIndex-1) * reqData.PageCount
	customList,err := DataBase.Instance.QueryCustomerData(from,reqData.PageCount,reqData.OrderBy,reqData.UserName)
	if err != nil{
		retMsg["code"] = 404
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	rowsCount ,err := DataBase.Instance.GetDataCount(reqData.UserName)
	if err != nil{
		retMsg["code"] = 405
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	retMsg["code"] = 0
	retMsg["data"] = customList
	retMsg["count"] = rowsCount
	c.JSON(http.StatusOK, retMsg)
}

//更新数据

func API_UpdateCustomerData(c *gin.Context)  {
	retMsg := make(map[string]interface{})
	respBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		retMsg["code"] = 401
		retMsg["msg"] = "invalid request"
		c.JSON(http.StatusOK, retMsg)
		return
	}
	if c.Request.Header.Get("AUTH") != "123456789"{
		retMsg["code"] = 402
		retMsg["msg"] = "invalid auth"
		c.JSON(http.StatusOK, retMsg)
		return
	}
	var reqData []Model.CustomerInformation
	err = json.Unmarshal(respBody,&reqData)
	if err != nil{
		retMsg["code"] = 403
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	err = DataBase.Instance.UpdateCustomerData(reqData)
	if err != nil{
		retMsg["code"] = 404
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
	}
	retMsg["code"] = 0
	retMsg["msg"] = "ok"
	c.JSON(http.StatusOK, retMsg)
}
