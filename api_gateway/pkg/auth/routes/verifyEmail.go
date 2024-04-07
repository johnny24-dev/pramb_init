package routes

import (
	"api_gateway/pkg/auth/pb"
	"api_gateway/pkg/domain"
	"api_gateway/pkg/utils"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func VerifyEmail(ctx *gin.Context, c pb.AuthServiceClient) {
	user := domain.User{}

	err := ctx.Bind(&user)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.VerifyEmail(context.Background(), &pb.VerifyEmailRequest{
		Otp: user.Otp,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "OTP does not match",
			"err":     err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Verify successfully",
			"data":    res,
		})
	}

}
