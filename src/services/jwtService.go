package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

// GenerateToken implements JWTService
func (jwts *jwtService) GenerateToken(userID string) string {
	claims := &jwtCustomClaim{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(1, 0, 0).Unix(),
			Issuer:    jwts.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tk, err := token.SignedString([]byte(jwts.secretKey))
	if err != nil {
		panic(err.Error())
	}
	return tk
}

// ValidateToken implements JWTService
// Xác thực token
func (jwts *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method %v", t.Header["alg"])
		}
		return []byte(jwts.secretKey), nil
	})
}

func NewJWTService() JWTService {
	return &jwtService{
		issuer:    os.Getenv("ISSUER"),
		secretKey: GetSecretKey("SECRET_ADMIN"),
	}
}

func GetSecretKey(key string) string {
	secretKey := os.Getenv(key)
	if secretKey == "" {
		secretKey = "default"
	}
	return secretKey
}
