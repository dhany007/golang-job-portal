package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dhany007/golang-job-portal/models"
	"github.com/dhany007/golang-job-portal/models/response"
	"github.com/dhany007/golang-job-portal/services/utils"
	"github.com/julienschmidt/httprouter"
)

var (
	key = utils.GetEnv("KEY_JWT", "secretkey") // make default key "secretkey" if not provided
)

func Authentication(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		bearerToken := r.Header.Get("Authorization")

		authHeader := strings.Split(bearerToken, "Bearer ")
		if len(authHeader) != 2 {
			log.Println("[middleware] [Authentication] while Split bearerToken")
			response.Result(w, response.ErrorTokenInvalid)
			return
		}

		jwtToken := authHeader[1]
		token, err := jwt.Parse(jwtToken, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("token invalid")
			}
			return []byte(key), nil
		})

		if err != nil {
			log.Println("[middleware] [Authentication] while jwt.Parse")
			response.Result(w, response.ErrorTokenInvalid)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok && !token.Valid {
			log.Println("[middleware] [Authentication] while token.Claims")
			response.Result(w, response.ErrorUnauthorized)
			return
		}

		now := time.Now().Unix()
		expired := claims["exp"].(float64)

		if now >= int64(expired) {
			log.Println("[middleware] [Authentication] while expired")
			response.Result(w, response.ErrorTokenExpired)
			return
		}

		ctx := context.WithValue(r.Context(), models.Authorization{}, claims)

		next(w, r.WithContext(ctx), p)
	}
}
