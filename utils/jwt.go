package utils

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecretKey = os.Getenv("JWT_SECRET")

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
		"iat":    time.Now().Unix(),
	})
	return token.SignedString([]byte(jwtSecretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token " + err.Error())
	}

	isTokenValid := parsedToken.Valid
	if !isTokenValid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	//* we can extract values from claims like this if needed
	// email, ok := claims["email"].(string)
	userId, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	return int64(userId), nil
}

//TODO: add separate fucntion to get userId from verified token