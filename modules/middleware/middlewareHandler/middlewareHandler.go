package middlewarehandler

import (
	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface{}

	middlewareHandler struct {
		cfg               *config.Config
		middlewareUsecase middlewareusecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(cfg *config.Config, middlewareUsecase middlewareusecase.MiddlewareUsecaseService) MiddlewareHandlerService {
	return &middlewareHandler{
		cfg:               cfg,
		middlewareUsecase: middlewareUsecase,
	}
}
