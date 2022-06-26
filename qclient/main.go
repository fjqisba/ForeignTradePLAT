package main

import (
	"qclient/TradeApp"
)

func main()  {
	app := TradeApp.NewTradeApp()
	app.InitTradeApp()
	app.Run()
}
