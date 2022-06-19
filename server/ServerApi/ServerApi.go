package ServerApi

import (
	"github.com/gin-gonic/gin"
	"log"
)

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

	this.g.POST("/login", API_Login)



	return nil
}

func (this *ServerApi)RunServerApi()error{
	log.Println("Open ServerApi Port:8888")
	return this.g.Run("0.0.0.0:8888")
}