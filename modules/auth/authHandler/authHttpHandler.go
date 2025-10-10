package authhandler

import (
	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authusecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(cfg *config.Config, authUsecase authusecase.AuthUsecaseService) AuthHttpHandlerService {
	return &authHttpHandler{cfg, authUsecase}
}
