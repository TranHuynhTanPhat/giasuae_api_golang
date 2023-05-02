package repositories

import (
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	InsertCategory(ctg *entities.Category) error
	UpdateCategory(ctg *entities.Category) error
	DeleteCategory(id int) error
	FindAllCategory() []entities.Category
	FindByID(id int) entities.Category
	FilterCategory(value ...interface{}) []entities.Category
}

type categoryConnection struct {
	connection *gorm.DB
}

// FilterCategory implements CategoryRepository
func (db *categoryConnection) FilterCategory(value ...interface{}) []entities.Category {
	var ctgs []entities.Category
	db.connection.Where("type = ?", value[0]).Find(&ctgs)
	return ctgs
}

// DeleteCategory implements CategoryRepository
func (db *categoryConnection) DeleteCategory(id int) error {
	if err := db.connection.Table("categories_of_newclasses").Where("id_category = ?", id).Delete(&entities.CategoriesOfNewclasses{}); err.Error != nil {
		return err.Error
	}

	if err := db.connection.Table("categories_of_tutors").Where("id_category = ?", id).Delete(&entities.CategoriesOfTutor{}); err.Error != nil {
		return err.Error
	}
	if err := db.connection.Table("salaryinfos").Where("id_category = ?", id).Delete(&entities.Salaryinfo{}); err.Error != nil {
		return err.Error
	}

	if err := db.connection.Table("categories").Delete(&entities.Category{}, id); err.Error != nil {
		return err.Error
	}
	return nil
}

// FindAllCategory implements CategoryRepository
func (db *categoryConnection) FindAllCategory() []entities.Category {
	var categories []entities.Category
	db.connection.Find(&categories)
	return categories
}

// FindByID implements CategoryRepository
func (db *categoryConnection) FindByID(id int) entities.Category {
	var category entities.Category
	db.connection.First(&category, id)
	return category
}

// InsertCategory implements CategoryRepository
func (db *categoryConnection) InsertCategory(ctg *entities.Category) error {
	err := db.connection.Save(&ctg)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// UpdateCategory implements CategoryRepository
func (db *categoryConnection) UpdateCategory(ctg *entities.Category) error {
	err := db.connection.Save(&ctg)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func NewCategoryRepository(dbConn *gorm.DB) CategoryRepository {
	return &categoryConnection{
		connection: dbConn,
	}
}
