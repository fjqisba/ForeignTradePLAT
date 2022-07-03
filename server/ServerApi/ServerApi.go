package ServerApi

import (
	"github.com/gin-gonic/gin"
	"log"
)

const AuthToken = "123456789"

type ServerApi struct {
	g *gin.Engine
}

func NewServerApi()*ServerApi {
	gin.SetMode(gin.ReleaseMode)
	return &ServerApi{
		g: gin.New(),
	}
}

func (this *ServerApi)InitServerApi()error{

	//用户管理
	this.g.GET("/getUserList",API_GetUserList)
	this.g.POST("/deleteUser",API_DeleteUser)
	this.g.POST("/addUser",API_AddUser)
	this.g.POST("/getUserLevel",API_GetUserLevel)

	this.g.POST("/uploadCustomerData",API_Upload)
	this.g.POST("/getAllCustomerData",API_GetAllCustomerData)
	this.g.POST("/getCustomerData",API_GetCustomerData)
	this.g.POST("/updateCustomerData",API_UpdateCustomerData)

	//任务分配
	this.g.POST("/assignTask",API_AssignTask)
	return nil
}

func (this *ServerApi)RunServerApi()error{
	log.Println("Open ServerApi Port:8888")
	return this.g.Run("0.0.0.0:8888")
}