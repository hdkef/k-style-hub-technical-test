package repository

import (
	"kstylehub/features/member"
	"kstylehub/features/member/models"
)

func convertToModels(u *member.Member) models.Member {
	return models.Member{
		ID:        u.ID,
		Username:  u.Username,
		SkinType:  u.SkinType,
		SkinColor: u.SkinColor,
		Gender:    u.Gender,
	}
}

func convertToCore(u *models.Member) member.Member {
	return member.Member{
		ID:        u.ID,
		Username:  u.Username,
		SkinType:  u.SkinType,
		SkinColor: u.SkinColor,
		Gender:    u.Gender,
	}
}

func convertToCoreList(u []models.Member) []member.Member {
	data := []member.Member{}
	for _, v := range u {
		data = append(data, convertToCore(&v))
	}
	return data
}
