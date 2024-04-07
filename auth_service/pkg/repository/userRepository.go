package repository

import (
	"auth_service/pkg/domain"
	interfaces "auth_service/pkg/repository/interfaces"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepo(db *gorm.DB) interfaces.UserRepo {
	return &userDatabase{
		DB: db,
	}
}

func (r *userDatabase) UpdatePassword(user domain.User) error {
	var res_user domain.User
	r.DB.First(&res_user, user.Id)
	res_user.Password = user.Password
	res_user.IsVerified = user.IsVerified
	res_user.Otp = user.Otp

	result := r.DB.Save(res_user).Error
	return result
}

func (r *userDatabase) CheckExitsByPId(user domain.User) int64 {
	var count int64
	r.DB.Model(&user).Where("pid = ?", user.Pid).Count(&count)
	return count
}

func (r *userDatabase) FindUserByPId(user domain.User) (domain.User, error) {
	result := r.DB.First(&user, "pid LIKE ?", user.Pid).Error
	return user, result
}

func (r *userDatabase) FindByUserEmail(user domain.User) (domain.User, error) {
	result := r.DB.First(&user, "email LIKE ?", user.Email).Error
	return user, result
}

func (r *userDatabase) Create(user domain.User) error {

	result := r.DB.Create(&user).Error
	return result
}
func (r *userDatabase) FindUserByOtp(user domain.User) (domain.User, error) {
	result := r.DB.Where("otp LIKE ?", user.Otp).First(&user)
	return user, result.Error
}
func (r *userDatabase) NullTheOtp(user domain.User) int64 {
	var userData domain.User
	result := r.DB.Model(&userData).Where("id = ?", user.Id).Update("otp", nil)
	return result.RowsAffected
}

func (r *userDatabase) FindUserById(userid uint) (domain.User, error) {
	user := domain.User{}
	result := r.DB.First(&user, "id = ?", userid).Error
	return user, result
}

func (r *userDatabase) IsOtpVerified(username string) string {
	var otp string
	r.DB.Raw("select otp from users where username LIKE ?", username).Scan(&otp)
	return otp
}

func (r *userDatabase) DeleteUser(user domain.User) error {
	result := r.DB.Exec("DELETE FROM users WHERE email LIKE ?", user.Email).Error
	return result
}

func (r *userDatabase) UpdateOtp(user domain.User) error {
	result := r.DB.Model(&user).Where("id = ?", user.Id).Update("otp", user.Otp)
	return result.Error
}

func (r *userDatabase) VerifyUser(user domain.User) (domain.User, error) {
	result := r.DB.Model(&user).Where("id = ?", user.Id).Update("is_verified", true)
	return user, result.Error
}

func (r *userDatabase) ChangePassword(user domain.User) error {
	result := r.DB.Model(&user).Where("id = ?", 1).Update("password", user.Password)
	return result.Error
}
