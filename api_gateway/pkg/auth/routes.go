package auth

import (
	"api_gateway/pkg/auth/routes"
	"api_gateway/pkg/config"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) *ServiceAuth {
	svc := &ServiceAuth{
		client: InitServiceClient(cfg),
	}

	authorize := InitAuthMiddleware(svc)

	r.POST("/api/v1/login", svc.Login)
	r.POST("/api/v1/sendEmailVerifyCode", svc.SendEmailVerifyCode)
	r.POST("/api/v1/verifyEmail", svc.VerifyEmail)
	r.POST("/api/v1/updatePassword", svc.UpdatePassword)

	user := r.Group("/api/v1/user")
	{
		// user.POST("/register", svc.Register)
		// user.POST("/register/validate", svc.RegitserValidate)
		// user.POST("/login", svc.Login)
		// user.POST("/forget/password", svc.ForgotPassword)
		// user.POST("/forget/password/validate", svc.RegitserValidate)
		// user.PATCH("/forget/password/validate/newpassword", svc.ChangePassword)
		user.GET("/view", authorize.AuthRequired, svc.GetUserInfo)
	}

	return svc

}

func (svc *ServiceAuth) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.client)
}

func (svc *ServiceAuth) RegitserValidate(ctx *gin.Context) {
	routes.RegisterValidate(ctx, svc.client)
}

func (svc *ServiceAuth) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.client)
}

func (svc *ServiceAuth) ForgotPassword(ctx *gin.Context) {
	routes.ForgotPassword(ctx, svc.client)
}

func (svc *ServiceAuth) ChangePassword(ctx *gin.Context) {
	routes.ChangePassword(ctx, svc.client)
}

// ShowAccount godoc
// @Summary      Show an account
// @Tags         account
// @Accept       json
// @Produce      json
// @Header       200,400,default  {string}  Token     "token"
func (svc *ServiceAuth) GetUserInfo(ctx *gin.Context) {
	routes.GetUserInfo(ctx, svc.client)
}

func (svc *ServiceAuth) SendEmailVerifyCode(ctx *gin.Context) {
	routes.SendEmailVerifyCode(ctx, svc.client)
}

func (svc *ServiceAuth) VerifyEmail(ctx *gin.Context) {
	routes.VerifyEmail(ctx, svc.client)
}

func (svc *ServiceAuth) UpdatePassword(ctx *gin.Context) {
	routes.UpdatePassword(ctx, svc.client)
}
