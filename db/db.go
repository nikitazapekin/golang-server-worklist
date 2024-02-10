package db

import (
	"database/sql"
	"fmt"
	"log"
    "strings"
	"math/rand"
	"net/http"
    "github.com/lib/pq"
	_ "github.com/lib/pq"
	"time"
  // "strconv"

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
    CreateTableOfOffers()
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
		document VARCHAR(255),
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
type VacancyData struct {
    ID            int      `json:"id"`
    Title         string   `json:"title"`
    Description   string   `json:"description"`
    Skills        []string `json:"skills"`
    WorkingPerDay string   `json:"workingPerDay"`
    Location      string   `json:"location"`
    Salary        string   `json:"salary"`
    Owner         string   `json:"owner"`
    ImageSet      []string `json:"image_set"`
    DataOfPublication string `json:"data_of_publication"`
  //  Comments     []string `json:"comments"`
    LastTimeOfRise string `json:"last_time_of_rise"`
}

/* ПОСЛЕДНЯЯ ВЕРСИЯ ПОЛЯ vacancy_data
func CreateTableOfOffers() {
	if DB == nil {
		log.Fatal("Database connection is not established. Call Connect function first.")
	}
	query := `
	CREATE TABLE IF NOT EXISTS vacancy_data (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255),
        describtion VARCHAR(255),
        skills VARCHAR(255)[] ,
        workingPerDay VARCHAR(255),
        location VARCHAR(255),
           salary VARCHAR(255)
	);
	`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table user_data created successfully.")
}  */


func CreateTableOfOffers() {
	if DB == nil {
		log.Fatal("Database connection is not established. Call Connect function first.")
	}
	query := `
	CREATE TABLE IF NOT EXISTS vacancy_data (
        id SERIAL PRIMARY KEY,
        title VARCHAR(255),
        describtion VARCHAR(255),
        skills VARCHAR(255)[] ,
        workingPerDay VARCHAR(255),
        location VARCHAR(255),
           salary VARCHAR(255),
           owner VARCHAR(255),
           image_set VARCHAR(255)[] ,
           data_of_publication VARCHAR(255),
           last_time_of_rise VARCHAR(255)
           );
           `
         //  comments VARCHAR(255)[] ,
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}
	fmt.Println("Table user_data created successfully.")
}


 func GetAmountOfOffers() ([]VacancyData, error) {
    initDatabase()
    fmt.Println("INIT")
    if DB == nil {
        return nil, fmt.Errorf("Database connection is not established. Call Connect function first.")
    }

query := `
    SELECT id, title, describtion, skills, workingPerDay, location, salary
    FROM vacancy_data
`
    rows, err := DB.Query(query)
    if err != nil {
        return nil, fmt.Errorf("Failed to retrieve vacancy data: %v", err)
    }
    defer rows.Close()

    fmt.Println("rows")
    fmt.Println(rows)
amountOfOffers :=0;
    vacancyData := make([]VacancyData, 0)
    for rows.Next() {
        var vd VacancyData
        var skills string
        if err := rows.Scan(&vd.ID, &vd.Title, &vd.Description, &skills, &vd.WorkingPerDay, &vd.Location, &vd.Salary); err != nil {
            return nil, fmt.Errorf("Failed to scan vacancy data: %v", err)
        }
        vd.Skills = strings.Split(skills, ",")
        vacancyData = append(vacancyData, vd)
    }
    
 fmt.Println("FOUND ELEMS")
 fmt.Println(vacancyData)
 fmt.Println("FIRST")
 fmt.Println(vacancyData[0].Skills[0])
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error in rows: %v", err)
    }
    for i := range vacancyData {
        fmt.Println(i)
        amountOfOffers++;
    }
    return vacancyData, nil

   // return strconv.Itoa(amountOfOffers), nil
} 



/*
image_set VARCHAR(255)[] ,
data_of_publication VARCHAR(255),
comments VARCHAR(255)[] ,
last_time_of_rise VARCHAR(255)
 */


 func GetAllVacancyData(limit, page int) ([]VacancyData, error) {             
    initDatabase()
    fmt.Println("INIT")
    if DB == nil {
        return nil, fmt.Errorf("Database connection is not established. Call Connect function first.")
    }
    offset := limit * (page - 1)

    query := `
    SELECT id, title, describtion, skills, workingPerDay, location, salary, image_set, data_of_publication, owner
    FROM vacancy_data
    LIMIT $1 OFFSET $2
`

    rows, err := DB.Query(query, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("Failed to retrieve vacancy data: %v", err)
    }
    defer rows.Close()

    fmt.Println("rows")
    fmt.Println(rows)

    vacancyData := make([]VacancyData, 0)
    for rows.Next() {
        var vd VacancyData
        var skills string
        var images string
        if err := rows.Scan(&vd.ID, &vd.Title, &vd.Description, &skills, &vd.WorkingPerDay, &vd.Location, &vd.Salary, &images, &vd.DataOfPublication, &vd.Owner); err != nil {
            return nil, fmt.Errorf("Failed to scan vacancy data: %v", err)
        }
        vd.Skills = strings.Split(skills, ",")
        fmt.Println("IMAGEEEEEEEEEEEEEESS")
      /*  fmt.Println(&images)
        if(&images==nil){

        } */
     //   vd.ImageSet = strings.Split(images, ",")
     
        vacancyData = append(vacancyData, vd)
       // vacancyData = append(vacancyData, vd)
    } 
  /*  for rows.Next() {
        var vd VacancyData
        var skills string
        var images *string // Изменяем тип переменной на *string, чтобы проверить nil
    
        if err := rows.Scan(&vd.ID, &vd.Title, &vd.Description, &skills, &vd.WorkingPerDay, &vd.Location, &vd.Salary, &images, &vd.DataOfPublication, &vd.Owner); err != nil {
            return nil, fmt.Errorf("Failed to scan vacancy data: %v", err)
        }
    
        if images == nil { // Проверяем, является ли images nil
            vd.ImageSet = []string{} // Устанавливаем пустой слайс, если nil
        } else {
            vd.ImageSet = strings.Split(*images, ",") // Разбиваем строку, если images не nil
        }
    
        vd.Skills = strings.Split(skills, ",")
        vacancyData = append(vacancyData, vd)
    } */
    

    fmt.Println("FOUND ELEMS")
    fmt.Println(vacancyData)
    fmt.Println("FIRST")
    fmt.Println(vacancyData[0].Skills[0])
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error in rows: %v", err)
    }
    for i := range vacancyData {
        vacancyData[i].Description = strings.ReplaceAll(vacancyData[i].Description, `"`, "")
        for j := range vacancyData[i].Skills {
            vacancyData[i].Skills[j] = strings.ReplaceAll(vacancyData[i].Skills[j], `"`, "")
            vacancyData[i].Skills[j] = strings.ReplaceAll(vacancyData[i].Skills[j], `{`, "")
            vacancyData[i].Skills[j] = strings.ReplaceAll(vacancyData[i].Skills[j], `}`, "")
        }
    }

    return vacancyData, nil
}
 /*
 func GetAllVacancyData(limit, page int) ([]VacancyData, error) {                 CТАРРРРРРРРРРРРРРРРРРРРРРРРРРРРРРРРРРОООООООООООООООООООООООЕЕЕЕЕЕЕЕЕЕЕЕЕЕЕЕЕЕ
    initDatabase()
    fmt.Println("INIT")
    if DB == nil {
        return nil, fmt.Errorf("Database connection is not established. Call Connect function first.")
    }
    offset := limit * (page - 1)

    query := `
    SELECT id, title, describtion, skills, workingPerDay, location, salary
    FROM vacancy_data
    LIMIT $1 OFFSET $2
`

    rows, err := DB.Query(query, limit, offset)
    if err != nil {
        return nil, fmt.Errorf("Failed to retrieve vacancy data: %v", err)
    }
    defer rows.Close()

    fmt.Println("rows")
    fmt.Println(rows)

    vacancyData := make([]VacancyData, 0)
    for rows.Next() {
        var vd VacancyData
        var skills string
        if err := rows.Scan(&vd.ID, &vd.Title, &vd.Description, &skills, &vd.WorkingPerDay, &vd.Location, &vd.Salary); err != nil {
            return nil, fmt.Errorf("Failed to scan vacancy data: %v", err)
        }
        vd.Skills = strings.Split(skills, ",")
        vacancyData = append(vacancyData, vd)
    }

    fmt.Println("FOUND ELEMS")
    fmt.Println(vacancyData)
    fmt.Println("FIRST")
    fmt.Println(vacancyData[0].Skills[0])
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error in rows: %v", err)
    }
    for i := range vacancyData {
        vacancyData[i].Description = strings.ReplaceAll(vacancyData[i].Description, `"`, "")
        for j := range vacancyData[i].Skills {
            vacancyData[i].Skills[j] = strings.ReplaceAll(vacancyData[i].Skills[j], `"`, "")
            vacancyData[i].Skills[j] = strings.ReplaceAll(vacancyData[i].Skills[j], `{`, "")
            vacancyData[i].Skills[j] = strings.ReplaceAll(vacancyData[i].Skills[j], `}`, "")
        }
    }

    return vacancyData, nil
} */
/*
  owner VARCHAR(255),
           image_set VARCHAR(255)[] ,
           data_of_publication VARCHAR(255),
          comments VARCHAR(255)[] ,
          last_time_of_rise VARCHAR(255)
*/

func InsertDataIntoOffers(w http.ResponseWriter, title string, describtion string, skills []string, workingPerDay string, location string, salary string,  token string, arrayOfPictures []string) error {
    now := time.Now()
    dateString := now.Format("02,01,2006")
    fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA")
    fmt.Println(arrayOfPictures)

if DB == nil {
    http.Error(w, "Database connection is not established. Call Connect function first.", http.StatusInternalServerError)
    return fmt.Errorf("Database connection is not established. Call Connect function first.")
}
query := "INSERT INTO vacancy_data   (title, describtion, skills,  workingPerDay, location, salary, owner, data_of_publication, last_time_of_rise, image_set) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10);" //$16


_, err := DB.Exec(query,
title, describtion,// skills,
pq.Array(skills),
workingPerDay, location, salary,
token, dateString, dateString,
pq.Array(arrayOfPictures),
)
if err != nil {
fmt.Println("Error at auth")
http.Error(w, fmt.Sprintf("Failed to insert data into user_data: %v", err), http.StatusInternalServerError)
return fmt.Errorf("Failed to insert data into user_data: %v", err)
}

fmt.Println("Data inserted into user_data successfully.") 
return nil

}
//СТАРАЯ ВЕРСИЯ
  /*  func InsertDataIntoOffers(w http.ResponseWriter, title string, describtion string, skills []string, workingPerDay string, location string, salary string,  token string) error {
        now := time.Now()
        dateString := now.Format("02,01,2006")
        fmt.Println("TIME NOW" +dateString)

    if DB == nil {
        http.Error(w, "Database connection is not established. Call Connect function first.", http.StatusInternalServerError)
        return fmt.Errorf("Database connection is not established. Call Connect function first.")
    }
    query := "INSERT INTO vacancy_data   (title, describtion, skills,  workingPerDay, location, salary) VALUES ($1, $2, $3, $4, $5, $6);" //$16
   
   
    _, err := DB.Exec(query,
title, describtion,// skills,
pq.Array(skills),
 workingPerDay, location, salary,
)
if err != nil {
    fmt.Println("Error at auth")
    http.Error(w, fmt.Sprintf("Failed to insert data into user_data: %v", err), http.StatusInternalServerError)
    return fmt.Errorf("Failed to insert data into user_data: %v", err)
}

fmt.Println("Data inserted into user_data successfully.") 
return nil
  
}
 */


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
       // Avatar []byte `json:"avatar"`
       Avatar              string `json:"avatar"`
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

    func FindOfferById(id string) (VacancyData, error) {
        fmt.Println("EMAILLLLLLLLLLLLLLLLLL")
        fmt.Println(id)
        if DB == nil {
            return VacancyData{}, fmt.Errorf("Database connection is not established. Call Connect function first.")
        }
    
        query := "SELECT * FROM vacancy_data WHERE id = $1"
        row := DB.QueryRow(query, id)
     var vacancy VacancyData

     var skills string
     var images string
err := row.Scan(
    &vacancy.ID,
    &vacancy.Title,
    &vacancy.Description,
    &skills, 
    &vacancy.WorkingPerDay,
    &vacancy.Location,
    &vacancy.Salary,
    &vacancy.Owner,
    &images,
    &vacancy.DataOfPublication,
    &vacancy.LastTimeOfRise,
)

if err != nil {
    return VacancyData{}, fmt.Errorf("Failed to scan data: %v", err)
}

vacancy.Skills = strings.Split(skills, ",")
vacancy.ImageSet = strings.Split(images, ",")
        if err == sql.ErrNoRows {
            return VacancyData{}, fmt.Errorf("User not found with email: %s", id)
        } else if err != nil {
            return VacancyData{}, fmt.Errorf("Failed to query user_data: %v", err)
        }
    fmt.Println(vacancy)
        return vacancy, nil
    }

    func UpdateUser(user UserData, currentUserEmail string) error {
        fmt.Println("User id db")
        fmt.Println("current email" +currentUserEmail)
        fmt.Println("Tel " +user.Telephone)
        fmt.Println("doc" +user.Document)
        if DB == nil {
            return fmt.Errorf("Database connection is not established. Call Connect function first.")
        }
       query := `
            UPDATE user_data
            SET country=$2, city=$3, telephone=$4, password=$5,
                education=$6, describtion=$7, experience=$8,  document=$9, email=$10
            WHERE email = $1
        ` 
        _, err := DB.Exec(query,
            currentUserEmail,
            user.Country, user.City, user.Telephone, user.Password,
            user.Education, user.Describtion, user.Experience, user.Document, user.Email,
           
        ) 
    
        if err != nil {
           return fmt.Errorf("Failed to update user_data: %v", err)
        }
    
        fmt.Println("User updated successfully.")
        return nil
    }
    func UpdateUserAvatar(user UserData, newAvatar string) error {
        fmt.Println("User id db")
    //    fmt.Println("current email" +currentUserEmail)
        fmt.Println("Tel " +user.Telephone)
        fmt.Println("doc" +user.Document)
        if DB == nil {
            return fmt.Errorf("Database connection is not established. Call Connect function first.")
        }
       query := `
            UPDATE user_data
            SET avatar=$2
            WHERE email = $1
        ` 
        _, err := DB.Exec(query,
          user.Email,
          newAvatar,
        ) 
        if err != nil {
           return fmt.Errorf("Failed to update user_data: %v", err)
        }
    
        fmt.Println("User updated successfully.")
        return nil
    }
    const createTableQuery = `
CREATE TABLE IF NOT EXISTS user_logos (
    id SERIAL PRIMARY KEY,
    user_email VARCHAR(255),
    logo_url VARCHAR(255)
);
`

func initDatabase() error {
    _, err := DB.Exec(createTableQuery)
    if err != nil {
        return fmt.Errorf("Failed to create user_logos table: %v", err)
    }

    return nil
}

func AddLogo(logoURL, userEmail string) error {
    initDatabase()
    fmt.Println("INIT")
    if DB == nil {
        return fmt.Errorf("Database connection is not established. Call Connect function first.")
    }

    query := `
        INSERT INTO user_logos (user_email, logo_url)
        VALUES ($1, $2)
    `

    _, err := DB.Exec(query, userEmail, logoURL)
    if err != nil {
        return fmt.Errorf("Failed to insert logo record: %v", err)
    }

    fmt.Println("Logo URL added successfully.")
    return nil
}