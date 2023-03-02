package middleware

import (
	"fmt"
	"kstylehub/app/config"
	"kstylehub/features/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type JWTMiddleware struct {
	config *config.AppConfig
}

// Handle implements JWTMiddleware_
func (u *JWTMiddleware) Handle(next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		//get bearer token from header
		bearer := c.Request.Header.Get("AUTHORIZATION")
		if len(bearer) <= 7 {
			utils.SendError(c, http.StatusUnauthorized, "jwt invalid or does not exist")
			return
		}

		//remove BEARER_
		bearer = bearer[7:]

		//parse token to jwt
		jwt, err := parseTokenString(bearer, u.config.JWT_SECRET)
		if err != nil {
			utils.SendError(c, http.StatusUnauthorized, "jwt invalid")
			return
		}

		//set jwt to context
		c.Set("jwt", jwt)
		//continue
		next(c)
	}
}

type JWTMiddleware_ interface {
	Handle(next gin.HandlerFunc) gin.HandlerFunc
}

func NewJWTMiddleware(config *config.AppConfig) JWTMiddleware_ {
	return &JWTMiddleware{
		config: config,
	}
}

func parseTokenString(tokenString string, secret string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	return token, err
}
