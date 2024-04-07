package routes

import (
	"api_gateway/pkg/auth/pb"
	"api_gateway/pkg/domain"
	"api_gateway/pkg/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, p pb.AuthServiceClient) {
	body := domain.User{}
	err := ctx.BindJSON(&body)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := p.Register(context.Background(), &pb.RegisterRequest{
		Email:    body.Email,
		Password: body.Password,
		Phone:    body.Phone,
	})

	errs, _ := utils.ExtractError(err)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"Success": false,
			"Message": "Registering the User Failed",
			"Error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Successfully sent the OTP",
			"data":    res,
		})
	}

}

func RegisterValidate(ctx *gin.Context, p pb.AuthServiceClient) {
	body := domain.User{}
	err := ctx.Bind(&body)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := p.RegisterValidate(context.Background(), &pb.RegisterValidateRequest{
		Otp: body.Otp,
	})

	errs, _ := utils.ExtractError(err)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "OTP Verification Failed",
			"Error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Successfully registered the user",
			"data":    res,
		})
	}

}
