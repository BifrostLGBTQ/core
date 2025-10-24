package helpers

import (
	"bifrost/models/jwtclaims"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

func GenerateUserJWT(user_id uuid.UUID, publicId int64) (string, error) {
	var jwtSecret = []byte(os.Getenv("USER_JWT_SECRET"))

	claims := jwt.MapClaims{
		"user_id":  user_id,
		"publicId": publicId,
		"exp":      time.Now().AddDate(0, 0, 30).Unix(),
		"version":  "0.1.2",
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, error := token.SignedString(jwtSecret)
	result := fmt.Sprintf("Bearer %s", tokenString)
	return result, error
}

func DecodeUserJWT(tokenString string) (*jwtclaims.UserJWTClaims, error) {
	fmt.Println("DecodeUserJWT:Token:", tokenString)
	token, err := jwt.ParseWithClaims(tokenString, &jwtclaims.UserJWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("USER_JWT_SECRET")), nil
	})
	if err != nil {
		fmt.Println("DecodeUserJWT:Error:1", err)
		return nil, err
	}
	if !token.Valid {
		fmt.Println("DecodeUserJWT:Error:2")
		return nil, errors.New("invalid jwt token")
	}
	myClaims, ok := token.Claims.(*jwtclaims.UserJWTClaims)
	if !ok {
		fmt.Println("DecodeUserJWT:Error:3")
		return nil, errors.New("couldn't parse token claims")
	}
	fmt.Println("DecodeUserJWT:Passed:1")
	return myClaims, nil
}
