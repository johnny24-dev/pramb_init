package interfaces

import "auth_service/pkg/domain"

type UserUseCase interface {
	SendEmailVerifyCode(user domain.User, typeSend int64) error
	VerifyEmail(user domain.User) error
	UpdatePassword(user domain.User) (domain.User, error)
	Register(user domain.User) error
	RegisterValidate(user domain.User) (domain.User, error)
	Login(user domain.User) (domain.User, error)
	ValidateJwtUser(userid uint) (domain.User, error)
	ForgotPassword(user domain.User) error
	ChangePassword(user domain.User) error
	GetUserInfo(user domain.User) (domain.User, error)
}
