package playerhandler

import playerusecase "github.com/Hayato360/go_shop/modules/player/playerUsecase"

type (
	playerGrpcHandlerService struct {
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler(playerUsecase playerusecase.PlayerUsecaseService) *playerGrpcHandlerService{
	return &playerGrpcHandlerService{playerUsecase : playerUsecase}
}