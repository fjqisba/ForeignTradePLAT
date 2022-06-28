package main

import (
	"qclient/TradeApp"
)

func main()  {
	app := TradeApp.NewTradeAppClient()
	app.InitTradeApp()
	app.Run()
}
