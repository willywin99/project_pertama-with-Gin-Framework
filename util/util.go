package util

import (
	"errors"
	"project_pertama/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func CreateResponse(isSuccess bool, data any, errorMessage string) model.Response {
	return model.Response{
		Success: isSuccess,
		Data:    data,
		Error:   errorMessage,
	}
}

func Hash(data []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(data, 8)
}

func HashMatched(hash []byte, plain []byte) bool {
	err := bcrypt.CompareHashAndPassword(hash, plain)
	return err == nil
}

func GenerateJWTToken(isAdmin bool, userId string) (string, error) {
	claims := jwt.MapClaims{
		"admin": isAdmin,
		"sub":   userId,
		"exp":   time.Now().Add(5 * time.Minute).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("s4y4_5uka_9074nG"))
}

func GetJWTClaims(tokenString string) (map[string]any, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid method")
		}
		return []byte("s4y4_5uka_9074nG"), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
}

func GetSubFromClaims(claims any) (any, error) {
	mapClaims, ok := claims.(map[string]any)
	if !ok {
		return nil, errors.New("not map")
	}

	sub, ok := mapClaims["sub"]
	if !ok {
		return nil, errors.New("not found")
	}

	return sub, nil
}
