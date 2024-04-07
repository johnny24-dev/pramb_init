package routes

import (
	"api_gateway/pkg/auth/pb"
	"api_gateway/pkg/utils"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(ctx *gin.Context, p pb.AuthServiceClient) {

	uid, err := strconv.Atoi(ctx.GetString("userId"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "View Profile Failed",
			"err":     err,
		})
	}

	res, err := p.GetUserInfo(context.Background(), &pb.UserInfoRequest{
		Uid: int64(uid),
	})
	errs, _ := utils.ExtractError(err)
	if err != nil {
		ctx.JSON(http.StatusConflict, gin.H{
			"Success": false,
			"Message": "Fetching the User Failed",
			"Error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Successfully",
			"data":    res,
		})
	}
}
