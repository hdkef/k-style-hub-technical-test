package likereview

import "github.com/gin-gonic/gin"

type LikeReview struct {
	IDReview int `validate:"required"`
	IDMember int `validate:"required"`
}

type LikeReviewDelivery_ interface {
	Upsert(c *gin.Context)
	Delete(c *gin.Context)
}

type LikeReviewService_ interface {
	Upsert(idReview, idMember int) error
	Delete(idReview, idMember int) error
}

type LikeReviewRepository_ interface {
	Upsert(idReview, idMember int) error
	Delete(idReview, idMember int) error
}
