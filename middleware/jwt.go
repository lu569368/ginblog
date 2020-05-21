package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Claims struct {
	Username string
	jwt.StandardClaims
}

var jwtkey = []byte("a_secret_create")

func Getting(ctx *gin.Context) {
	tokenString := ctx.GetHeader("Authorization")
	//vcalidate token formate
	if tokenString == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足1"})
		ctx.Abort()
		return
	}

	token, claims, err := ParseToken(tokenString)
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足2"})
		ctx.Abort()
		return
	}
	ctx.Set("username", claims.Username)
}

func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	Claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, Claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtkey, nil
	})
	return token, Claims, err
}
