package routes

import (
	"api_gateway/pkg/auth/pb"
	"api_gateway/pkg/domain"
	"api_gateway/pkg/utils"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func SendEmailVerifyCode(ctx *gin.Context, c pb.AuthServiceClient) {
	emailRequest := domain.EmailRequest{}
	err := ctx.Bind(&emailRequest)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	typeSend, err := strconv.Atoi(emailRequest.Type)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "TypeSend must be number",
			"err":     err,
		})
	}

	res, err := c.SendEmailVerifyCode(context.Background(), &pb.SendEmailVerifyRequest{
		Email: emailRequest.Email,
		Type:  int64(typeSend),
	})

	errs, _ := utils.ExtractError(err)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Send Otp credentials failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "OTP successfully sent",
			"data":    res,
		})
	}
}
