package delivery

import (
	"kstylehub/features/member"
	"kstylehub/features/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MemberDelivery struct {
	service member.MemberService_
}

// Delete implements member.MemberDelivery_
func (u *MemberDelivery) Delete(c *gin.Context) {
	//get id from param
	id, err := utils.ExtractParamID(c)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	//get id from jwt
	jwtID, err := utils.ExtractJWTMemberID(c)
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, err.Error())
		return
	}
	err = u.service.Delete(id, jwtID)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResult(c, "", "deleted")
}

// GetAll implements member.MemberDelivery_
func (u *MemberDelivery) GetAll(c *gin.Context) {
	//get page from param
	pageStr, valid := c.GetQuery("page")
	if !valid {
		utils.SendError(c, http.StatusBadRequest, "page queryparam is needed")
		return
	}
	//get limit from param
	limitStr, valid := c.GetQuery("limit")
	if !valid {
		utils.SendError(c, http.StatusBadRequest, "limit queryparam is needed")
		return
	}
	//convert page & limit
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid page")
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "invalid limit")
		return
	}
	data, err := u.service.GetAll(page, limit)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResult(c, convertToResponseList(data), "")
}

// Insert implements member.MemberDelivery_
func (u *MemberDelivery) Insert(c *gin.Context) {
	//get payload
	var data member.MemberRequest
	err := c.Bind(&data)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "unknown payload")
		return
	}
	err = u.service.Insert(convertToCore(&data))
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResult(c, "", "created")
}

// Update implements member.MemberDelivery_
func (u *MemberDelivery) Update(c *gin.Context) {
	//get id from param
	id, err := utils.ExtractParamID(c)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	//get id from jwt
	jwtID, err := utils.ExtractJWTMemberID(c)
	if err != nil {
		utils.SendError(c, http.StatusUnauthorized, err.Error())
		return
	}
	//get payload
	var data member.MemberRequest
	err = c.Bind(&data)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, "unknown payload")
		return
	}
	//update
	err = u.service.Update(id, jwtID, convertToCore(&data))
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResult(c, "", "updated")
}

func NewMemberDelivery(service member.MemberService_) member.MemberDelivery_ {
	return &MemberDelivery{
		service: service,
	}
}
