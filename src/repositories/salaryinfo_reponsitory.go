package repositories

import (
	"giasuaeapi/src/entities"

	"github.com/mashingan/smapping"
	"gorm.io/gorm"
)

type SalaryinfoRepository interface {
	InsertSalaryinfo(sal *entities.Salaryinfo) error
	UpdateSalaryinfo(sal *entities.Salaryinfo) error
	DeleteSalaryinfo(id int) error
	FindAllSalaryinfo() []entities.SalaryinfoView
	FindByID(id int) entities.SalaryinfoDetail
	FindByType(type_teacher int) []entities.SalaryinfoView
}

type salaryinfoConnection struct {
	connection *gorm.DB
}

// FindByType implements SalaryinfoRepository
func (db *salaryinfoConnection) FindByType(type_teacher int) []entities.SalaryinfoView {
	var sal []entities.SalaryinfoView
	db.connection.Table("salaryinfos").Select("sal.id, sal.type_teacher, sal.two_sessions, sal.three_sessions, sal.four_sessions, sal.five_sessions, sal.created_at, categories.name AS id_category ").Joins("sal left join categories on categories.id = sal.id_category").Where("type_teacher = ?", type_teacher).Find(&sal)
	return sal
}

// DeleteSalaryinfo implements SalaryinfoRepository
func (db *salaryinfoConnection) DeleteSalaryinfo(id int) error {
	err := db.connection.Delete(&entities.Salaryinfo{}, id)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// FindAllSalaryinfo implements SalaryinfoRepository
func (db *salaryinfoConnection) FindAllSalaryinfo() []entities.SalaryinfoView {
	var sals []entities.SalaryinfoView
	// db.connection.Table("salaryinfos").Select("sal.id, sal.type_teacher, sal.two_sessions, sal.three_sessions, sal.four_sessions, sal.five_sessions, sal.created_at, categories.name AS id_category ").Joins("sal left join categories on categories.id = sal.id_category").Find(&sals)
	db.connection.Table("salaryinfos").Select("sal.id, sal.type_teacher, sal.two_sessions, sal.three_sessions, sal.four_sessions, sal.five_sessions, sal.created_at, categories.name AS id_category ").Joins("sal left join categories on categories.id = sal.id_category").Find(&sals)
	return sals
}

// FindByID implements SalaryinfoRepository
func (db *salaryinfoConnection) FindByID(id int) entities.SalaryinfoDetail {
	var sal entities.Salaryinfo
	var salDetail entities.SalaryinfoDetail
	var ctg entities.Category
	db.connection.Limit(1).Table("salaryinfos").Where("id = ?", id).Scan(&sal)
	smapping.FillStruct(&salDetail, smapping.MapFields(&sal))
	db.connection.First(&ctg, sal.ID_category)
	salDetail.ID_category = ctg
	return salDetail
}

// InsertSalaryinfo implements SalaryinfoRepository
func (db *salaryinfoConnection) InsertSalaryinfo(sal *entities.Salaryinfo) error {
	err := db.connection.Table("salaryinfos").Create(&sal)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

// UpdateSalaryinfo implements SalaryinfoRepository
func (db *salaryinfoConnection) UpdateSalaryinfo(sal *entities.Salaryinfo) error {
	err := db.connection.Save(&sal)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func NewSalaryinfoRepository(dbConn *gorm.DB) SalaryinfoRepository {
	return &salaryinfoConnection{
		connection: dbConn,
	}
}
