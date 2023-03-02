package review

import "kstylehub/features/member"

type Review struct {
	IDReview  int
	IDProduct int
	TotalLike int
	IDMember  int
	Desc      string
	Member    member.Member
}

type ReviewResponse struct {
	IDReview  int    `json:"id_review"`
	IDProduct int    `json:"id_product"`
	TotalLike int    `json:"total_like"`
	IDMember  int    `json:"id_member"`
	Desc      string `json:"desc"`
	Member    member.MemberResponse
}

type ReviewRepository_ interface {
	GetAllByProductID(id int) ([]Review, error)
}
