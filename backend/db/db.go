package db

import (
	categoryClient "backend/clients/category"
	courseClient "backend/clients/course"
	userClient "backend/clients/users"

	categoryModel "backend/model/category"
	courseModel "backend/model/courses"

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
	// DB Connections Paramters

	DBName := "test_db"         //Nombre de la base de datos local de ustedes
	DBUser := "root"            //usuario de la base de datos, habitualmente root
	DBPass := "amoamifamilia99" //password del root en la instalacion
	DBHost := "127.0.0.1"       //host de la base de datos. hbitualmente 127.0.0.1

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build

	courseClient.Db = db
	userClient.Db = db
	categoryClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&categoryModel.Category{})
	db.AutoMigrate(&courseModel.Course{})
	db.AutoMigrate(&courseModel.CourseCategory{})
	db.AutoMigrate(&userModel.UserCourses{})

	log.Info("Finishing Migration Database Tables")
}
