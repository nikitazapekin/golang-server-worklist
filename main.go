
 package main

import (
	"fmt"
	



	"log"
"server/db"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	r "server/route"
	"github.com/labstack/echo/v4/middleware"
)


func main() {
	fmt.Println("STATIC DIR:", "static")
	e := echo.New()
	db.Connect()
	e.Use(middleware.CORS())
	r.InitRoutes(e)
	err := e.Start(":5000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
 