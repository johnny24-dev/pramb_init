package di

import (
	"auth_service/pkg/api"
	"auth_service/pkg/api/handler"
	"auth_service/pkg/config"
	"auth_service/pkg/db"
	"auth_service/pkg/repository"
	"auth_service/pkg/usecase"

	"github.com/google/wire"
)

func InitApi_E(cfg config.Config) (*api.ServerHttp, error) {
	wire.Build(
		db.ConnectToDb,
		repository.NewUserRepo,
		usecase.NewUserUseCase,
		usecase.NewJWTUseCase,
		handler.NewUserHandler,
		api.NewServerHttp,
	)
	return &api.ServerHttp{}, nil
}
