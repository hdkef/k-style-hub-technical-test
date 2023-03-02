package delivery

import (
	"kstylehub/features/likereview"
	"kstylehub/features/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeReviewDelivery struct {
	service likereview.LikeReviewService_
}

// Delete implements likereview.LikeReviewDelivery_
func (u *LikeReviewDelivery) Delete(c *gin.Context) {
	//get id_review from param
	idReview, err := utils.ExtractParamID(c)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	//get id_member from jwt
	idMember, err := utils.ExtractJWTMemberID(c)
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, err.Error())
		return
	}
	err = u.service.Delete(idReview, idMember)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResult(c, "", "deleted")
}

// Upsert implements likereview.LikeReviewDelivery_
func (u *LikeReviewDelivery) Upsert(c *gin.Context) {
	//get id_review from param
	idReview, err := utils.ExtractParamID(c)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	//get id_member from jwt
	idMember, err := utils.ExtractJWTMemberID(c)
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, err.Error())
		return
	}
	err = u.service.Upsert(idReview, idMember)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResult(c, "", "upserted")
}

func NewLikeReviewDelivery(service likereview.LikeReviewService_) likereview.LikeReviewDelivery_ {
	return &LikeReviewDelivery{
		service: service,
	}
}
