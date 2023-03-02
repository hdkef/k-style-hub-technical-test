package delivery

import "kstylehub/features/member"

func ConvertToMemberResponse(u *member.Member) member.MemberResponse {
	return convertToResponse(u)
}

func convertToResponse(u *member.Member) member.MemberResponse {
	return member.MemberResponse{
		ID:        u.ID,
		Username:  u.Username,
		SkinType:  u.SkinType,
		SkinColor: u.SkinColor,
		Gender:    u.Gender,
	}
}

func convertToResponseList(u []member.Member) []member.MemberResponse {
	data := []member.MemberResponse{}
	for _, v := range u {
		data = append(data, convertToResponse(&v))
	}
	return data
}

func convertToCore(u *member.MemberRequest) member.Member {
	return member.Member{
		Username:  u.Username,
		SkinType:  u.SkinType,
		SkinColor: u.SkinColor,
		Gender:    u.Gender,
	}
}
