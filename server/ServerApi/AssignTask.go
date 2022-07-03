package ServerApi

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"server/DataBase"
	"server/Model"
)

func API_AssignTask(c *gin.Context)  {
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
	var reqData Model.AssignTaskReq
	err = json.Unmarshal(respBody,&reqData)
	if err != nil{
		retMsg["code"] = 403
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	err = DataBase.Instance.AssignTask(reqData)
	if err != nil{
		retMsg["code"] = 404
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
	}
	retMsg["code"] = 0
	retMsg["msg"] = "ok"
	c.JSON(http.StatusOK, retMsg)
}

