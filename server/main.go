package main

import (
	"log"
	"server/ServerApi"
)

func main()  {
	apiServer := ServerApi.NewServerApi()
	err := apiServer.InitServerApi()
	if err != nil{
		log.Panicln(err)
	}
	err = apiServer.RunServerApi()
	if err != nil{
		log.Panicln(err)
	}
}
