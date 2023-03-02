package member

import "github.com/gin-gonic/gin"

type Member struct {
	ID        int
	Username  string `validate:"required"`
	Gender    string `validate:"required"`
	SkinType  string `validate:"required"`
	SkinColor string `validate:"required"`
}

type MemberRequest struct {
	Username  string
	Gender    string
	SkinType  string
	SkinColor string
}

type MemberResponse struct {
	ID        int
	Username  string
	Gender    string
	SkinType  string
	SkinColor string
}

type MemberDelivery_ interface {
	GetAll(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Insert(c *gin.Context)
}

type MemberService_ interface {
	GetAll(page int, limit int) ([]Member, error)
	Update(id int, jwtID int, member Member) error
	Delete(id int, jwtID int) error
	Insert(member Member) error
}

type MemberRepository_ interface {
	GetOneByID(id int) (Member, error)
	GetAll(page int, limit int) ([]Member, error)
	Update(id int, member Member) error
	Delete(id int) error
	Insert(member Member) error
}
