package authusecase

import "github.com/Hayato360/go_shop/modules/auth/authRepository"

type (
	AuthUsecaseService interface{}

	authUsecase struct {
		authRepository authrepository.AuthRepositoryService
	}
)

func NewAuthUsecase(authRepository authrepository.AuthRepositoryService) AuthUsecaseService {
	return &authUsecase{authRepository: authRepository}
}