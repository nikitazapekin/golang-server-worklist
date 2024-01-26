
package db
import (
	"database/sql"
	"fmt"
	//"io/ioutil"
//	"os"
	//"path/filepath"
	"log"
	_ "github.com/lib/pq"
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
    DB = db 
    CreateTable()
    InsertData("John Doe", "password123")
    InsertData("Тшллл", "password12ццацйа3")
    err = db.Ping()
    if err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }
}
func PingDB() error {
    err := DB.Ping()
    if err != nil {
        return err
    }
    return nil
}

func CreateTable() {
    // Ensure that the Connect function has been called before creating the table
    if DB == nil {
        log.Fatal("Database connection is not established. Call Connect function first.")
    }

    // SQL query to create the table
    query := `
        CREATE TABLE IF NOT EXISTS TESTTABLE (
            name VARCHAR(255),
            password VARCHAR(255)
        );
    `

    _, err := DB.Exec(query)
    if err != nil {
        log.Fatalf("Failed to create table: %v", err)
    }

    fmt.Println("Table TESTTABLE created successfully.")
}

func InsertData(name, password string) {
    // Ensure that the Connect function has been called before inserting data
    if DB == nil {
        log.Fatal("Database connection is not established. Call Connect function first.")
    }

    // SQL query to insert data into the table
    query := "INSERT INTO TESTTABLE (name, password) VALUES ($1, $2);"

    _, err := DB.Exec(query, name, password)
    if err != nil {
        log.Fatalf("Failed to insert data into TESTTABLE: %v", err)
    }

    fmt.Println("Data inserted into TESTTABLE successfully.")
}


/*
package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

var (
	DB *sql.DB
)

func Connect() {
	fmt.Println("DB WORK")

	dbConnStr := "user=Nikita password=Backend dbname=golang-database sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", dbConnStr)
	if err != nil {
		log.Fatalf("Failed to open database connection: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}
}

func CreateTables() error {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	filePath := filepath.Join(wd, "/migrations/users_schema.sql")
	sqlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

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

func CreateUsersTable() error {
	// SQL statement to create the table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS createdAtGolang (
		id SERIAL PRIMARY KEY,
		username VARCHAR(50) NOT NULL,
		password VARCHAR(50) NOT NULL
	);
	`

	_, err := DB.Exec(createTableSQL)
	if err != nil {
		return err
	}

	// Insert a sample user
	_, err = DB.Exec("INSERT INTO createdAtGolang (username, password) VALUES ('sampleUser', 'samplePassword')")
	if err != nil {
		return err
	}

	return nil
}
*/