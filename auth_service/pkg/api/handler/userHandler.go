package handler

import (
	"context"
	"errors"
	"net/http"

	"auth_service/pkg/domain"
	"auth_service/pkg/pb"
	interfaces "auth_service/pkg/usecase/interfaces"
)

type UserHandler struct {
	UseCase    interfaces.UserUseCase
	jwtUseCase interfaces.JwtUseCase
	pb.AuthServiceServer
}

func NewUserHandler(useCase interfaces.UserUseCase, jwtUseCase interfaces.JwtUseCase) *UserHandler {
	return &UserHandler{
		UseCase:    useCase,
		jwtUseCase: jwtUseCase,
	}
}

func (h *UserHandler) SendEmailVerifyCode(ctx context.Context, req *pb.SendEmailVerifyRequest) (*pb.SendEmailVerifyResponse, error) {
	user := domain.User{
		Email: req.GetEmail(),
	}

	err := h.UseCase.SendEmailVerifyCode(user, req.GetType())
	if err != nil {
		return &pb.SendEmailVerifyResponse{
			Status: http.StatusBadRequest,
			Error:  "Error",
		}, err
	}
	return &pb.SendEmailVerifyResponse{
		Status: http.StatusOK,
		Error:  "nil",
	}, nil
}

func (h *UserHandler) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {
	user := domain.User{
		Otp: req.GetOtp(),
	}
	err := h.UseCase.VerifyEmail(user)
	if err != nil {
		return &pb.VerifyEmailResponse{
			Status: http.StatusBadRequest,
			Error:  "Error",
		}, err
	}
	return &pb.VerifyEmailResponse{
		Status: http.StatusOK,
		Error:  "nil",
	}, nil
}

func (h *UserHandler) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordReponse, error) {
	user := domain.User{
		Otp:      req.GetOtp(),
		Password: req.GetPassword(),
	}
	userDetails, err := h.UseCase.UpdatePassword(user)
	if err != nil {
		return &pb.UpdatePasswordReponse{
			Status: http.StatusBadRequest,
			Error:  "Error",
		}, err
	}
	accessToken, err := h.jwtUseCase.GenerateAccessToken(int(userDetails.Id), userDetails.Email, "user")
	if err != nil {
		return &pb.UpdatePasswordReponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error in Generating JWT token",
		}, err
	}

	refreshToken, err := h.jwtUseCase.GenerateRefreshToken(int(userDetails.Id), userDetails.Email, "user")
	if err != nil {
		return &pb.UpdatePasswordReponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error in Generating JWT token",
		}, err
	}

	return &pb.UpdatePasswordReponse{
		Status:       http.StatusOK,
		Error:        "nil",
		Accesstoken:  accessToken,
		RefreshToken: refreshToken,
	}, nil

}

func (h *UserHandler) GetUserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.UserInfoResponse, error) {
	user := domain.User{
		Id: uint(req.GetUid()),
	}
	userData, err := h.UseCase.GetUserInfo(user)
	if err != nil {
		return &pb.UserInfoResponse{
			Status:      http.StatusNotFound,
			Error:       "Error",
			Id:          int64(userData.Id),
			UserName:    *userData.Username,
			Pid:         int64(userData.Pid),
			Phone:       userData.Phone,
			Email:       userData.Email,
			VipLevel:    int64(userData.VipLevel),
			IsVerified:  userData.IsVerified,
			IsAdmin:     userData.IsAdmin,
			CanTrade:    userData.CanTrade,
			CanWithdraw: userData.CanWithdraw,
			CanDeposit:  userData.CanDeposit,
		}, err
	}
	if userData.Username == nil {
		userData.Username = &userData.Email
	}
	return &pb.UserInfoResponse{
		Status:      http.StatusOK,
		Error:       "nil",
		Id:          int64(userData.Id),
		UserName:    *userData.Username,
		Pid:         int64(userData.Pid),
		Phone:       userData.Phone,
		Email:       userData.Email,
		VipLevel:    int64(userData.VipLevel),
		IsVerified:  userData.IsVerified,
		IsAdmin:     userData.IsAdmin,
		CanTrade:    userData.CanTrade,
		CanWithdraw: userData.CanWithdraw,
		CanDeposit:  userData.CanDeposit,
	}, nil
}

func (h *UserHandler) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	user := domain.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
		Phone:    req.GetPhone(),
	}
	err := h.UseCase.Register(user)

	if err != nil {
		return &pb.RegisterResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error",
		}, err
	}
	return &pb.RegisterResponse{
		Status: http.StatusOK,
		Error:  "nil",
	}, nil
}

func (h *UserHandler) RegisterValidate(ctx context.Context, req *pb.RegisterValidateRequest) (*pb.RegisterValidateResponse, error) {
	user := domain.User{
		Otp: req.GetOtp(),
	}
	user, err := h.UseCase.RegisterValidate(user)
	if err != nil {
		return &pb.RegisterValidateResponse{
			Status: http.StatusNotFound,
			Error:  "Error",
			Id:     int64(user.Id),
		}, err
	}
	return &pb.RegisterValidateResponse{
		Status: http.StatusOK,
		Error:  "nil",
		Id:     int64(user.Id),
	}, nil
}

func (h *UserHandler) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	user := domain.User{
		Email:    req.GetEmail(),
		Password: req.GetPassword(),
	}

	userDetails, err := h.UseCase.Login(user)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "Error in logging the user",
		}, err
	}
	accessToken, err := h.jwtUseCase.GenerateAccessToken(int(userDetails.Id), userDetails.Email, "user")
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error in Generating JWT token",
		}, err
	}

	refreshToken, err := h.jwtUseCase.GenerateRefreshToken(int(userDetails.Id), userDetails.Email, "user")
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnprocessableEntity,
			Error:  "Error in Generating JWT token",
		}, err
	}

	return &pb.LoginResponse{
		Status:       http.StatusOK,
		Accesstoken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *UserHandler) ForgotPassword(ctx context.Context, req *pb.ForgotPasswordRequest) (*pb.ForgotPasswordResponse, error) {
	user := domain.User{
		Email: req.GetEmail(),
	}
	err := u.UseCase.ForgotPassword(user)
	if err != nil {
		return &pb.ForgotPasswordResponse{
			Status: http.StatusNotFound,
			Error:  "Error in Forget Passsword",
		}, err
	}
	return &pb.ForgotPasswordResponse{
		Status: http.StatusOK,
	}, nil

}
func (u *UserHandler) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.ForgotPasswordResponse, error) {
	user := domain.User{
		Id:       uint(req.Id),
		Password: req.GetPassword(),
	}
	err := u.UseCase.ChangePassword(user)
	if err != nil {
		return &pb.ForgotPasswordResponse{
			Status: http.StatusNotFound,
			Error:  "Error in changing the password",
		}, err
	}
	return &pb.ForgotPasswordResponse{
		Status: http.StatusOK,
	}, nil
}

// Jwt Validation
func (u *UserHandler) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {

	userData := domain.User{}
	ok, claims := u.jwtUseCase.VerifyToken(req.Accesstoken)
	if !ok {
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Error:  "Token Verification Failed",
		}, errors.New("Token failed")
	}
	userData, err := u.UseCase.ValidateJwtUser(claims.Userid)
	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusUnauthorized,
			Userid: int64(userData.Id),
			Error:  "User not found with essesntial token credential",
			Source: claims.Source,
		}, err
	}
	return &pb.ValidateResponse{
		Status: http.StatusOK,
		Userid: int64(userData.Id),
		Source: claims.Source,
	}, nil

}
