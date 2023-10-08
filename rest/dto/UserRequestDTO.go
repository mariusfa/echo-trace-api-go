package dto

import (
	"echo/biz/domain"
)

type UserRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (ur *UserRequestDTO) ToDomain() domain.UserRequest {
	return domain.UserRequest{
		Username: ur.Username,
		Password: ur.Password,
	}
}
