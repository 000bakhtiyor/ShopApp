package config

import (
	"fmt"
	"shopapp/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


var secretKey = []byte("salom-dunyo")

func CreateToken (user models.User) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{
			"username": user.Username,
			"user_id": user.UserID,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		},
	)

	tokenString, err := token.SignedString(secretKey)
    if err != nil {
    	return "", err
    }

 	return tokenString, nil
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token)   (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v",token.Header["alg"])
			}
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}