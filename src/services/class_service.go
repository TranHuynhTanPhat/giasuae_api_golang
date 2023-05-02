package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type ClassService interface {
	InsertClass(c *entities.Class) error
	UpdateClass(c *entities.Class) error
	DeleteClass(id int) error
	FindAllClass() []entities.Class
	FindByID(id int) entities.Class
}
type classService struct {
	ClassRepository repositories.ClassRepository
}

// DeleteClass implements ClassService
func (svc *classService) DeleteClass(id int) error {
	return svc.ClassRepository.DeleteClass(id)
}

// FindAllClass implements ClassService
func (svc *classService) FindAllClass() []entities.Class {
	return svc.ClassRepository.FindAllClass()
}

// FindByID implements ClassService
func (svc *classService) FindByID(id int) entities.Class {
	return svc.ClassRepository.FindByID(id)
}

// InsertClass implements ClassService
func (svc *classService) InsertClass(c *entities.Class) error {
	return svc.ClassRepository.InsertClass(c)
}

// UpdateClass implements ClassService
func (svc *classService) UpdateClass(c *entities.Class) error {
	return svc.ClassRepository.UpdateClass(c)
}

func NewClassITService(classRepo repositories.ClassRepository) ClassService {
	return &classService{
		ClassRepository: classRepo,
	}
}
