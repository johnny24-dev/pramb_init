package usecase

import (
	"errors"

	"auth_service/pkg/domain"
	interfaces "auth_service/pkg/repository/interfaces"
	userCase "auth_service/pkg/usecase/interfaces"
	"auth_service/pkg/utils"
)

type userUseCase struct {
	Repo interfaces.UserRepo
}

func (use *userUseCase) SendEmailVerifyCode(user domain.User, typeSend int64) error {
	userData, err := use.Repo.FindByUserEmail(user)
	if typeSend == 0 {
		if err == nil {
			if !userData.IsVerified {
				err := use.Repo.DeleteUser(user)
				if err != nil {
					return errors.New("Could Not delete unethenticated user")
				}
			} else {
				return errors.New("Email Address already exists")
			}
		}
		// generate pay_id for user
		var pid int64
		for {
			// Generate a random 10-digit number
			pid = utils.RandomPId()
			_user := domain.User{
				Pid: uint(pid),
			}
			// Check if the number exists in the database
			count := use.Repo.CheckExitsByPId(_user)
			// If the number doesn't exist, break the loop
			if count == 0 {
				break
			}
		}
		user.Pid = uint(pid)
		otp := utils.Otpgeneration(user.Email)
		user.Otp = otp
		err = use.Repo.Create(user)
		if err != nil {
			return err
		}
		return nil
	} else if typeSend == 1 {
		if err != nil {
			return errors.New("Email doesn't exist")
		}
		otp := utils.Otpgeneration(user.Email)
		userData.Otp = otp

		err = use.Repo.UpdateOtp(userData)
		if err != nil {
			return errors.New("Could not update the OTP")
		}
		return nil
	}
	return errors.New("Type is uncorrect")
}

func (use *userUseCase) VerifyEmail(user domain.User) error {
	_, err := use.Repo.FindUserByOtp(user)
	if err != nil {
		return errors.New("Enterd wrong OTP")
	}
	return nil
}

func (use *userUseCase) UpdatePassword(user domain.User) (domain.User, error) {
	userData, err := use.Repo.FindUserByOtp(user)
	if err != nil {
		return userData, errors.New("Enterd wrong OTP")
	}

	userData.Password = utils.HashPassword(user.Password)
	userData.IsVerified = true
	userData.Otp = ""
	//update user to DB

	err = use.Repo.UpdatePassword(userData)

	if err != nil {
		return userData, err
	}

	return userData, nil
}

func (use *userUseCase) GetUserInfo(user domain.User) (domain.User, error) {
	userData, err := use.Repo.FindUserById(user.Id)
	if err != nil {
		return userData, errors.New("User id is not found!")
	}
	return userData, nil
}

func (use *userUseCase) Register(user domain.User) error {

	validationError := utils.ValidateUser(user)
	if validationError != nil {
		return validationError
	}

	// generate pay_id for user
	var pid int64
	for {
		// Generate a random 10-digit number
		pid = utils.RandomPId()
		_user := domain.User{
			Pid: uint(pid),
		}
		// Check if the number exists in the database
		count := use.Repo.CheckExitsByPId(_user)
		// If the number doesn't exist, break the loop
		if count == 0 {
			break
		}
	}
	user.Pid = uint(pid)
	userData, err := use.Repo.FindByUserEmail(user)
	if err == nil {
		if !userData.IsVerified {
			err := use.Repo.DeleteUser(user)
			if err != nil {
				return errors.New("Could Not delete unethenticated user")
			}
		} else {
			return errors.New("Email Address already exists")
		}
	}

	otp := utils.Otpgeneration(user.Email)
	user.Otp = otp

	user.Password = utils.HashPassword(user.Password)

	err = use.Repo.Create(user)
	if err != nil {
		return err
	}
	return nil
}

func (use *userUseCase) RegisterValidate(user domain.User) (domain.User, error) {
	user, err := use.Repo.FindUserByOtp(user)
	if err != nil {
		return user, errors.New("Enterd wrong OTP")
	}

	rows := use.Repo.NullTheOtp(user)
	if rows == 0 {
		return user, errors.New("Could not update the OTP")
	}

	user, err = use.Repo.VerifyUser(user)
	if err != nil {
		return user, errors.New("Could not verify the user")
	}
	return user, nil
}

func (use *userUseCase) Login(user domain.User) (domain.User, error) {
	var userDetatils domain.User
	var err error
	if user.Email != "" {
		userDetatils, err = use.Repo.FindByUserEmail(user)
		if err != nil {
			return userDetatils, errors.New("User not found")
		}
	}

	if userDetatils.IsVerified == false {
		err := use.Repo.DeleteUser(userDetatils)
		if err != nil {
			return userDetatils, errors.New("Could not delete unauthenticateduser")
		}
		return userDetatils, errors.New("User not Authenticated, Register again")

	}

	if !utils.VerifyPassword(user.Password, userDetatils.Password) {
		return userDetatils, errors.New("Password is not matched or worg")
	}
	return userDetatils, nil
}

func (use *userUseCase) ForgotPassword(user domain.User) error {
	user, err := use.Repo.FindByUserEmail(user)
	if err != nil {
		return errors.New("Email Address not found!")
	}

	otp := utils.Otpgeneration(user.Email)
	user.Otp = otp

	err = use.Repo.UpdateOtp(user)
	if err != nil {
		return errors.New("Could not update the OTP")
	}
	return nil
}

func (use *userUseCase) ChangePassword(user domain.User) error {
	// user.Password = utils.HashPassword(user.Password)
	err := use.Repo.ChangePassword(user)
	if err != nil {
		return errors.New("Could not change the password")
	}
	return nil
}

func (use *userUseCase) ValidateJwtUser(userId uint) (domain.User, error) {
	user, err := use.Repo.FindUserById(userId)
	if err != nil {
		return user, errors.New("Unauthorized User")
	}
	return user, nil
}

func NewUserUseCase(repo interfaces.UserRepo) userCase.UserUseCase {
	return &userUseCase{
		Repo: repo,
	}
}
