package db

import (
	"database/sql"
	"fmt"
	//	"fmt"
	//"os"
	//	"time"
	"io/ioutil"
	"os"
    "path/filepath"
	"log"

	_ "github.com/lib/pq"
	"github.com/go-pg/migrations/v8"
)

var (
	DB *sql.DB
)

func Connect() {
fmt.Println("DB WORK")

	dbConnStr := "user=Nikita password=Backend dbname=golang-database sslmode=disable"
	db, err := sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}
	defer db.Close()


	oldVersion, newVersion, err := migrations.Run(db, "up")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Migrated from version %d to %d\n", oldVersion, newVersion)

	err = db.Ping()
	//err = CreateTables()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
	/*_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS example_table (
			id SERIAL PRIMARY KEY,
			test VARCHAR
		)
	`)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	_, err = db.Exec("INSERT INTO example_table (test) VALUES ($1)", "Hello, World!")
	if err != nil {
		log.Fatalf("Failed to insert data into table: %v", err)
	}
*/
}
/*
func CreateTables() error {
    // Read SQL file content
    sqlFile, err := ioutil.ReadFile("server/migrations/users_schema")
    if err != nil {
        return err
    }

    // Execute SQL statements
    _, err = DB.Exec(string(sqlFile))
    if err != nil {
        return err
    }

    return nil
} */


func CreateTables() error {
    // Get the current working directory
    wd, err := os.Getwd()
    if err != nil {
        return err
    }

 
   // filePath := filepath.Join(wd, "C:\Users\wotbl\fullstackworklist\server\server\migrations\users_schema.sql")
	//	"server/migrations/users_schema.sql")
	filePath := filepath.Join(wd, "/migrations/users_schema.sql")
    // Read SQL file content C:/Users/wotbl/fullstackworklist/
    sqlFile, err := ioutil.ReadFile(filePath)
    if err != nil {
        return err
    }

    // Execute SQL statements
    _, err = DB.Exec(string(sqlFile))
    if err != nil {
        return err
    }

    return nil
}

func PingDB() error {
	err := DB.Ping()
	if err != nil {
		return err
	}
	return nil
}  