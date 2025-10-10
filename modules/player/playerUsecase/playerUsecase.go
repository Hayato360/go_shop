package playerusecase

import "github.com/Hayato360/go_shop/modules/player/playerRepository"

type (
	PlayerUsecaseService interface{}

	playerUsecase struct {
		playerRepository playerrepository.PlayerRepositoryService
	}
)

func NewPlayerUsecase(playerRepository playerrepository.PlayerRepositoryService) PlayerUsecaseService {
	return &playerUsecase{playerRepository: playerRepository}
}

