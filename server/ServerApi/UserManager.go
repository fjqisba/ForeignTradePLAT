package ServerApi

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"server/DataBase"
)

//添加新用户

func API_AddUser(c *gin.Context) {

	var retMsg map[string]interface{} = make(map[string]interface{})
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
	err = DataBase.Instance.AddUser(string(respBody))
	if err != nil{
		retMsg["code"] = 403
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	retMsg["code"] = 0
	return
}


//删除用户

func API_DeleteUser(c *gin.Context) {
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
	err = DataBase.Instance.DeleteUser(string(respBody))
	if err != nil{
		retMsg["code"] = 403
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	retMsg["code"] = 0
	return
}

func API_GetUserLevel(c *gin.Context) {

	retMsg := make(map[string]interface{})
	respBody, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		retMsg["code"] = 401
		retMsg["msg"] = "invalid request"
		c.JSON(http.StatusOK, retMsg)
		return
	}
	userLevel,err := DataBase.Instance.GetUserLevel(string(respBody))
	if err != nil{
		retMsg["code"] = 402
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	retMsg["code"] = 0
	retMsg["userLevel"] = userLevel
	c.JSON(http.StatusOK, retMsg)
	return
}

//获取用户列表

func API_GetUserList(c *gin.Context) {
	retMsg := make(map[string]interface{})
	if c.Request.Header.Get("AUTH") != AuthToken{
		retMsg["code"] = 401
		retMsg["msg"] = "invalid auth"
		c.JSON(http.StatusOK, retMsg)
		return
	}
	userList,err := DataBase.Instance.GetUserList()
	if err != nil{
		retMsg["code"] = 402
		retMsg["msg"] = err.Error()
		c.JSON(http.StatusOK, retMsg)
		return
	}
	retMsg["code"] = 0
	retMsg["userList"] = userList
	c.JSON(http.StatusOK, retMsg)
	return
}