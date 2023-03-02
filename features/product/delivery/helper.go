package delivery

import (
	"kstylehub/features/product"
	"kstylehub/features/review"
	helper "kstylehub/features/review/delivery"
)

func convertToResponse(u *product.Product) product.ProductResponse {
	data := product.ProductResponse{
		ID:    u.ID,
		Name:  u.Name,
		Price: u.Price,
	}

	var reviewResponse []review.ReviewResponse
	for _, v := range u.Review {
		reviewResponse = append(reviewResponse, helper.ConvertToReviewResponse(&v))
	}
	data.Review = reviewResponse
	return data
}
