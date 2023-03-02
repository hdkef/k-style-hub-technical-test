package delivery

import (
	"kstylehub/features/product"
	"kstylehub/features/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductDelivery struct {
	service product.ProductService_
}

// GetOneByID implements product.ProductDelivery_
func (u *ProductDelivery) GetOneByID(c *gin.Context) {
	//get id from param
	id, err := utils.ExtractParamID(c)
	if err != nil {
		utils.SendError(c, http.StatusBadRequest, err.Error())
		return
	}
	data, err := u.service.GetOneByID(id)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendResult(c, convertToResponse(&data), "")
}

func NewProductDelivery(service product.ProductService_) product.ProductDelivery_ {
	return &ProductDelivery{
		service: service,
	}
}
