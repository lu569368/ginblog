package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/student/ginblog/datamodels"
)

var jwtkey = []byte("a_secret_create")

type Claims struct {
	Username string
	jwt.StandardClaims
}

func ReleaseToken(user datamodels.LoginUser) (string, error) {
	expireTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(), //过期时间
			IssuedAt:  time.Now().Unix(),
			Issuer:    "127.0.0.1",
			Subject:   "user token",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtkey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
