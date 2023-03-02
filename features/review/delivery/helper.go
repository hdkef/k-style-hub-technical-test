package delivery

import (
	helper "kstylehub/features/member/delivery"
	"kstylehub/features/review"
)

func ConvertToReviewResponse(u *review.Review) review.ReviewResponse {
	return convertToResponse(u)
}

func convertToResponse(u *review.Review) review.ReviewResponse {
	return review.ReviewResponse{
		IDReview:  u.IDReview,
		IDProduct: u.IDProduct,
		IDMember:  u.IDMember,
		TotalLike: u.TotalLike,
		Desc:      u.Desc,
		Member:    helper.ConvertToMemberResponse(&u.Member),
	}
}
