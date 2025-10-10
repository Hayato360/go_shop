package middlewareusecase

import "github.com/Hayato360/go_shop/modules/middleware/middlewareRepository"

type (
	MiddlewareUsecaseService interface{}

	middlewareUsecase struct {
		middlewareRepository middlewarerepository.MiddlewareRepositoryService
	}
)

func NewMiddlewareUsecase(middlewareRepository middlewarerepository.MiddlewareRepositoryService) MiddlewareUsecaseService {
	return &middlewareUsecase{middlewareRepository: middlewareRepository}
}