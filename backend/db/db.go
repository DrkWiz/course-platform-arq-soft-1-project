package db

import (
	categoryClient "backend/clients/category"
	courseClient "backend/clients/course"
	fileClient "backend/clients/files"
	userClient "backend/clients/users"
	"fmt"
	"os"

	categoryModel "backend/model/category"
	courseModel "backend/model/courses"
	fileModel "backend/model/files"
	userModel "backend/model/users"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	// DB Connections Parameters
	DBName := os.Getenv("DB_NAME")
	DBUser := os.Getenv("DB_USER")
	DBPass := os.Getenv("DB_PASSWORD")
	DBHost := os.Getenv("DB_HOST")
	DBPort := os.Getenv("DB_PORT")

	if DBName == "" || DBUser == "" || DBHost == "" || DBPort == "" {
		log.Fatal("Database connection parameters missing")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", DBUser, DBPass, DBHost, DBPort, DBName)
	db, err = gorm.Open("mysql", dsn)

	if err != nil {
		log.Error("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all Clients that we build
	courseClient.Db = db
	userClient.Db = db
	categoryClient.Db = db
	fileClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&categoryModel.Category{})
	db.AutoMigrate(&courseModel.Course{})
	db.AutoMigrate(&courseModel.CourseCategory{})
	db.AutoMigrate(&userModel.UserCourses{})
	db.AutoMigrate(&fileModel.File{})

	log.Info("Finishing Migration Database Tables")
}
