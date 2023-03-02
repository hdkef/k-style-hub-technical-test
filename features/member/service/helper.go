package service

import (
	"errors"
	"kstylehub/features/member"
)

func validateENUM(m member.Member) error {
	if m.SkinColor != "Type 1" && m.SkinColor != "Type 2" && m.SkinColor != "Type 3" && m.SkinColor != "Type 4" && m.SkinColor != "Type 5" && m.SkinColor != "Type 6" {
		return errors.New("make sure skincolor is between Type 1 - Type 2")
	}
	if m.Gender != "Female" && m.Gender != "Male" {
		return errors.New("gender can only be Male or Female")
	}
	if m.SkinType != "Normal" && m.SkinType != "Dry" && m.SkinType != "Oily" {
		return errors.New("skintype can only be Normal, Dry or Oily")
	}
	return nil
}
