package services

import "github.com/suttapak/cacf/dto"

type AuthService interface {
	SignIn(dto.SignInDTO) (*dto.SignInReply, error)
}
