package auth

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"api_gateway/pkg/auth/pb"
	"api_gateway/pkg/utils"

	"github.com/gin-gonic/gin"
)

type AuthMiddlewareConfig struct {
	svc *ServiceAuth
}

func InitAuthMiddleware(Svc *ServiceAuth) AuthMiddlewareConfig {
	return AuthMiddlewareConfig{
		svc: Svc,
	}
}

func (c *AuthMiddlewareConfig) AuthRequired(ctx *gin.Context) {
	authorization := ctx.Request.Header.Get("Authorization")

	if authorization == "" {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	token := strings.Split(authorization, "Bearer ")
	if len(token) < 2 {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	res, err := c.svc.client.Validate(context.Background(), &pb.ValidateRequest{
		Accesstoken: token[1],
	})

	errs, _ := utils.ExtractError(err)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Error": errs,
		})
	}

	str := strconv.FormatInt(res.Userid, 10)
	ctx.Set("userId", str)
	ctx.Next()
}
