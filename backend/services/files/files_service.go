package files

import (
	fileClient "backend/clients/files"
	"backend/dto"
	fileModel "backend/model/files"
	"path/filepath"

	userService "backend/services/users"

	e "backend/utils/errors"
)

type fileService struct{}

type fileServiceInterface interface {
	GetFileById(idFile int) (dto.FileMinDto, e.ApiError)
	SaveFile(file []byte, path string, idCourse int, token string) e.ApiError
	GetFilesByCourse(idCourse int) ([]dto.FileMinDto, e.ApiError)
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
	// Retrieve user information from token
	user, err := userService.UsersService.GetUsersByToken(token)
	if err != nil {
		return err
	}
	idUser := user.IdUser

	// Extract filename from path
	fileName := filepath.Base(path)

	// Create file model
	fileToCreate := fileModel.File{
		Name:     fileName,
		Path:     path,
		IdCourse: idCourse,
		IdUser:   idUser,
	}

	// Attempt to create file entry in database
	err = fileClient.CreateFile(&fileToCreate)
	if err != nil {
		return err
	}

	// Attempt to save file content
	err = fileClient.SaveFile(file, path)
	if err != nil {
		// Rollback: Delete created file entry if saving content fails
		fileClient.DeleteFile(fileToCreate.IdFile)
		return err
	}

	return nil
}

func (s *fileService) GetFilesByCourse(idCourse int) ([]dto.FileMinDto, e.ApiError) {
	files, err := fileClient.GetFilesByCourse(idCourse)
	if err != nil {
		return nil, err
	}

	var filesMin []dto.FileMinDto
	for _, file := range files {
		fileMin := dto.FileMinDto{
			IdFile: file.IdFile,
			Name:   file.Name,
			Path:   file.Path,
		}
		filesMin = append(filesMin, fileMin)
	}

	return filesMin, nil
}
