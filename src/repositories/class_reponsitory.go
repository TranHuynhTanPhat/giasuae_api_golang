package repositories

import (
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type ClassRepository interface {
	InsertClass(c *entities.Class) error
	UpdateClass(c *entities.Class) error
	DeleteClass(id int) error
	FindAllClass() []entities.Class
	FindByID(id int) entities.Class
}

type classConnection struct {
	connection *gorm.DB
}

// DeleteClass implements ClassRepository
func (db *classConnection) DeleteClass(id int) error {
	if err := db.connection.Table("classes_of_newclasses").Where("id_class = ?", id).Delete(&entities.ClassesOfNewclasses{}); err.Error != nil {
		return err.Error
	}

	if err := db.connection.Table("classes_of_tutors").Where("id_class = ?", id).Delete(&entities.ClassesOfTutor{}); err.Error != nil {
		return err.Error
	}
	if err := db.connection.Table("classes").Delete(&entities.Class{}, id); err.Error != nil {
		return err.Error
	}
	return nil
}

// FindAllClass implements ClassRepository
func (db *classConnection) FindAllClass() []entities.Class {
	var classes []entities.Class
	db.connection.Find(&classes)
	return classes
}

// FindByID implements ClassRepository
func (db *classConnection) FindByID(id int) entities.Class {
	var class entities.Class
	db.connection.First(&class, id)
	return class
}

// InsertClass implements ClassRepository
func (db *classConnection) InsertClass(c *entities.Class) error {
	err := db.connection.Save(&c)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// UpdateClass implements ClassRepository
func (db *classConnection) UpdateClass(c *entities.Class) error {
	err := db.connection.Save(&c)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func NewClassITRepository(dbConn *gorm.DB) ClassRepository {
	return &classConnection{
		connection: dbConn,
	}
}
