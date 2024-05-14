package db

import (
	courseClient "backend/clients/course"
	professorClient "backend/clients/professor"
	studentClient "backend/clients/student"
	userClient "backend/clients/users"

	courseModel "backend/model/courses"
	professorModel "backend/model/professor"
	studentModel "backend/model/student"
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
	DBName := "softArch"  //Nombre de la base de datos local de ustedes
	DBUser := "root"      //usuario de la base de datos, habitualmente root
	DBPass := "Nano01234" //password del root en la instalacion
	DBHost := "127.0.0.1" //host de la base de datos. hbitualmente 127.0.0.1
	// ------------------------

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	studentClient.Db = db
	professorClient.Db = db
	courseClient.Db = db
	userClient.Db = db

}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&userModel.User{})
	db.AutoMigrate(&studentModel.Student{})
	db.AutoMigrate(&professorModel.Professor{})
	db.AutoMigrate(&courseModel.Course{})

	log.Info("Finishing Migration Database Tables")
}
