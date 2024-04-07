package routes

import (
	"api_gateway/pkg/auth/pb"
	"api_gateway/pkg/domain"
	"api_gateway/pkg/utils"
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func UpdatePassword(ctx *gin.Context, c pb.AuthServiceClient) {
	updatePasswordRequest := domain.UpdatePasswordRequest{}
	err := ctx.Bind(&updatePasswordRequest)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	log.Println(updatePasswordRequest)

	// check compare password and confirm password

	isSame := strings.Compare(updatePasswordRequest.Password, updatePasswordRequest.ConfirmPassword)

	if isSame != 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Confirm Password doesn't match with Password",
			"err":     "Error",
		})
		return
	}

	res, err := c.UpdatePassword(context.Background(), &pb.UpdatePasswordRequest{
		Otp:      updatePasswordRequest.Otp,
		Password: updatePasswordRequest.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Error",
			"err":     err,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Password has been updated",
			"data":    res,
		})
	}
}
