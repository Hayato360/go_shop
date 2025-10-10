package playerhandler

import (
	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/modules/player/playerUsecase"
)

type (
	PlayerQueueHandlerService interface{}

	playerQueueHandler struct {
		cfg           *config.Config
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(cfg *config.Config, playerUsecase playerusecase.PlayerUsecaseService) PlayerQueueHandlerService {
	return &playerQueueHandler{playerUsecase: playerUsecase}
}