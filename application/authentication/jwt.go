package authentication

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
)

type jsonWebToken struct{}

type JWTDecoder interface {
	Decode(string) (jwt.MapClaims, error)
}

func (j jsonWebToken) Decode(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	} else {
		fmt.Println(err)

		return nil, errors.New("error decoding JWT")
	}

}

var (
	JsonWebToken = jsonWebToken{}
)