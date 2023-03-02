package service

import (
	"errors"
	"kstylehub/features/member"

	"github.com/go-playground/validator/v10"
)

type MemberService struct {
	repo      member.MemberRepository_
	validator validator.Validate
}

// Delete implements member.MemberService_
func (u *MemberService) Delete(id int, jwtID int) error {
	//validate id and jwtID
	if id <= 0 && jwtID <= 0 {
		return errors.New("invalid id or jwt id")
	}
	//if jwtID != id reject
	if jwtID != id {
		return errors.New("unauthorized request to delete")
	}
	return u.repo.Delete(id)
}

// GetAll implements member.MemberService_
func (u *MemberService) GetAll(page int, limit int) ([]member.Member, error) {
	//validate page and limit
	if page <= 0 && limit <= 1 {
		return nil, errors.New("invalid page or limit")
	}
	return u.repo.GetAll(page, limit)
}

// Insert implements member.MemberService_
func (u *MemberService) Insert(member member.Member) error {
	//validate
	err := u.validator.Struct(member)
	if err != nil {
		return errors.New("something is missing")
	}
	err = validateENUM(member)
	if err != nil {
		return err
	}
	//insert
	return u.repo.Insert(member)
}

// Update implements member.MemberService_
func (u *MemberService) Update(id int, jwtID int, member member.Member) error {
	//validate id and jwtID
	if id <= 0 && jwtID <= 0 {
		return errors.New("invalid id or jwt id")
	}
	//if jwt id is not equal id reject
	if jwtID != id {
		return errors.New("unauthorized request to update")
	}

	//validate
	err := u.validator.Struct(member)
	if err != nil {
		return errors.New("something is missing")
	}
	err = validateENUM(member)
	if err != nil {
		return err
	}
	return u.repo.Update(id, member)
}

func NewMemberService(repo member.MemberRepository_, validator validator.Validate) member.MemberService_ {
	return &MemberService{
		repo:      repo,
		validator: validator,
	}
}
