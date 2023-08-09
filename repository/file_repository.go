package repository

import (
	"backend/model"

	"gorm.io/gorm"
)

type IFileRepository interface {
	GetAllFiles(files *[]model.File) error
	GetFileReviews(file *model.File, fileId uint) error
	GetFileByUsername(files *[]model.File, username string) error
	CreateFile(file *model.File) error
}

type fileRepository struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) IFileRepository {
	return &fileRepository{db}
}

func (fr *fileRepository) GetAllFiles(files *[]model.File) error {
	if err := fr.db.Find(files).Error; err != nil {
		return err
	}
	return nil
}

func (fr *fileRepository) GetFileReviews(file *model.File, fileId uint) error {
	if err := fr.db.Where("file_id=?", fileId).First(file).Error; err != nil {
		return err
	}
	return nil
}

func (fr *fileRepository) GetFileByUsername(files *[]model.File, username string) error {
	if err := fr.db.Where("username = ?", username).Find(files).Error; err != nil {
		return err
	}
	return nil
}

func (fr *fileRepository) CreateFile(file *model.File) error {
	if err := fr.db.Create(file).Error; err != nil {
		return err
	}
	return nil
}
