package repository

import (
	"kstylehub/features/review"
	"kstylehub/features/review/models"
)

func convertToCore(u *models.Review) review.Review {
	return review.Review{
		IDReview:  u.IDReview,
		IDProduct: u.IDProduct,
		TotalLike: u.TotalLike,
		IDMember:  u.IDMember,
		Desc:      u.Desc.String,
		Member:    u.Member,
	}
}

func convertToCoreList(u []models.Review) []review.Review {
	data := []review.Review{}
	for _, v := range u {
		data = append(data, convertToCore(&v))
	}
	return data
}
