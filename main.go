 
 package main

import (
	//"fmt"
	"log"
	//"net/http"
	"server/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	r "server/route"
)
func main() {
	e := echo.New()
	db.Connect()
	e.Use(middleware.CORS())
	r.InitRoutes(e)
 r.InitWebsocketRoutes(e)
	err := e.Start(":5000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}


	//handleHistory()
}
 




