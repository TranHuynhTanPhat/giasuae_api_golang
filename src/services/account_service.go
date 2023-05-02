package services

import (
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
)

type AccountService interface {
	InsertAccount(acc *entities.Account)
	UpdateAccount(acc *entities.Account)
	DeleteAccount(acc *entities.Account)
	FindAllAccount() []entities.Account
	FindByID(id int) entities.Account
	FilterAccount(username string, isTutor int) []entities.Account
	UpdatePassword(pass string, id int) error
}
type accountService struct {
	AccountReponsitory repositories.AccountReponsitory
}

// UpdatePassword implements AccountService
func (svc *accountService) UpdatePassword(pass string, id int) error {
	return svc.AccountReponsitory.UpdatePassword(pass, id)
}

// FilterAccount implements AccountService
func (svc *accountService) FilterAccount(username string, isTutor int) []entities.Account {
	return svc.AccountReponsitory.FilterAccount(username, isTutor)
}

// DeleteAccount implements AccountService
func (accsv *accountService) DeleteAccount(acc *entities.Account) {
	accsv.AccountReponsitory.DeleteAccount(acc)
}

// FindAllAccount implements AccountService
func (accsv *accountService) FindAllAccount() []entities.Account {
	return accsv.AccountReponsitory.FindAllAccount()
}

// FindByID implements AccountService
func (accsv *accountService) FindByID(id int) entities.Account {
	return accsv.AccountReponsitory.FindByID(id)
}

// InsertAccount implements AccountService
func (accsv *accountService) InsertAccount(acc *entities.Account) {
	accsv.AccountReponsitory.InsertAccount(acc)
}

// UpdateAccount implements AccountService
func (accsv *accountService) UpdateAccount(acc *entities.Account) {
	accsv.AccountReponsitory.UpdateAccount(acc)
}

func NewAccountService(accRepo repositories.AccountReponsitory) AccountService {
	return &accountService{
		AccountReponsitory: accRepo,
	}
}
