package dto

import "github.com/student/ginblog/datamodels"

type UserDto struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func ToUserDto(user datamodels.User) UserDto {
	return UserDto{
		ID:       user.ID,
		Username: user.Username,
	}
}
