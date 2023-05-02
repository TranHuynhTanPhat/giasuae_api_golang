package repositories

import (
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type SubjectRepository interface {
	InsertSubject(s *entities.Subject) error
	UpdateSubject(s *entities.Subject) error
	DeleteSubject(id int) error
	FindAllSubject() []entities.Subject
	FindByID(id int) entities.Subject
}

type subjectConnection struct {
	connection *gorm.DB
}

// DeleteSubject implements SubjectRepository
func (db *subjectConnection) DeleteSubject(id int) error {
	if err := db.connection.Table("subjects_of_newclasses").Where("id_subject = ?", id).Delete(&entities.SubjectsOfNewclasses{}); err.Error != nil {
		return err.Error
	}

	if err := db.connection.Table("subjects_of_tutors").Where("id_subject = ?", id).Delete(&entities.SubjectsOfTutor{}); err.Error != nil {
		return err.Error
	}
	if err := db.connection.Table("subjects").Delete(&entities.Subject{}, id); err.Error != nil {
		return err.Error
	}
	return nil
}

// FindAllSubject implements SubjectRepository
func (db *subjectConnection) FindAllSubject() []entities.Subject {
	var subjects []entities.Subject
	db.connection.Find(&subjects)
	return subjects
}

// FindByID implements SubjectRepository
func (db *subjectConnection) FindByID(id int) entities.Subject {
	var sub entities.Subject
	db.connection.First(&sub, id)
	return sub
}

// InsertSubject implements SubjectRepository
func (db *subjectConnection) InsertSubject(s *entities.Subject) error {
	err := db.connection.Save(&s)
	if err != nil {
		return err.Error
	}
	return nil
}

// UpdateSubject implements SubjectRepository
func (db *subjectConnection) UpdateSubject(s *entities.Subject) error {
	err := db.connection.Save(&s)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func NewSubjectRepository(dbConn *gorm.DB) SubjectRepository {
	return &subjectConnection{
		connection: dbConn,
	}
}
