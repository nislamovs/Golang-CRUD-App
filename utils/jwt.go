package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateJWT(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte("secret"))
}

func VerifyJWT(tokenString string) (float64, error) {
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Unexpected signing method")
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return 0, errors.New("Cannot parse token : " + err.Error())
	}

	if !parsedToken.Valid {
		return 0, errors.New("Token is not valid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("Cannot cast claims to MapClaims")
	}

	//email := claims["email"].(string)
	userId, _ := claims["userId"].(float64)

	return userId, nil
}
