package model

type File struct {
	IdFile   int    `gorm:"PRIMARY_KEY;AUTO_INCREMENT;column:id_file"`
	Name     string `gorm:"type:VARCHAR(350);not null"`
	Path     string `gorm:"type:VARCHAR(350);not null"`
	IdCourse int    `gorm:"type:int;not null"`
	IdUser   int    `gorm:"type:int;not null"`
}

type Files []File
