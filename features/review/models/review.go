package models

import (
	"database/sql"
	"kstylehub/features/member"
)

type Review struct {
	IDReview  int
	IDProduct int
	TotalLike int
	IDMember  int
	Desc      sql.NullString
	Member    member.Member
}
