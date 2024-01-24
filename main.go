/* 
package main

import (
	"log"
	"server/db"
	//"net/http"
	r "server/route"
	//"server/route"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	db.Connect()
	// Инициализация маршрутов
	r.InitRoutes(e)

	// Запуск сервера на порту 8080
	err := e.Start(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}   */



/*
package main

import (
	"database/sql"
	"log"
	//"os"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	r "server/route"
)

func main() {
	// Database connection string
	dbConnStr := "user=postgres password=Belorus2010 dbname=golang-database sslmode=disable"

	// Create a database connection
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Echo instance
	e := echo.New()

	// Pass the database connection to the route initialization
	r.InitRoutes(e)

	// Start server
	err = e.Start(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
} */



package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
	r "server/route"
)

func main() {
	// Database connection string
	dbConnStr := "user=postgres password=Belorus2010 dbname=golang-database sslmode=disable"

	// Create a database connection
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()

	// Check the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Create a new table "example_table" with a "test" field
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS example_table (
			id SERIAL PRIMARY KEY,
			test VARCHAR
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	// Insert some data into the "example_table"
	_, err = db.Exec("INSERT INTO example_table (test) VALUES ($1)", "Hello, World!")
	if err != nil {
		log.Fatalf("Failed to insert data into table: %v", err)
	}

	// Echo instance
	e := echo.New()

	// Pass the database connection to the route initialization
	r.InitRoutes(e)

	// Start server
	//err = e.Start(":8080")
	err = e.Start(":8000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}







 


//https://github.com/adispartadev/golang-mvc-rest-api/blob/master/main.go
//https://github.com/josephspurrier/gowebapp/blob/master/README.md

 //go mod init server    
 //go get -u github.com/labstack/echo/v4
//go get github.com/labstack/echo/v4/middleware@v4.11.4
//go run server
//go get github.com/fsnotify/fsnotify
//go get github.com/lib/pq
