package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type SubjectService interface {
	InsertSubject(s *entities.Subject) error
	UpdateSubject(s *entities.Subject) error
	DeleteSubject(id int) error
	FindAllSubject() []entities.Subject
	FindByID(id int) entities.Subject
}
type subjectService struct {
	SubjectRepository repositories.SubjectRepository
}

// DeleteSubject implements SubjectService
func (svc *subjectService) DeleteSubject(id int) error {
	return svc.SubjectRepository.DeleteSubject(id)
}

// FindAllSubject implements SubjectService
func (services *subjectService) FindAllSubject() []entities.Subject {
	return services.SubjectRepository.FindAllSubject()
}

// FindByID implements SubjectService
func (subRepo *subjectService) FindByID(id int) entities.Subject {
	return subRepo.SubjectRepository.FindByID(int(id))
}

// InsertSubject implements SubjectService
func (subRepo *subjectService) InsertSubject(s *entities.Subject) error {
	return subRepo.SubjectRepository.InsertSubject(s)
}

// UpdateSubject implements SubjectService
func (svc *subjectService) UpdateSubject(s *entities.Subject) error {
	return svc.SubjectRepository.UpdateSubject(s)
}

func NewSubjectService(subRepo repositories.SubjectRepository) SubjectService {
	return &subjectService{
		SubjectRepository: subRepo,
	}
}
