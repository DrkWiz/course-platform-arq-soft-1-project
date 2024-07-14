package db

import (
	categoryClient "backend/clients/category"
	courseClient "backend/clients/course"
	fileClient "backend/clients/files"
	userClient "backend/clients/users"

	e "backend/utils/errors"
	"fmt"
	"os"

	categoryModel "backend/model/category"
	courseModel "backend/model/courses"
	fileModel "backend/model/files"
	userModel "backend/model/users"

	userServices "backend/services/users"

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

	err := SeedDatabaseUsers()
	if err != nil {
		log.Error(err)
	}

	err1 := SeedDatabaseCourses()
	if err1 != nil {
		log.Error(err1)
	}

	err2 := SeedDatabaseUsersCourses()
	if err2 != nil {
		log.Error(err2)
	}

	err3 := SeedDatabaseCategories()
	if err3 != nil {
		log.Error(err3)
	}

	err4 := SeedDatabaseCoursesCategories()
	if err4 != nil {
		log.Error(err4)
	}

	log.Info("Finishing Migration Database Tables")
}

func SeedDatabaseUsers() e.ApiError {

	hashPasswordAdmin, err := userServices.UsersService.HashPassword("admin")
	if err != nil {
		log.Error("Error Hashing Password")
		return err
	}

	admin := userModel.User{
		Name:     "Super Duper Secret EasterEgg 1",
		Username: "admin",
		Password: hashPasswordAdmin,
		Email:    "admin@admin.admin",
		IsAdmin:  true,
	}
	db.Create(&admin)

	hashPasswordUser, err := userServices.UsersService.HashPassword("user")
	if err != nil {
		log.Error("Error Hashing Password")
		return e.NewInternalServerApiError("Error hashing password", err)
	}

	user := userModel.User{
		Name:     "Super Duper Secret EasterEgg 2",
		Username: "user",
		Password: hashPasswordUser,
		Email:    "user@user.user",
		IsAdmin:  false,
	}
	db.Create(&user)

	log.Info("Finishing Seeding users Database")

	return nil
}

func SeedDatabaseCourses() e.ApiError {

	course1 := courseModel.Course{
		Name:        "Semiconductores",
		Description: "curso de teoria de semiconductores",
		Price:       200,
		PicturePath: "Semiconductordevices.jpg",
		StartDate:   "2024-06-02",
		EndDate:     "2026-06-02",
		IdOwner:     1,
		IsActive:    true,
	}
	db.Create(&course1)

	course2 := courseModel.Course{
		Name:        "Electromagnetismo",
		Description: "Teoria y practicas de electromagentismo explicado para cualquiera",
		Price:       120,
		PicturePath: "electromagnetismo.jpg",
		StartDate:   "2024-06-02",
		EndDate:     "2026-06-02",
		IdOwner:     1,
		IsActive:    true,
	}
	db.Create(&course2)

	course3 := courseModel.Course{
		Name:        "programacion en C++ ",
		Description: "programacion en C++ explicado para principiantes",
		Price:       330,
		PicturePath: "cpp.png",
		StartDate:   "2024-06-02",
		EndDate:     "2026-06-02",
		IdOwner:     1,
		IsActive:    true,
	}
	db.Create(&course3)

	course4 := courseModel.Course{
		Name:        "Circuitismo",
		Description: "teoria y metodo de diseño de circuitos para Formula 1",
		Price:       10,
		PicturePath: "circuitos.jpg",
		StartDate:   "2024-06-02",
		EndDate:     "2026-06-02",
		IdOwner:     1,
		IsActive:    true,
	}
	db.Create(&course4)

	course5 := courseModel.Course{
		Name:        "Arquitectura de Software",
		Description: "La materia mas piola del semestre sin lugar a dudas,profe por favor apruebenos con un 10 ",
		Price:       421,
		PicturePath: "softarch.jpg",
		StartDate:   "2024-06-02",
		EndDate:     "2026-06-02",
		IdOwner:     1,
		IsActive:    true,
	}
	db.Create(&course5)

	course6 := courseModel.Course{
		Name:        "Bromeismo (Bromo)",
		Description: "Curso para hacer Un elemento químico de número atómico 35 situado en el grupo de los halógenos (grupo XVII) de la tabla periódica de los elementos. Su símbolo es Br. Bromo a temperatura ambiente es un líquido rojo, volátil y denso",
		Price:       4123,
		PicturePath: "bromas.jpg",
		StartDate:   "2024-06-02",
		EndDate:     "2026-06-02",
		IdOwner:     1,
		IsActive:    true,
	}
	db.Create(&course6)

	log.Info("Finishing Seeding course Database")

	return nil
}

func SeedDatabaseUsersCourses() e.ApiError {

	adminCourses1 := userModel.UserCourses{
		IdUser:   1,
		IdCourse: 1,
		Comment:  "nice course",
		Rating:   4,
	}
	db.Create(&adminCourses1)

	adminCourses2 := userModel.UserCourses{
		IdUser:   1,
		IdCourse: 2,
		Comment:  "nice course",
		Rating:   2,
	}
	db.Create(&adminCourses2)

	adminCourses3 := userModel.UserCourses{
		IdUser:   1,
		IdCourse: 3,
		Comment:  "nice course",
		Rating:   3,
	}
	db.Create(&adminCourses3)

	adminCourses4 := userModel.UserCourses{
		IdUser:   1,
		IdCourse: 4,
		Comment:  "a course",
		Rating:   1,
	}
	db.Create(&adminCourses4)

	adminCourses5 := userModel.UserCourses{
		IdUser:   1,
		IdCourse: 5,
		Comment:  "nice course",
		Rating:   5,
	}
	db.Create(&adminCourses5)

	adminCourses6 := userModel.UserCourses{
		IdUser:   1,
		IdCourse: 6,
		Comment:  "nice course",
		Rating:   3,
	}
	db.Create(&adminCourses6)

	userCourses1 := userModel.UserCourses{
		IdUser:   2,
		IdCourse: 1,
		Comment:  "gran profesor",
		Rating:   5,
	}
	db.Create(&userCourses1)

	userCourses6 := userModel.UserCourses{
		IdUser:   2,
		IdCourse: 6,
		Comment:  "tremendo chiste, ¿NO?",
		Rating:   4,
	}
	db.Create(&userCourses6)

	userCourses3 := userModel.UserCourses{
		IdUser:   2,
		IdCourse: 3,
		Comment:  "pozole",
		Rating:   5,
	}
	db.Create(&userCourses3)

	userCourses5 := userModel.UserCourses{
		IdUser:   2,
		IdCourse: 5,
		Comment:  "alto curso",
		Rating:   5,
	}
	db.Create(&userCourses5)

	log.Info("Finishing Seeding user courses Database")

	return nil
}

func SeedDatabaseCategories() e.ApiError {

	category1 := categoryModel.Category{
		Name: "tecnologia",
	}
	db.Create(&category1)

	category2 := categoryModel.Category{
		Name: "programacion",
	}
	db.Create(&category2)

	category3 := categoryModel.Category{
		Name: "principiantes",
	}
	db.Create(&category3)

	category4 := categoryModel.Category{
		Name: "futurista",
	}
	db.Create(&category4)

	category5 := categoryModel.Category{
		Name: "avanzado",
	}
	db.Create(&category5)

	category6 := categoryModel.Category{
		Name: "bromo :)",
	}
	db.Create(&category6)

	log.Info("Finishing Seeding category Database")

	return nil
}

func SeedDatabaseCoursesCategories() e.ApiError {

	course1Category1 := courseModel.CourseCategory{
		IdCourse:   1,
		IdCategory: 1,
	}
	db.Create(&course1Category1)

	course2Category1 := courseModel.CourseCategory{
		IdCourse:   2,
		IdCategory: 1,
	}
	db.Create(&course2Category1)

	course3Category1 := courseModel.CourseCategory{
		IdCourse:   3,
		IdCategory: 1,
	}
	db.Create(&course3Category1)

	course4Category1 := courseModel.CourseCategory{
		IdCourse:   4,
		IdCategory: 1,
	}
	db.Create(&course4Category1)

	course5Category1 := courseModel.CourseCategory{
		IdCourse:   5,
		IdCategory: 1,
	}
	db.Create(&course5Category1)

	course6Category1 := courseModel.CourseCategory{
		IdCourse:   6,
		IdCategory: 1,
	}
	db.Create(&course6Category1)

	course6Category6 := courseModel.CourseCategory{
		IdCourse:   6,
		IdCategory: 6,
	}
	db.Create(&course6Category6)

	course2Category4 := courseModel.CourseCategory{
		IdCourse:   2,
		IdCategory: 4,
	}
	db.Create(&course2Category4)

	course2Category2 := courseModel.CourseCategory{
		IdCourse:   2,
		IdCategory: 2,
	}
	db.Create(&course2Category2)

	course3Category2 := courseModel.CourseCategory{
		IdCourse:   3,
		IdCategory: 2,
	}
	db.Create(&course3Category2)

	course3Category5 := courseModel.CourseCategory{
		IdCourse:   3,
		IdCategory: 5,
	}
	db.Create(&course3Category5)

	course4Category3 := courseModel.CourseCategory{
		IdCourse:   4,
		IdCategory: 3,
	}
	db.Create(&course4Category3)

	log.Info("Finishing Seeding courses categories Database")

	return nil
}
