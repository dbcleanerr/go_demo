package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	signingKey := []byte("welcome") // 加密key
	claims := MyClaims{
		Username: "zyy",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 2*60*60, // 失效时间
			Issuer:    "zyy",                       // 签发人
			NotBefore: time.Now().Unix() - 60,      // 生效时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(signingKey)
	fmt.Printf("%v,%v\n", ss, err)

	fmt.Println("parse")
	token2, err := jwt.ParseWithClaims(ss, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	fmt.Printf("%v,%v\n", token2, err)

	fmt.Println(token2.Claims.(*MyClaims).ExpiresAt)
	fmt.Println(token2.Claims.(*MyClaims).Username)
}
