


package main

import (
//	"database/sql"
	"log"
"server/db"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	r "server/route"
	"github.com/labstack/echo/v4/middleware"
)
func main() {
	e := echo.New()
	db.Connect()
	e.Use(middleware.CORS())
	r.InitRoutes(e)

	err := e.Start(":5000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}






//Nikita Backend
 


//https://github.com/adispartadev/golang-mvc-rest-api/blob/master/main.go
//https://github.com/josephspurrier/gowebapp/blob/master/README.md

 //go mod init server    
 //go get -u github.com/labstack/echo/v4
//go get github.com/labstack/echo/v4/middleware@v4.11.4
//go run server
//go get github.com/fsnotify/fsnotify
//go get github.com/lib/pq
//go get -u github.com/go-pg/migrations/v8

//