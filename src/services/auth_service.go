package services

import (
	"fmt"
	"giasuaeapi/src/entities"
	"giasuaeapi/src/repositories"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	VerifyCredential(usernam string, password string) (interface{}, error)
	CreateUser(user *entities.Account) error
	// FindByEmail(email string) entity.User
	// IsDuplicateEmail(email string) bool
}

type authService struct {
	AcountRepository repositories.AccountReponsitory
}

// CreateUser implements AuthService
func (svc *authService) CreateUser(user *entities.Account) error {
	return svc.AcountRepository.InsertAccount(user)
}

// VerifyCredential implements AuthService
func (svc *authService) VerifyCredential(username string, password string) (interface{}, error) {
	res := svc.AcountRepository.VerifyCredential(username)
	if res == false {
		return nil, fmt.Errorf("Không tìm thấy tên tài khoản")
	}
	if acv, ok := res.(entities.AccountWithToken); ok {
		isComparePass := comparePass(acv.Password, []byte(password))
		if !isComparePass {
			return nil, fmt.Errorf("Sai mật khẩu")
		} else if acv.State == 0 {
			return nil, fmt.Errorf("Tài khoản bị khóa")
		}
	}
	return res, nil
}

func NewAuthService(accRepo repositories.AccountReponsitory) AuthService {
	return &authService{
		AcountRepository: accRepo,
	}
}

func comparePass(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
