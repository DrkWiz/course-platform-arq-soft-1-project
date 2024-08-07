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

/*
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
*/

func SeedDatabaseUsers() e.ApiError {

	hashPasswordAdmin, err := userServices.UsersService.HashPassword("admin")
	if err != nil {
		log.Error("Error Hashing Password")
		return err
	}

	admin := userModel.User{
		Username: "admin",
	}

	// Verificar si el usuario admin ya existe
	var existingAdmin userModel.User
	if err := db.Where("username = ?", admin.Username).First(&existingAdmin).Error; err != nil {
		// Si no existe, crearlo
		admin.Name = "Super Duper Secret EasterEgg 1"
		admin.Password = hashPasswordAdmin
		admin.Email = "admin@admin.admin"
		admin.IsAdmin = true
		if err := db.Create(&admin).Error; err != nil {
			log.Error("Error creating admin user")
			return e.NewInternalServerApiError("Error creating admin user", err)
		}
	} else {
		log.Info("Admin user already exists, skipping creation")
	}

	hashPasswordUser, err := userServices.UsersService.HashPassword("user")
	if err != nil {
		log.Error("Error Hashing Password")
		return e.NewInternalServerApiError("Error hashing password", err)
	}

	user := userModel.User{
		Username: "user",
	}

	// Verificar si el usuario user ya existe
	var existingUser userModel.User
	if err := db.Where("username = ?", user.Username).First(&existingUser).Error; err != nil {
		// Si no existe, crearlo
		user.Name = "Super Duper Secret EasterEgg 2"
		user.Password = hashPasswordUser
		user.Email = "user@user.user"
		user.IsAdmin = false
		if err := db.Create(&user).Error; err != nil {
			log.Error("Error creating user")
			return e.NewInternalServerApiError("Error creating user", err)
		}
	} else {
		log.Info("User already exists, skipping creation")
	}

	log.Info("Finishing Seeding users Database")

	return nil
}

/*
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
*/

func SeedDatabaseCourses() e.ApiError {

	courses := []courseModel.Course{
		{
			Name:        "Semiconductores",
			Description: "curso de teoria de semiconductores",
			Price:       200,
			PicturePath: "Semiconductordevices.jpg",
			StartDate:   "2024-06-02",
			EndDate:     "2026-06-02",
			IdOwner:     1,
			IsActive:    true,
		},
		{
			Name:        "Electromagnetismo",
			Description: "Teoria y practicas de electromagentismo explicado para cualquiera.",
			Price:       120,
			PicturePath: "electromagnetismo.jpg",
			StartDate:   "2024-06-02",
			EndDate:     "2026-06-02",
			IdOwner:     1,
			IsActive:    true,
		},
		{
			Name:        "programacion en C++ ",
			Description: "programacion en C++ explicado para principiantes con ejemplos practicos.",
			Price:       330,
			PicturePath: "cpp.png",
			StartDate:   "2024-06-02",
			EndDate:     "2026-06-02",
			IdOwner:     1,
			IsActive:    true,
		},
		{
			Name:        "Circuitismo",
			Description: "teoria y metodo de diseño de circuitos para Formula 1, 2, 3 y 4.",
			Price:       10,
			PicturePath: "circuitos.jpg",
			StartDate:   "2024-06-02",
			EndDate:     "2026-06-02",
			IdOwner:     1,
			IsActive:    true,
		},
		{
			Name:        "Arquitectura de Software",
			Description: "La materia mas piola del semestre sin lugar a dudas, profe por favor apruebenos con un 10 ",
			Price:       421,
			PicturePath: "softarch.jpg",
			StartDate:   "2024-06-02",
			EndDate:     "2026-06-02",
			IdOwner:     1,
			IsActive:    true,
		},
		{
			Name:        "Bromeismo (Bromo)",
			Description: "Curso para hacer Un elemento químico de número atómico 35 situado en el grupo de los halógenos (grupo XVII) de la tabla periódica de los elementos. Su símbolo es Br. Bromo a temperatura ambiente es un líquido rojo, volátil y denso",
			Price:       4123,
			PicturePath: "bromas.jpg",
			StartDate:   "2024-06-02",
			EndDate:     "2026-06-02",
			IdOwner:     1,
			IsActive:    true,
		},
	}

	for _, course := range courses {
		// Verificar si el curso ya existe
		var existingCourse courseModel.Course
		if err := db.Where("name = ?", course.Name).First(&existingCourse).Error; err != nil {
			// Si no existe, crearlo
			if err := db.Create(&course).Error; err != nil {
				log.Error("Error creating course")
				return e.NewInternalServerApiError("Error creating course", err)
			}
		} else {
			log.Info("Course already exists, skipping creation")
		}
	}

	log.Info("Finishing Seeding courses Database")

	return nil
}

/*
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
*/

func SeedDatabaseUsersCourses() e.ApiError {

	userCourses := []userModel.UserCourses{
		{
			IdUser:   1,
			IdCourse: 1,
			Comment:  "nice course",
			Rating:   4,
		},
		{
			IdUser:   1,
			IdCourse: 2,
			Comment:  "nice course",
			Rating:   2,
		},
		{
			IdUser:   1,
			IdCourse: 3,
			Comment:  "nice course",
			Rating:   3,
		},
		{
			IdUser:   1,
			IdCourse: 4,
			Comment:  "a course",
			Rating:   1,
		},
		{
			IdUser:   1,
			IdCourse: 5,
			Comment:  "nice course",
			Rating:   5,
		},
		{
			IdUser:   1,
			IdCourse: 6,
			Comment:  "nice course",
			Rating:   3,
		},
		{
			IdUser:   2,
			IdCourse: 1,
			Comment:  "gran profesor",
			Rating:   5,
		},
		{
			IdUser:   2,
			IdCourse: 6,
			Comment:  "tremendo chiste, ¿NO?",
			Rating:   4,
		},
		{
			IdUser:   2,
			IdCourse: 3,
			Comment:  "pozole",
			Rating:   5,
		},
		{
			IdUser:   2,
			IdCourse: 5,
			Comment:  "alto curso",
			Rating:   5,
		},
	}

	for _, userCourse := range userCourses {
		// Verificar si el registro de usuario y curso ya existe
		var existingUserCourse userModel.UserCourses
		if err := db.Where("id_user = ? AND id_course = ?", userCourse.IdUser, userCourse.IdCourse).First(&existingUserCourse).Error; err != nil {
			// Si no existe, crearlo
			if err := db.Create(&userCourse).Error; err != nil {
				log.Error("Error creating user course")
				return e.NewInternalServerApiError("Error creating user course", err)
			}
		} else {
			log.Info("User course already exists, skipping creation")
		}
	}

	log.Info("Finishing Seeding user courses Database")

	return nil
}

/*
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
*/

func SeedDatabaseCategories() e.ApiError {

	categories := []categoryModel.Category{
		{
			Name: "tecnologia",
		},
		{
			Name: "programacion",
		},
		{
			Name: "principiantes",
		},
		{
			Name: "futurista",
		},
		{
			Name: "avanzado",
		},
		{
			Name: "bromo :)",
		},
	}

	for _, category := range categories {
		// Verificar si la categoría ya existe
		var existingCategory categoryModel.Category
		if err := db.Where("name = ?", category.Name).First(&existingCategory).Error; err != nil {
			// Si no existe, crearla
			if err := db.Create(&category).Error; err != nil {
				log.Error("Error creating category")
				return e.NewInternalServerApiError("Error creating category", err)
			}
		} else {
			log.Info("Category already exists, skipping creation")
		}
	}

	log.Info("Finishing Seeding category Database")

	return nil
}

/*
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
*/

func SeedDatabaseCoursesCategories() e.ApiError {

	coursesCategories := []courseModel.CourseCategory{
		{
			IdCourse:   1,
			IdCategory: 1,
		},
		{
			IdCourse:   2,
			IdCategory: 1,
		},
		{
			IdCourse:   3,
			IdCategory: 1,
		},
		{
			IdCourse:   4,
			IdCategory: 1,
		},
		{
			IdCourse:   5,
			IdCategory: 1,
		},
		{
			IdCourse:   6,
			IdCategory: 1,
		},
		{
			IdCourse:   6,
			IdCategory: 6,
		},
		{
			IdCourse:   2,
			IdCategory: 4,
		},
		{
			IdCourse:   2,
			IdCategory: 2,
		},
		{
			IdCourse:   3,
			IdCategory: 2,
		},
		{
			IdCourse:   3,
			IdCategory: 5,
		},
		{
			IdCourse:   4,
			IdCategory: 3,
		},
	}

	for _, courseCategory := range coursesCategories {
		// Verificar si el registro de curso y categoría ya existe
		var existingCourseCategory courseModel.CourseCategory
		if err := db.Where("id_course = ? AND id_category = ?", courseCategory.IdCourse, courseCategory.IdCategory).First(&existingCourseCategory).Error; err != nil {
			// Si no existe, crearlo
			if err := db.Create(&courseCategory).Error; err != nil {
				log.Error("Error creating course category")
				return e.NewInternalServerApiError("Error creating course category", err)
			}
		} else {
			log.Info("Course category already exists, skipping creation")
		}
	}

	log.Info("Finishing Seeding courses categories Database")

	return nil
}
