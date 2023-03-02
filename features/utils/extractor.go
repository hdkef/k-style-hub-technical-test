package utils

import (
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func ExtractParamID(c *gin.Context) (int, error) {
	idStr := c.Params.ByName("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid id param")
	}
	return id, nil
}

func ExtractJWTMemberID(c *gin.Context) (int, error) {
	jwtCtx, exist := c.Get("jwt")
	if !exist {
		return 0, errors.New("jwt not found")
	}
	jwtData, valid := jwtCtx.(*jwt.Token)
	if valid && jwtData.Valid {
		claims := jwtData.Claims.(jwt.MapClaims)
		ID, v := claims["id"].(float64)
		if v {
			return int(ID), nil
		}
		return 0, errors.New("jwt not found")
	}
	return 0, errors.New("jwt not found")
}
