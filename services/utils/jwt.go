package utils

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/models/response"
)

var (
	key = GetEnv("KEY_JWT", "secretkey") // make default key "secretkey" if not provided
)

func GenerateToken(ttl time.Duration, id string, email string, role int) (token string, err error) {
	now := time.Now().Unix()
	exp := time.Now().Add(ttl * time.Minute).Unix()

	claims := jwt.MapClaims{
		"id":    id,
		"email": email,
		"role":  role,
		"iat":   now,
		"exp":   exp,
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = parseToken.SignedString([]byte(key))
	if err != nil {
		return "", err
	}

	return token, nil
}

func ValidateToken(tokenStr string) (result jwt.MapClaims, err error) {
	var (
		tokenInvalid error = errors.New(fmt.Sprint(response.ErrorTokenInvalid))
		tokenExpired error = errors.New(fmt.Sprint(response.ErrorTokenExpired))
	)

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, tokenInvalid
		}

		return []byte(key), nil
	})

	if err != nil {
		return nil, tokenInvalid
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return nil, tokenInvalid
	}

	now := time.Now().Unix()
	expired := claims["exp"].(float64)

	if now >= int64(expired) {
		return nil, tokenExpired
	}

	return claims, nil
}

func GetAuthorization(ctx context.Context) (result models.Authorization) {
	auth := ctx.Value(models.Authorization{}).(jwt.MapClaims)

	result = models.Authorization{
		ID:    auth["id"].(string),
		Email: auth["email"].(string),
		Role:  int(auth["role"].(float64)),
	}

	return
}
