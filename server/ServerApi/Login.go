package ServerApi

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"server/DataBase"
)

func API_Login(c *gin.Context) {

	var retMsg map[string]interface{} = make(map[string]interface{})
	type loginRequest struct {
		UserName string `json:"userName"`
		PassWord string `json:"passWord"`
	}

	respBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		retMsg["code"] = 401
		retMsg["msg"] = "invalid request"
		c.JSON(http.StatusOK, retMsg)
		return
	}

	var loginReq loginRequest
	err = json.Unmarshal(respBody,&loginReq)
	if err != nil{
		retMsg["code"] = 401
		retMsg["msg"] = "parse req error"
		c.JSON(http.StatusOK, retMsg)
		return
	}

	userRole := DataBase.Instance.GetUserRole(loginReq.UserName,loginReq.PassWord)
	retMsg["code"] = 0
	retMsg["userRole"] = userRole
	c.JSON(http.StatusOK, retMsg)
	return
}