package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type TutorService interface {
	InsertTutor(tutor *entities.TutorReq) error
	UpdateTutor(tutor *entities.TutorReq) error
	DeleteTutor(id int) error
	FindAllTutor() []entities.TutorSet
	FindByID(id int) entities.TutorDetail
	FilterTutor(subID int, classID int, cateID int, gender string, isnow string) []entities.TutorSet
}
type tutorService struct {
	TutorRepository repositories.TutorRepository
}

// FilterTutor implements TutorService
func (svc *tutorService) FilterTutor(subID int, classID int, cateID int, gender string, isnow string) []entities.TutorSet {
	return svc.TutorRepository.FilterTutor(subID, classID, cateID, gender, isnow)
}

// DeleteTutor implements TutorService
func (svc *tutorService) DeleteTutor(id int) error {
	return svc.TutorRepository.DeleteTutor(id)
}

// UpdateTutor implements TutorService
func (svc *tutorService) UpdateTutor(tutor *entities.TutorReq) error {
	return svc.TutorRepository.UpdateTutor(tutor)
}

// FindAllTutor implements TutorService
func (svc *tutorService) FindAllTutor() []entities.TutorSet {
	return svc.TutorRepository.FindAllTutor()
}

// FindByID implements TutorService
func (svc *tutorService) FindByID(id int) entities.TutorDetail {
	return svc.TutorRepository.FindByID(id)
}

// InsertTutor implements TutorService
func (svc *tutorService) InsertTutor(tutor *entities.TutorReq) error {
	return svc.TutorRepository.InsertTutor(tutor)
}

func NewTutorService(repo repositories.TutorRepository) TutorService {
	return &tutorService{
		TutorRepository: repo,
	}
}
