package files

import (
	fileClient "backend/clients/files"
	"backend/dto"
	fileModel "backend/model/files"
	"log"
	"strings"

	userService "backend/services/users"

	e "backend/utils/errors"
)

type fileService struct{}

type fileServiceInterface interface {
	GetFileById(idFile int) (dto.FileMinDto, e.ApiError)
	SaveFile(file []byte, path string, idCourse int, token string) e.ApiError
}

var (
	FileService fileServiceInterface
)

func init() {
	FileService = &fileService{}
}

func (s *fileService) GetFileById(idFile int) (dto.FileMinDto, e.ApiError) {
	file, err := fileClient.GetFileById(idFile)

	if err != nil {
		return dto.FileMinDto{}, err
	}

	fileMin := dto.FileMinDto{
		IdFile: file.IdFile,
		Name:   file.Name,
		Path:   file.Path,
	}

	return fileMin, nil
}

func (s *fileService) SaveFile(file []byte, path string, idCourse int, token string) e.ApiError {

	user, err := userService.UsersService.GetUsersByToken(token)
	if err != nil {
		return err
	}
	idUser := user.IdUser

	log.Println("User: ", user)

	log.Println("Name: ", strings.Split(path, "uploads\\files\\"))

	fileToCreate := fileModel.File{
		Name:     strings.Split(path, "uploads\\files\\")[1],
		Path:     path,
		IdCourse: idCourse,
		IdUser:   idUser,
	}

	err = fileClient.CreateFile(&fileToCreate)
	if err != nil {
		return err
	}

	log.Println("File created: ", fileToCreate)

	err = fileClient.SaveFile(file, path)
	if err != nil {
		fileClient.DeleteFile(fileToCreate.IdFile)
		return err
	}

	return nil
}
