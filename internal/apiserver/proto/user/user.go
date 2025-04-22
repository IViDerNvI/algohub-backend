package user

import v1 "github.com/ividernvi/algohub/model/v1"

func ToUser(u *UserInfo) *v1.User {
	return &v1.User{
		Email:      u.Email,
		UserName:   u.Username,
		Bio:        u.Bio,
		NickName:   u.Nickname,
		Phone:      u.Phone,
		Company:    u.Company,
		Location:   u.Location,
		ProfileURL: u.ProfileUrl,
	}
}

func ToErrorResponse(err error) *ErrorResponse {
	if err == nil {
		return &ErrorResponse{
			Code:    200,
			Message: "OK",
		}
	}
	return &ErrorResponse{
		Code:    500,
		Message: err.Error(),
	}
}

func ToUserInfo(u *v1.User) *UserInfo {
	return &UserInfo{
		Email:      u.Email,
		Username:   u.UserName,
		Bio:        u.Bio,
		Nickname:   u.NickName,
		Phone:      u.Phone,
		Company:    u.Company,
		Location:   u.Location,
		ProfileUrl: u.ProfileURL,
	}
}
