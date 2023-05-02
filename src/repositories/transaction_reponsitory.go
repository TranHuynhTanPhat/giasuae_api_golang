package repositories

import (
	"giasuaeapi/src/entities"

	"gorm.io/gorm"
)

type TransRepository interface {
	InsertTrans(trans *entities.TransactionhistoriesReq) error
	FindAllTrans() []entities.Transactionhistories
	FindByIDAcc(id int) entities.Transactionhistories
	FilterTrans(key interface{}) []entities.Transactionhistories
	Statistics(statis *entities.Statistics, month string, year string, day string)
}

type transConnection struct {
	connection *gorm.DB
}

// Statistics implements TransRepository
func (db *transConnection) Statistics(statis *entities.Statistics, month string, year string, day string) {
	var monthType int64
	var yearType int64
	var numNewclassReceived int64
	var numNewclass int64
	var percentNewclass float64
	var countTrans int64
	db.connection.Table("transactionhistories").Select("sum(amount)").Where("created_at LIKE ?", year+"%-"+month+"-%").Row().Scan(&monthType)
	db.connection.Table("transactionhistories").Select("sum(amount)").Where("created_at LIKE ?", year+"%").Row().Scan(&yearType)

	db.connection.Table("newclasses").Select("COUNT(status)").Count(&numNewclass)
	db.connection.Table("newclasses").Select("COUNT(status)").Where("status = 1").Count(&numNewclassReceived)
	if numNewclass == 0 {
		percentNewclass = 0
	} else {
		percentNewclass = float64(numNewclassReceived) / float64(numNewclass) * 100
	}
	db.connection.Table("transactionhistories").Select("COUNT(id)").Where("created_at LIKE ?", "%"+year+"-"+month+"-"+day+"%").Count(&countTrans)
	statis.Budget_month = int(monthType)
	statis.Budget_year = int(yearType)
	statis.Status_newclass = int(percentNewclass)
	statis.Count_trans = int(countTrans)
}

// FilterTrans implements TransRepository
func (db *transConnection) FilterTrans(key interface{}) []entities.Transactionhistories {
	var transs []entities.Transactionhistories
	db.connection.Table("transactionhistories").Select("trans.id, trans.amount, trans.content, trans.status, trans.created_at, accounts.username AS id_account ").Joins("trans left join accounts on accounts.id = trans.id_account").Where("trans.id LIKE ? OR trans.amount LIKE ? OR trans.content LIKE ? OR trans.status LIKE ? OR trans.created_at LIKE ? OR accounts.username LIKE ? ", key, key, key, key, key, key).Find(&transs)
	return transs
}

// FindAllTrans implements TransRepository
func (db *transConnection) FindAllTrans() []entities.Transactionhistories {
	var transs []entities.Transactionhistories
	db.connection.Table("transactionhistories").Select("trans.id, trans.amount, trans.content, trans.status, trans.created_at, accounts.username AS id_account ").Joins("trans left join accounts on accounts.id = trans.id_account").Find(&transs)
	return transs
}

// FindByIDAcc implements TransRepository
func (db *transConnection) FindByIDAcc(id int) entities.Transactionhistories {
	var trans entities.Transactionhistories
	db.connection.Where("id_account = ?", id).First(&trans)
	return trans
}

// InsertTrans implements TransRepository
func (db *transConnection) InsertTrans(trans *entities.TransactionhistoriesReq) error {
	err := db.connection.Table("transactionhistories").Save(&trans)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func NewTransRepository(dbConn *gorm.DB) TransRepository {
	return &transConnection{
		connection: dbConn,
	}
}
