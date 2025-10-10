package playerhandler

import (
	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/modules/player/playerUsecase"
)

type (
	PlayerHttpHandlerService interface{}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewPlayerHttpHandler(cfg *config.Config,playerUsecase playerusecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{playerUsecase: playerUsecase}
}
