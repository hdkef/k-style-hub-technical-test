package repository

import (
	"database/sql"
	"errors"
	mr "kstylehub/features/member"
	"kstylehub/features/review"
	"kstylehub/features/review/models"
)

type ReviewRepository struct {
	db         *sql.DB
	memberRepo mr.MemberRepository_
}

// GetAllByProductID implements review.ReviewRepository_
func (u *ReviewRepository) GetAllByProductID(id int) ([]review.Review, error) {
	//get review
	var data []models.Review
	rows, err := u.db.Query("SELECT id_review, id_product, id_member, desc_review FROM review WHERE id_product = ?", id)
	if err != nil {
		return nil, errors.New("internal database error")
	}
	for rows.Next() {
		var review models.Review
		err = rows.Scan(&review.IDReview, &review.IDProduct, &review.IDMember, &review.Desc)
		if err != nil {
			return nil, errors.New("internal database error")
		}

		//get member
		member, err := u.memberRepo.GetOneByID(review.IDMember)
		if err != nil {
			return nil, errors.New("internal database error")
		}
		review.Member = member

		//get total like
		var totalLike int = 0
		err = u.db.QueryRow("SElECT COUNT(id_review) FROM like_review WHERE id_review = ?", review.IDReview).Scan(&totalLike)

		if err != nil {
			return nil, errors.New("internal database error")
		}

		review.TotalLike = totalLike

		data = append(data, review)
	}

	return convertToCoreList(data), nil
}

func NewReviewRepository(db *sql.DB, memberRepo mr.MemberRepository_) review.ReviewRepository_ {
	return &ReviewRepository{
		db:         db,
		memberRepo: memberRepo,
	}
}
