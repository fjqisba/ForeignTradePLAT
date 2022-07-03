package ServerApi

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"server/DataBase"
	"server/Model"
)

func API_Upload(c *gin.Context) {

	retMsg := make(map[string]interface{})
	respBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		retMsg["code"] = 401
		retMsg["msg"] = "invalid request"
		c.JSON(http.StatusOK, retMsg)
		return
	}

	var infoList []Model.CustomerInformation
	err = json.Unmarshal(respBody,&infoList)
	if err != nil{
		retMsg["code"] = 401
		retMsg["msg"] = "parse req error"
		c.JSON(http.StatusOK, retMsg)
		return
	}

	err = DataBase.Instance.ImportCustomerInfo(infoList)
	if err != nil{
		retMsg["code"] = 402
		retMsg["msg"] = err
		c.JSON(http.StatusOK, retMsg)
		return
	}

	retMsg["code"] = 0
	c.JSON(http.StatusOK, retMsg)
	return
}