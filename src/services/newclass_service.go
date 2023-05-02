package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type NewClassService interface {
	InsertNewClass(nc *entities.NewClassesReq) entities.NewClassesReq
	UpdateNewClass(nc *entities.NewClassesReq) error
	DeleteNewClass(id int) error
	FindAllNewClass(page int, pagesize int) []entities.NewclasssesSet
	FindByID(id int) entities.NewclassesDetail
	FilterNewClass(subID int, classID int, cateID int) []entities.NewclasssesSet
	UpdateStatusNewClass(status int, id int) error
}
type newClassService struct {
	NewClassRepository repositories.NewClassRepository
}

// UpdateStatusNewClass implements NewClassService
func (svc *newClassService) UpdateStatusNewClass(status int, id int) error {
	return svc.NewClassRepository.UpdateStatusNewClass(status, id)
}

// FilterNewClass implements NewClassService
func (svc *newClassService) FilterNewClass(subID int, classID int, cateID int) []entities.NewclasssesSet {
	return svc.NewClassRepository.FilterNewClass(subID, classID, cateID)
}

// DeleteNewClass implements NewClassService
func (svc *newClassService) DeleteNewClass(id int) error {
	return svc.NewClassRepository.DeleteNewClass(id)
}

// FindAllNewClass implements NewClassService
func (svc *newClassService) FindAllNewClass(page int, pagesize int) []entities.NewclasssesSet {
	return svc.NewClassRepository.FindAllNewClass(page, pagesize)
}

// FindByID implements NewClassService
func (svc *newClassService) FindByID(id int) entities.NewclassesDetail {
	return svc.NewClassRepository.FindByID(id)
}

// InsertNewClass implements NewClassService
func (svc *newClassService) InsertNewClass(nc *entities.NewClassesReq) entities.NewClassesReq {
	return svc.NewClassRepository.InsertNewClass(nc)
}

// UpdateNewClass implements NewClassService
func (svc *newClassService) UpdateNewClass(nc *entities.NewClassesReq) error {
	return svc.NewClassRepository.UpdateNewClass(nc)
}

func NewNewClassService(newRepo repositories.NewClassRepository) NewClassService {
	return &newClassService{
		NewClassRepository: newRepo,
	}
}
