package repositories

import (
	"fmt"
	"giasuaeapi/src/entities"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type TutorRepository interface {
	InsertTutor(tutor *entities.TutorReq) error
	UpdateTutor(tutor *entities.TutorReq) error
	DeleteTutor(id int) error
	FindAllTutor() []entities.TutorSet
	FindByID(id int) entities.TutorDetail
	FilterTutor(subID int, classID int, cateID int, gender string, isnow string) []entities.TutorSet
}

type tutorConnection struct {
	connection *gorm.DB
}

// FilterTutor implements TutorRepository
func (db *tutorConnection) FilterTutor(subID int, classID int, cateID int, gender string, isnow string) []entities.TutorSet {
	var tutors []entities.TutorSet
	db.connection.Table("tutors").
		Joins("INNER JOIN subjects_of_tutors ON tutors.id IN (SELECT id_tutor FROM subjects_of_tutors WHERE subjects_of_tutors.id_subject = ?) OR ? = 0 ", subID, subID).
		Joins("INNER JOIN classes_of_tutors ON tutors.id IN (SELECT id_tutor FROM classes_of_tutors WHERE classes_of_tutors.id_class = ?) OR ? = 0 ", classID, classID).
		Joins("INNER JOIN categories_of_tutors ON tutors.id IN (SELECT id_tutor FROM categories_of_tutors WHERE categories_of_tutors.id_category = ?) OR ? = 0 ", cateID, cateID).
		// Joins("INNER JOIN tutors ON tutorS.gender LIKE ?", gender).
		Select(queyGetAllTutor()).
		Where("(tutors.gender = ? OR ? = '') AND (tutors.isnow = ? OR ? = '')", gender, gender, isnow, isnow).
		Group("tutors.id").
		Find(&tutors)
	return tutors
}

// DeleteTutor implements TutorRepository
func (db *tutorConnection) DeleteTutor(id int) error {
	if err := db.connection.Table("classes_of_tutors").Where("id_tutor = ?", id).Delete(&entities.ClassesOfTutor{}); err.Error != nil {
		return err.Error
	}
	if err := db.connection.Table("subjects_of_tutors").Where("id_tutor = ?", id).Delete(&entities.SubjectsOfTutor{}); err.Error != nil {
		return err.Error
	}
	if err := db.connection.Table("categories_of_tutors").Where("id_tutor = ?", id).Delete(&entities.CategoriesOfTutor{}); err.Error != nil {
		return err.Error
	}

	if err := db.connection.Table("tutors").Delete(&entities.TutorDefault{}, id); err.Error != nil {
		return err.Error
	}
	return nil
}

// UpdateTutor implements TutorRepository
func (db *tutorConnection) UpdateTutor(tutor *entities.TutorReq) error {
	delListClassOfTutor(db, tutor.ID)
	delListSubjectOfTutor(db, tutor.ID)
	delListCategoryOfTutor(db, tutor.ID)
	var tutorDF entities.TutorDefault
	var subOfTT []entities.SubjectsOfTutor
	var clasOfTT []entities.ClassesOfTutor
	var ctgOfTT []entities.CategoriesOfTutor
	smapping.FillStruct(&tutorDF, smapping.MapFields(&tutor))
	err := db.connection.Table("tutors").Save(&tutorDF)
	if err.Error != nil {
		return err.Error
	}
	for _, value := range tutor.Subjects {
		subOfTT = append(subOfTT, entities.SubjectsOfTutor{
			ID_tutor:   tutor.ID,
			ID_subject: value,
		})
	}
	for _, value := range tutor.Classes {
		clasOfTT = append(clasOfTT, entities.ClassesOfTutor{
			ID_tutor: tutor.ID,
			ID_class: value,
		})
	}
	for _, value := range tutor.Categories {
		ctgOfTT = append(ctgOfTT, entities.CategoriesOfTutor{
			ID_tutor:    tutor.ID,
			ID_category: value,
		})
	}

	err1 := db.connection.Table("subjects_of_tutors").Save(&subOfTT)
	err2 := db.connection.Table("classes_of_tutors").Save(&clasOfTT)
	err3 := db.connection.Table("categories_of_tutors").Save(&ctgOfTT)
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

// FindAllTutor implements TutorRepository
func (db *tutorConnection) FindAllTutor() []entities.TutorSet {
	var tutors []entities.TutorSet
	db.connection.Table("tutors").Select(queyGetAllTutor()).Group("tutors.id").Find(&tutors)
	return tutors
}

// FindByID implements TutorRepository
func (db *tutorConnection) FindByID(id int) entities.TutorDetail {
	var tutors entities.TutorDetail
	db.connection.Limit(1).Select("*, (SELECT accounts.username FROM accounts WHERE accounts.id = tutors.id_account) as id_account").Table("tutors").Where("id = ?", id).Scan(&tutors)
	tutors.Subjects = getListSubjectOfTutor(db, id)
	tutors.Classes = getListClassOfTutor(db, id)
tutors.Categories = getListCategoryOfTutor(db, id)
	return tutors
}

// InsertTutor implements TutorRepository
func (db *tutorConnection) InsertTutor(tutor *entities.TutorReq) error {
	var subOfTT []entities.SubjectsOfTutor
	var classOfTT []entities.ClassesOfTutor
	var ctgOfTT []entities.CategoriesOfTutor
	err := db.connection.Table("tutors").Create(&tutor)
	if err.Error != nil {
		return fmt.Errorf("Tài khoản này đã đăng ký gia sư")
	}
	for _, value := range tutor.Subjects {
		subOfTT = append(subOfTT, entities.SubjectsOfTutor{
			ID_tutor:   tutor.ID,
			ID_subject: value,
		})
	}
	for _, value := range tutor.Classes {
		classOfTT = append(classOfTT, entities.ClassesOfTutor{
			ID_tutor: tutor.ID,
			ID_class: value,
		})
	}
	for _, value := range tutor.Categories {
		ctgOfTT = append(ctgOfTT, entities.CategoriesOfTutor{
			ID_tutor:    tutor.ID,
			ID_category: value,
		})
	}

	db.connection.Create(&subOfTT)
	db.connection.Create(&classOfTT)
	db.connection.Create(&ctgOfTT)
	return nil
}

func NewTutorRepository(dbConn *gorm.DB) TutorRepository {
	return &tutorConnection{
		connection: dbConn,
	}
}

func getListSubjectOfTutor(db *tutorConnection, id int) []entities.Subject {
	var subjects []entities.Subject
	db.connection.Table("subjects").Joins("inner join subjects_of_tutors on id_tutor = ?", id).Where("id = id_subject").Scan(&subjects)
	return subjects
}
func getListClassOfTutor(db *tutorConnection, id int) []entities.Class {
	var classes []entities.Class
	db.connection.Table("classes").Joins("inner join classes_of_tutors on id_tutor = ?", id).Where("id = id_class").Scan(&classes)
	return classes
}

func getListCategoryOfTutor(db *tutorConnection, id int) []entities.Category {
	var categories []entities.Category
	db.connection.Table("categories").Joins("inner join categories_of_tutors on id_tutor = ?", id).Where("id = id_category").Scan(&categories)
	return categories
}

func delListSubjectOfTutor(db *tutorConnection, id int) {
	db.connection.Table("subjects_of_tutors").Where("id_tutor = ?", id).Delete(&entities.SubjectsOfTutor{})
}
func delListClassOfTutor(db *tutorConnection, id int) {
	db.connection.Table("classes_of_tutors").Where("id_tutor = ?", id).Delete(&entities.ClassesOfTutor{})
}
func delListCategoryOfTutor(db *tutorConnection, id int) {
	db.connection.Table("categories_of_tutors").Where("id_tutor = ?", id).Delete(&entities.CategoriesOfTutor{})
}

func queyGetAllTutor() string {
	return `tutors.id,
	tutors.name,
	tutors.address,
	tutors.email,
	tutors.phone,
	tutors.school,
	tutors.department,
	tutors.gender,
	tutors.graduate_year,
	tutors.isnow,
	tutors.describe,
	tutors.sobuoi,
	tutors.birth_year,
	(SELECT accounts.username FROM accounts WHERE accounts.id = tutors.id_account) as id_account,
	tutors.created_at,
	(SELECT 
		GROUP_CONCAT((SELECT c.name
			FROM
				classes c
			WHERE
				(c.id = ncc.id_class))
			SEPARATOR ', ')
		FROM
			classes_of_tutors ncc
		WHERE
			(ncc.id_tutor = tutors.id)) AS classes,
	(SELECT 
		GROUP_CONCAT((SELECT s.name
			FROM
				subjects s
			WHERE
				(s.id = ncs.id_subject))
			SEPARATOR ', ')
		FROM
			subjects_of_tutors ncs
		WHERE
			(ncs.id_tutor = tutors.id)) AS subjects,
	(SELECT 
		GROUP_CONCAT((SELECT ctg.name
			FROM
				categories ctg
			WHERE
				(ctg.id = cot.id_category))
			SEPARATOR ', ')
		FROM
			categories_of_tutors cot
		WHERE
			(cot.id_tutor = tutors.id)) AS categories`
}
