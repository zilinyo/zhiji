package models

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

type jwtCustomClaims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid   uint `json:"uid"`
	Admin bool `json:"admin"`
}

//生成token
func CreateToken(SecretKey []byte, issuer string, Uid uint, isAdmin bool) (tokenString string, err error) {
	claims := &jwtCustomClaims{
		jwt.StandardClaims{
			ExpiresAt: int64(time.Now().Add(time.Hour * 24).Unix()),
			Issuer:    issuer,
		},
		Uid,
		isAdmin,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(SecretKey)
	return
}

//解析token
func ParseToken(tokenSrt string, SecretKey []byte) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenSrt, func(*jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	claims = token.Claims
	return
}
