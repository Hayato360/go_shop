package authhandler

import "github.com/Hayato360/go_shop/modules/auth/authUsecase"

type (
	authGrpcHandler struct {
		authUsecase authusecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authusecase.AuthUsecaseService) *authGrpcHandler{
	return &authGrpcHandler{authUsecase: authUsecase}
}