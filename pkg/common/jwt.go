package common

import (
	"errors"
	"time"

	"github.com/arashi87/gin-template/pkg/setting"
	"github.com/dgrijalva/jwt-go"
)

type CustomClaims struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenToken(name, email string, id uint) (string, error) {
	// Get token expire time and jwt secret from env
	expireTime := time.Hour * time.Duration(setting.CONFIG.JWTExpire)
	secret := []byte(setting.CONFIG.JWTSecret)

	// custom claim, will use to return if token parse success
	claim := CustomClaims{id, name, email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expireTime).Unix(),
			Issuer:    setting.CONFIG.JWTIssuer,
		},
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(secret)
}

func ParseToken(token string) (*CustomClaims, error) {
	// parse token
	secret := []byte(setting.CONFIG.JWTSecret)
	result, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(secret), nil
	})

	// catch error
	if err != nil {
		return nil, err
	}

	// check token is valid
	if claims, ok := result.Claims.(*CustomClaims); ok && result.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}
