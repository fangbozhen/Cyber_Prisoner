package main

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

const SECRETKEY = "make QSC great again"

func generateToken(userid string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":  userid,
		"expired": time.Now().Unix() + 20,
	})
	tokenString, err := token.SignedString([]byte(SECRETKEY))
	if err != nil {
		fmt.Println("signed token error :", err)
		return ""
	}
	return tokenString
}

func parseToken(tokenString string) jwt.MapClaims {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(SECRETKEY), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims
	} else {
		fmt.Println("Parse token error: ", err)
		return nil
	}
}
