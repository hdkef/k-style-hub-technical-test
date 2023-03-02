package dummylogin

import (
	"database/sql"
	"kstylehub/app/config"
	"kstylehub/features/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type DummyLogin struct {
	db     *sql.DB
	config *config.AppConfig
}

type Login_ interface {
	Handler(c *gin.Context)
}

func NewDummyLogin(db *sql.DB, config *config.AppConfig) Login_ {
	return &DummyLogin{
		db:     db,
		config: config,
	}
}

func (u *DummyLogin) Handler(c *gin.Context) {
	var payload struct {
		IDMember int `json:"id_member"`
		Exp      int `json:"exp"`
	}

	err := c.Bind(&payload)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "unknown payload")
		return
	}

	if payload.IDMember <= 0 {
		utils.SendError(c, http.StatusBadRequest, "invalid id member")
		return
	}

	if payload.Exp <= 0 {
		utils.SendError(c, http.StatusBadRequest, "expired min 1")
		return
	}

	//go to database if member exists
	var id int
	err = u.db.QueryRow("SELECT id_member FROM member WHERE id_member = ?", payload.IDMember).
		Scan(&id)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "member does not exist")
		return
	}
	//generate jwt
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":      id,
		"expired": time.Now().Add(time.Duration(payload.Exp * int(time.Hour))).Unix(),
	})

	tokenString, err := token.SignedString([]byte(u.config.JWT_SECRET))
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "fail to signed token")
		return
	}

	utils.SendResult(c, tokenString, "logged in")
}
