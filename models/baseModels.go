package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbURI)

	conn, err := gorm.Open("postgres", dbURI)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	pingErr := db.DB().Ping()
	if pingErr != nil {
		fmt.Println(pingErr)
	} else {
		fmt.Println("Connected")
	}
}

// GetDB to get the DataBase
func GetDB() *gorm.DB {
	return db
}

// AccountsModel to get the DataBase
type AccountsModel struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Token    string `json:"token"`
}

// GetTableDataByID to get the DataBase
func GetTableDataByID(table string, id string) *AccountsModel {
	data := &AccountsModel{}
	err := GetDB().Table(table).Where("id = ?", id).First(data).Error
	if err != nil {
		fmt.Println("Error Connecting DB.", err)
	}
	data.Password = ""
	return data
}
