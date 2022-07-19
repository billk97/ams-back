package middlewares

import (
	"ams-back/utils"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"strings"
)

var jwtSecret = []byte("secret")

func JwtMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		jwtSecret = []byte(utils.Config.JwtSecret)
		authHeader := context.Request.Header.Get("Authorization")
		if authHeader == "" {
			apiError := utils.NewApiError("INVALID_CREDENTIALS",
				errors.New("authentication failed"),
				"This is a secure endpoint Bearer token is required",
			)
			apiError.Enhance(context)
			context.AbortWithStatusJSON(401, apiError)
			return
		}
		headerValue := strings.Split(authHeader, " ")
		if len(headerValue) < 2 {
			apiError := utils.NewApiError("INVALID_CREDENTIALS",
				errors.New("authentication failed"),
				"This is a secure endpoint Bearer token is required",
			)
			apiError.Enhance(context)
			context.AbortWithStatusJSON(401, apiError)
			return
		}
		token, err := jwt.Parse(headerValue[1], validateJwtToken)
		if err != nil {
			apiError := utils.NewApiError("INVALID_CREDENTIALS",
				errors.New("authentication failed"),
				"This is a secure endpoint Bearer token is required",
			)
			apiError.Enhance(context)
			context.AbortWithStatusJSON(401, apiError)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok && token.Valid {
			apiError := utils.NewApiError("INVALID_CREDENTIALS",
				errors.New("authentication failed"),
				"This is a secure endpoint Bearer token is required",
			)
			apiError.Enhance(context)
			context.AbortWithStatusJSON(401, apiError)
			return
		}
		for _, v := range claims {
			fmt.Println(v)
		}
		context.Next()
	}
}

func validateJwtToken(token *jwt.Token) (interface{}, error) {
	_, ok := token.Method.(*jwt.SigningMethodHMAC)
	if !ok {
		return nil, errors.New("unexpected signing method")
	}
	return jwtSecret, nil
}
