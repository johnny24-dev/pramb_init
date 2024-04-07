package utils

import (
	"errors"
	"fmt"
	"time"

	"auth_service/pkg/domain"

	"github.com/golang-jwt/jwt"
)

var w *domain.JwtWrapper

func ValidateToken(signedToken string) (claims *domain.JwtClaims, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&domain.JwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		fmt.Println("error :", err)
		return
	}

	claims, ok := token.Claims.(*domain.JwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil
}

func GenerateToken(user domain.User) (signedToken string, err error) {
	claims := &domain.JwtClaims{
		StandardClaims: jwt.StandardClaims{ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix(), Issuer: w.Issuer},
		Userid:         user.Id,
		Email:          user.Email,
		Source:         "",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}
