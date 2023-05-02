package repositories

import (
	"giasuaeapi/src/entities"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type NewClassRepository interface {
	InsertNewClass(nc *entities.NewClassesReq) entities.NewClassesReq
	UpdateNewClass(nc *entities.NewClassesReq) error
	DeleteNewClass(id int) error
	FindAllNewClass(page int, pagesize int) []entities.NewclasssesSet
	FindByID(id int) entities.NewclassesDetail
	FilterNewClass(subID int, classID int, cateID int) []entities.NewclasssesSet
	UpdateStatusNewClass(status int, id int) error
}

type newClassConnection struct {
	connection *gorm.DB
}

// UpdateStatusNewClass implements NewClassRepository
func (db *newClassConnection) UpdateStatusNewClass(status int, id int) error {
	err := db.connection.Table("newclasses").Where("id = ?", id).Update("status", status)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// FilterNewClass implements NewClassRepository
func (db *newClassConnection) FilterNewClass(subID int, classID int, cateID int) []entities.NewclasssesSet {
	var newclasses []entities.NewclasssesSet
	db.connection.Table("newclasses").
		Joins("INNER JOIN subjects_of_newclasses ON newclasses.id IN (SELECT id_newclass FROM subjects_of_newclasses WHERE subjects_of_newclasses.id_subject = ?) OR ? = 0 ", subID, subID).
		Joins("INNER JOIN classes_of_newclasses ON newclasses.id IN (SELECT id_newclass FROM classes_of_newclasses WHERE classes_of_newclasses.id_class = ?) OR ? = 0 ", classID, classID).
		Joins("INNER JOIN categories_of_newclasses ON newclasses.id IN (SELECT id_newclass FROM categories_of_newclasses WHERE categories_of_newclasses.id_category = ?) OR ? = 0 ", cateID, cateID).
		Select(queyGetAllNewClass()).Group("newclasses.id").Find(&newclasses)
	return newclasses
}

// DeleteNewClass implements NewClassRepository
func (db *newClassConnection) DeleteNewClass(id int) error {
	if err := db.connection.Table("classes_of_newclasses").Where("id_newclass = ?", id).Delete(&entities.ClassesOfNewclasses{}); err.Error != nil {
		return err.Error
	}
	if err := db.connection.Table("subjects_of_newclasses").Where("id_newclass = ?", id).Delete(&entities.SubjectsOfNewclasses{}); err.Error != nil {
		return err.Error
	}
	if err := db.connection.Table("categories_of_newclasses").Where("id_newclass = ?", id).Delete(&entities.CategoriesOfNewclasses{}); err.Error != nil {
		return err.Error
	}

	if err := db.connection.Table("newclasses").Delete(&entities.NewClassesDefault{}, id); err.Error != nil {
		return err.Error
	}
	return nil
}

// FindAllNewClass implements NewClassRepository
func (db *newClassConnection) FindAllNewClass(page int, pagesize int) []entities.NewclasssesSet {
	var newclasses []entities.NewclasssesSet
	db.connection.Limit(pagesize).Offset((page - 1) * pagesize).Table("newclasses").Select(queyGetAllNewClass()).Group("newclasses.id").Find(&newclasses)
	return newclasses
}

// FindByID implements NewClassRepository
func (db *newClassConnection) FindByID(id int) entities.NewclassesDetail {
	var newclasses entities.NewclassesDetail
	db.connection.Limit(1).Table("newclasses").Where("id = ?", id).Scan(&newclasses)
	newclasses.Subjects = getListSubjectOfNC(db, id)
	newclasses.Classes = getListClassOfNC(db, id)
	newclasses.Categories = getListCategoryOfNC(db, id)
	return newclasses
}

// InsertNewClass implements NewClassRepository
func (db *newClassConnection) InsertNewClass(nc *entities.NewClassesReq) entities.NewClassesReq {
	var subOfNC []entities.SubjectsOfNewclasses
	var classOfNC []entities.ClassesOfNewclasses
	var ctgOfNC []entities.CategoriesOfNewclasses
	db.connection.Table("newclasses").Create(&nc)
	for _, value := range nc.Subjects {
		subOfNC = append(subOfNC, entities.SubjectsOfNewclasses{
			ID_newclass: nc.ID,
			ID_subject:  value,
		})
	}
	for _, value := range nc.Classes {
		classOfNC = append(classOfNC, entities.ClassesOfNewclasses{
			ID_newclass: nc.ID,
			ID_class:    value,
		})
	}
	for _, value := range nc.Categories {
		ctgOfNC = append(ctgOfNC, entities.CategoriesOfNewclasses{
			ID_newclass: nc.ID,
			ID_category: value,
		})
	}
	db.connection.Create(&subOfNC)
	db.connection.Create(&classOfNC)
	db.connection.Create(&ctgOfNC)
	return *nc
}

// UpdateNewClass implements NewClassRepository
func (db *newClassConnection) UpdateNewClass(nc *entities.NewClassesReq) error {
	delListSubjectOfNC(db, nc.ID)
	delListClassOfNC(db, nc.ID)
	delListCategoryOfNC(db, nc.ID)
	var newClass entities.NewClassesDefault
	var subOfNC []entities.SubjectsOfNewclasses
	var clasOfNC []entities.ClassesOfNewclasses
	var ctgOfNC []entities.CategoriesOfNewclasses
	smapping.FillStruct(&newClass, smapping.MapFields(&nc))
	err := db.connection.Table("newclasses").Save(&newClass)
	if err.Error != nil {
		return err.Error
	}
	for _, value := range nc.Subjects {
		subOfNC = append(subOfNC, entities.SubjectsOfNewclasses{
			ID_newclass: nc.ID,
			ID_subject:  value,
		})
	}
	for _, value := range nc.Classes {
		clasOfNC = append(clasOfNC, entities.ClassesOfNewclasses{
			ID_newclass: nc.ID,
			ID_class:    value,
		})
	}
	for _, value := range nc.Categories {
		ctgOfNC = append(ctgOfNC, entities.CategoriesOfNewclasses{
			ID_newclass: nc.ID,
			ID_category: value,
		})
	}

	err1 := db.connection.Table("subjects_of_newclasses").Save(&subOfNC)
	err2 := db.connection.Table("classes_of_newclasses").Save(&clasOfNC)
	err3 := db.connection.Table("categories_of_newclasses").Save(&ctgOfNC)
	if err1.Error != nil {
		return err1.Error
	}
	if err2.Error != nil {
		return err1.Error
	}
	if err3.Error != nil {
		return err1.Error
	}
	return nil
}

func NewNewClassRepository(dbConn *gorm.DB) NewClassRepository {
	return &newClassConnection{
		connection: dbConn,
	}
}

func getListSubjectOfNC(db *newClassConnection, ncId int) []entities.Subject {
	var subjects []entities.Subject
	db.connection.Table("subjects").Joins("inner join subjects_of_newclasses on id_newclass = ?", ncId).Where("id = id_subject").Scan(&subjects)
	return subjects
}
func getListClassOfNC(db *newClassConnection, ncId int) []entities.Class {
	var classes []entities.Class
	db.connection.Table("classes").Joins("inner join classes_of_newclasses on id_newclass = ?", ncId).Where("id = id_class").Scan(&classes)
	return classes
}
func getListCategoryOfNC(db *newClassConnection, ncId int) []entities.Category {
	var categories []entities.Category
	db.connection.Table("categories").Joins("inner join categories_of_newclasses on id_newclass = ?", ncId).Where("id = id_category").Scan(&categories)
	return categories
}

func delListSubjectOfNC(db *newClassConnection, ncId int) {
	db.connection.Table("subjects_of_newclasses").Where("id_newclass = ?", ncId).Delete(&entities.SubjectsOfNewclasses{})
}
func delListClassOfNC(db *newClassConnection, ncId int) {
	db.connection.Table("classes_of_newclasses").Where("id_newclass = ?", ncId).Delete(&entities.ClassesOfNewclasses{})
}
func delListCategoryOfNC(db *newClassConnection, ncId int) {
	db.connection.Table("categories_of_newclasses").Where("id_newclass = ?", ncId).Delete(&entities.CategoriesOfNewclasses{})
}

func queyGetAllNewClass() string {
	return `newclasses.id,
	newclasses.address,
	newclasses.district,
	newclasses.sobuoi,
	newclasses.time,
	newclasses.salary,
	newclasses.require,
	newclasses.status,
	newclasses.contact,
	newclasses.created_at,
	(SELECT 
		GROUP_CONCAT((SELECT c.name
			FROM
				classes c
			WHERE
				(c.id = ncc.id_class))
			SEPARATOR ', ')
		FROM
			classes_of_newclasses ncc
		WHERE
			(ncc.id_newclass = newclasses.id)) AS classes,
	(SELECT 
		GROUP_CONCAT((SELECT s.name
			FROM
				subjects s
			WHERE
				(s.id = ncs.id_subject))
			SEPARATOR ', ')
		FROM
			subjects_of_newclasses ncs
		WHERE
			(ncs.id_newclass = newclasses.id)) AS subjects,
	(SELECT 
		GROUP_CONCAT((SELECT ctg.name
			FROM
				categories ctg
			WHERE
				(ctg.id = ncctg.id_category))
			SEPARATOR ', ')
		FROM
			categories_of_newclasses ncctg
		WHERE
			(ncctg.id_newclass = newclasses.id)) AS categories`
}
