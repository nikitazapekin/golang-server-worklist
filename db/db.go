
package db

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"
	"net/http"
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
	if DB == nil {
		log.Fatal("Database connection is not established. Call Connect function first.")
	}
	query := `
	CREATE TABLE IF NOT EXISTS user_data (
		username VARCHAR(255),
		country VARCHAR(255),
		city VARCHAR(255),
		telephone VARCHAR(255),
		email VARCHAR(255),
		password VARCHAR(255),
		role VARCHAR(255),
		registration_data VARCHAR(255),
		avatar VARCHAR(255),
		document BYTEA,
		favourite_offers JSONB,
		experience VARCHAR(255),
		education VARCHAR(255),
		last_time_at_network VARCHAR(255),
		chats JSONB,
        user_id VARCHAR(255),
        describtion VARCHAR(255)
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table user_data created successfully.")
}
 func InsertData(w http.ResponseWriter, username string, password string, country, city string, telephone string, email string) error {
    if DB == nil {
        http.Error(w, "Database connection is not established. Call Connect function first.", http.StatusInternalServerError)
        return fmt.Errorf("Database connection is not established. Call Connect function first.")
    }
    registrationData := time.Now().Format("02-01-2006")
    document := ""
    query := "INSERT INTO user_data (username, country, city, telephone, email, password, role, registration_data, avatar, document, favourite_offers, experience, education, last_time_at_network, chats, describtion, user_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17 );" //$16
    role := "user"
    avatar := "https://cdn-icons-png.flaticon.com/512/1946/1946429.png"
    favouriteOffers := "{}"
    experience := ""
    education := ""
    lastTimeAtNetwork := ""
    chats := "{}"
    describtion := ""
    user_id := generateUniqueUserID();  
    _, err := DB.Exec(query, username, country, city, telephone, email, password, role,
        registrationData,//userID, 
        avatar, document, favouriteOffers, experience, education,
        lastTimeAtNetwork, chats,
        describtion, user_id,
    )
    if err != nil {
        fmt.Println("Error at auth")
        http.Error(w, fmt.Sprintf("Failed to insert data into user_data: %v", err), http.StatusInternalServerError)
        return fmt.Errorf("Failed to insert data into user_data: %v", err)
    }

    fmt.Println("Data inserted into user_data successfully.")
    return nil
}

    func generateUniqueUserID() string {
        timestamp := time.Now().Unix()
        randomNumber := rand.Intn(100000000) 
        return fmt.Sprintf("%d%d", timestamp, randomNumber)
    }


    type UserData struct {
        Username             string `json:"username"`
        Country              string `json:"country"`
        City                 string `json:"city"`
        Telephone            string `json:"telephone"`
        Email                string `json:"email"`
        Password             string `json:"password"`
        Role                 string `json:"role"`
        RegistrationData     string `json:"registration_data"`
        Avatar               string `json:"avatar"`
        Document             string `json:"document"`
        FavouriteOffers      string `json:"favourite_offers"`
        Experience           string `json:"experience"`
        Education            string `json:"education"`
        LastTimeAtNetwork    string `json:"last_time_at_network"`
        Chats                string `json:"chats"`
        Describtion   string `json:"describtion"`
       User_id               string `json:"user_id"`
    }


    func FindUserByUsername(username string) (UserData, error) {
        fmt.Println("EMAILLLLLLLLLLLLLLLLLL")
        fmt.Println(username)
        if DB == nil {
            return UserData{}, fmt.Errorf("Database connection is not established. Call Connect function first.")
        }
    
        query := "SELECT * FROM user_data WHERE username = $1"
        row := DB.QueryRow(query, username)
    
        var user UserData
        err := row.Scan(
            &user.Username,
            &user.Country,
            &user.City,
            &user.Telephone,
            &user.Email,
            &user.Password,
            &user.Role,
            &user.RegistrationData,
            &user.Avatar,
            &user.Document,
            &user.FavouriteOffers,
            &user.Experience,
            &user.Education,
            &user.LastTimeAtNetwork,
            &user.Chats,
            &user.Describtion,
            &user.User_id,
        )
    
        if err == sql.ErrNoRows {
            fmt.Println(user)
            return UserData{}, fmt.Errorf("User not found with email: %s", username)
        } else if err != nil {
            return UserData{}, fmt.Errorf("Failed to query user_data: %v", err)
        }
    fmt.Println(user)
        return user, nil
    }
    func FindUserByEmail(email string) (UserData, error) {
        fmt.Println("EMAILLLLLLLLLLLLLLLLLL")
        fmt.Println(email)
        if DB == nil {
            return UserData{}, fmt.Errorf("Database connection is not established. Call Connect function first.")
        }
    
        query := "SELECT * FROM user_data WHERE email = $1"
        row := DB.QueryRow(query, email)
    
        var user UserData
        err := row.Scan(
            &user.Username,
            &user.Country,
            &user.City,
            &user.Telephone,
            &user.Email,
            &user.Password,
            &user.Role,
            &user.RegistrationData,
            &user.Avatar,
            &user.Document,
            &user.FavouriteOffers,
            &user.Experience,
            &user.Education,
            &user.LastTimeAtNetwork,
            &user.Chats,
            &user.Describtion,
            &user.User_id,
        )
    
        if err == sql.ErrNoRows {
            return UserData{}, fmt.Errorf("User not found with email: %s", email)
        } else if err != nil {
            return UserData{}, fmt.Errorf("Failed to query user_data: %v", err)
        }
    fmt.Println(user)
        return user, nil
    }

    func UpdateUser(user UserData, currentUserEmail string) error {
        fmt.Println("User id db")
        fmt.Println(user)
        if DB == nil {
            return fmt.Errorf("Database connection is not established. Call Connect function first.")
        }
        //, about='cdwcdcw'
      /*  query := `
        UPDATE user_data
        SET country='gaaaay', city='Min',  telephone='3232', password='Bellll323',
            education='322323', experience='3232cwdw'
        WHERE email = 'wotblitz362@mail.ru';
    ` 
    _, err := DB.Exec(query) */
       query := `
            UPDATE user_data
            SET country=$2, city=$3, telephone=$4, password=$5,
                education=$6, describtion=$7, experience=$8
            WHERE email = $1
        ` 
        _, err := DB.Exec(query,
            currentUserEmail,
          //  user.Email, 
        //  "wotblitz362@mail.ru",
            user.Country, user.City, user.Telephone, user.Password,
            user.Education, user.Describtion, user.Experience,
           
        ) 
    
        if err != nil {
         //   fmt.Errorf("Failed to update user_data: %v", err)
           // return err
           return fmt.Errorf("Failed to update user_data: %v", err)

          //  return fmt.Errorf("Failed to update user_data: %v", err)
        }
    
        fmt.Println("User updated successfully.")
        return nil
    }