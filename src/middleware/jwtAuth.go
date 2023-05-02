package middleware

import (
	"giasuaeapi/src/helper"
	"giasuaeapi/src/services"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorJWT(jwts services.JWTService) gin.HandlerFunc {
	return func(cx *gin.Context) {
		authHeader := cx.GetHeader("token")
		if authHeader == "" {
			response := helper.BuildResponseError("Yêu cầu thất bại", "Không tìm thấy token", nil)
			cx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
		token, err := jwts.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claims[user_is]: ", claims["user_id"])
			log.Println("Claims[issuer]: ", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildResponseError("Token không hợp lệ", err.Error(), nil)
			cx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
