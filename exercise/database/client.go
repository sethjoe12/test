package database

import (
	"log"
	"test/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

type Idb interface {
	DbConnectAndMigrate() (*gorm.DB, error)
}

type db struct{}

func DbConnectAndMigrate() (*gorm.DB, error) {
	Instance, dbError = gorm.Open(mysql.Open("root:sethjoe@tcp(0.0.0.0:3303)/test_demo?parseTime=true"), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")

	if dbError = Instance.AutoMigrate(&models.User{}); dbError != nil {
		log.Println(dbError)
	}

	return Instance, dbError
}

// func DbConnectAndMigrate() (*gorm.DB, error) {

// 	errorENV := godotenv.Load()
// 	if errorENV != nil {
// 		panic("Failed to load env file")
// 	}

// 	dbUser := os.Getenv("DB_USER")
// 	dbPass := os.Getenv("DB_PASS")
// 	dbHost := os.Getenv("DB_HOST")
// 	dbName := os.Getenv("DB_NAME")

// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)
// 	Instance, dbError = gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if dbError != nil {
// 		log.Fatal(dbError)
// 		panic("Cannot connect to DB")
// 	}
// 	log.Println("Connected to Database!")

// 	if dbError = Instance.AutoMigrate(&models.User{}); dbError != nil {
// 		log.Println(dbError)
// 	}

// 	return Instance, dbError
// }
