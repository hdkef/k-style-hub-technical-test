package product

import (
	"kstylehub/features/review"

	"github.com/gin-gonic/gin"
)

type Product struct {
	ID     int
	Name   string
	Price  float64
	Review []review.Review
}

type ProductResponse struct {
	ID     int                     `json:"id_product"`
	Name   string                  `json:"name_product"`
	Price  float64                 `json:"price"`
	Review []review.ReviewResponse `json:"review"`
}

type ProductDelivery_ interface {
	GetOneByID(c *gin.Context)
}

type ProductService_ interface {
	GetOneByID(id int) (Product, error)
}

type ProductRepository_ interface {
	GetOneByID(id int) (Product, error)
}
