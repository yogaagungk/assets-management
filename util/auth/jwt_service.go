package auth

import (
	"fmt"
	"log"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

// all documented for this file is
// from https://godoc.org/github.com/dgrijalva/jwt-go

const signingKey = "this-secret-key-32334"
const expiredIn = 216000

func GenerateToken(userAuth UserAuth) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiredIn,
		Issuer:    "my-asset-application",
		Subject:   userAuth.Username,
	})

	tokenString, _ := token.SignedString([]byte(signingKey))

	log.Println("Generated Token for user " + userAuth.Username)

	return tokenString
}

func ParseToken(tokenString string) string {
	token, _ := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(signingKey), nil
	})

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok {
		log.Println(claims.Subject)

		return claims.Subject
	} else {
		log.Println(claims.Subject + " - " + strconv.FormatBool(ok) + " - " + strconv.FormatBool(token.Valid))

		return ""
	}
}
