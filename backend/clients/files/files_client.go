package files

import (
	e "backend/utils/errors"
	"os"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"

	fileModel "backend/model/files"
)

var Db *gorm.DB

func GetFileById(id int) (fileModel.File, e.ApiError) {
	var file fileModel.File

	err := Db.Where("id_file = ?", id).First(&file).Error

	if err != nil {
		return file, e.NewNotFoundApiError("File not found")
	}

	log.Debug("File: ", file)

	return file, nil
}

func CreateFile(file *fileModel.File) e.ApiError {
	err := Db.Create(&file).Error
	if err != nil {
		return e.NewInternalServerApiError("Error creating file", err)
	}
	return nil

}

func UpdateFile(file fileModel.File) e.ApiError {
	err := Db.Save(&file).Error

	if err != nil {
		return e.NewInternalServerApiError("Error updating file", err)
	}
	return nil
}

func DeleteFile(id int) e.ApiError {

	err := Db.Where("id_file = ?", id).Delete(&fileModel.File{}).Error
	if err != nil {
		return e.NewInternalServerApiError("Error deleting file", err)
	}
	return nil
}

func GetFilesByCourse(idCourse int) ([]fileModel.File, e.ApiError) {
	var files []fileModel.File

	err := Db.Where("id_course = ?", idCourse).Find(&files).Error

	if err != nil {
		return files, e.NewNotFoundApiError("Files not found")
	}

	log.Debug("Files: ", files)

	return files, nil
}

func GetFilesByUser(idUser int) ([]fileModel.File, e.ApiError) {
	var files []fileModel.File

	err := Db.Where("id_user = ?", idUser).Find(&files).Error

	if err != nil {
		return files, e.NewNotFoundApiError("Files not found")
	}

	log.Debug("Files: ", files)

	return files, nil
}

func GetFilesByCourseAndUser(idCourse int, idUser int) ([]fileModel.File, e.ApiError) {
	var files []fileModel.File

	err := Db.Where("id_course = ? AND id_user = ?", idCourse, idUser).Find(&files).Error

	if err != nil {
		return files, e.NewNotFoundApiError("Files not found")
	}

	log.Debug("Files: ", files)

	return files, nil
}

func GetFiles() ([]fileModel.File, e.ApiError) {
	var files []fileModel.File
	err := Db.Find(&files).Error
	if err != nil {
		return files, e.NewNotFoundApiError("Files not found")
	}
	return files, nil
}

func SaveFile(file []byte, path string) e.ApiError {
	err := os.WriteFile(path, file, 0644) // se guarda el archivo en la ruta especificada con los permisos 0644 (lectura y escritura)
	// 0644 es un permiso de lectura y escritura para el propietario y solo lectura para los demás usuarios del sistema operativo
	if err != nil {
		return e.NewInternalServerApiError("Error saving file", err)
	}
	return nil
}
