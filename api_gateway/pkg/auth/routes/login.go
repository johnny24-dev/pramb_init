package routes

import (
	"api_gateway/pkg/auth/pb"
	"api_gateway/pkg/utils"
	"context"
	"net/http"

	"api_gateway/pkg/domain"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	user := domain.User{}
	err := ctx.Bind(&user)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Email:    user.Email,
		Password: user.Password,
	})

	errs, _ := utils.ExtractError(err)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"Success": false,
			"Message": "Login credentials failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "User successfully logged in",
			"data":    res,
		})
	}

}

func ForgotPassword(ctx *gin.Context, c pb.AuthServiceClient) {
	user := domain.User{}
	err := ctx.Bind(&user)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.ForgotPassword(context.Background(), &pb.ForgotPasswordRequest{
		Email: user.Email,
	})
	// extracting the error message from the GRPC error
	errs, _ := utils.ExtractError(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Forgot Password failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Otp Successfully Sent",
			"data":    res,
		})
	}
}

func ChangePassword(ctx *gin.Context, c pb.AuthServiceClient) {
	user := domain.User{}
	err := ctx.Bind(&user)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.ChangePassword(context.Background(), &pb.ChangePasswordRequest{
		Id:       int64(user.Id),
		Password: user.Password,
	}) // extracting the error message from the GRPC error
	errs, _ := utils.ExtractError(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Change Password Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Password Successfully Changed",
			"data":    res,
		})
	}
}
