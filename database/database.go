package database

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"os"
	"time"
)

var DB *sql.DB

func InitDatabase() {
	errENV := godotenv.Load()
	if errENV != nil {
		fmt.Println("Gagal loading .env file:", errENV)
	}
	var User = os.Getenv("DATABASE_USERNAME")
	var Password = os.Getenv("DATABASE_PASSWORD")
	var Network = os.Getenv("DATABASE_NETWORK")
	var Address = os.Getenv("DATABASE_ADDRESS")
	var DBName = os.Getenv("DATABASE_NAME")
	config := mysql.Config{
		User:   User,
		Passwd: Password,
		Net:    Network,
		Addr:   Address,
		DBName: DBName,
	}

	var err error
	DB, err = sql.Open("mysql", config.FormatDSN())

	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(20)
	DB.SetConnMaxLifetime(60 * time.Minute)
	DB.SetConnMaxIdleTime(60 * time.Minute)

	if err != nil {
		log.Fatalln("Database Connection err => ", err)
	}
	pingError := DB.Ping()
	if pingError != nil {
		log.Fatal(pingError)
	}
	fmt.Println(err)
	fmt.Println(pingError)
	fmt.Println("Connect!")

	// migrate -database "mysql://root@tcp(localhost:3306)/sharing_vision" -path database/migrations up
}

func init() {
	InitDatabase()
}
