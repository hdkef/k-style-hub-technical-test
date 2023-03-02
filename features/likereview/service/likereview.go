package service

import (
	"errors"
	"kstylehub/features/likereview"

	"github.com/go-playground/validator/v10"
)

type LikeReviewService struct {
	repo      likereview.LikeReviewRepository_
	validator validator.Validate
}

// Delete implements likereview.LikeReviewService_
func (u *LikeReviewService) Delete(idReview, idMember int) error {
	//validate id_member && id_review
	if idReview <= 0 && idMember <= 0 {
		return errors.New("invalid id_member or id_review")
	}
	return u.repo.Delete(idReview, idMember)
}

// Upsert implements likereview.LikeReviewService_
func (u *LikeReviewService) Upsert(idReview, idMember int) error {
	//validate id_member && id_review
	if idReview <= 0 && idMember <= 0 {
		return errors.New("invalid id_member or id_review")
	}
	return u.repo.Upsert(idReview, idMember)
}

func NewLikeReviewService(repo likereview.LikeReviewRepository_, validator validator.Validate) likereview.LikeReviewService_ {
	return &LikeReviewService{
		repo:      repo,
		validator: validator,
	}
}
