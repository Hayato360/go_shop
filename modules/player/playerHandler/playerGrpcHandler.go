package playerhandler

import (
	"context"

	playerPb "github.com/Hayato360/go_shop/modules/player/playerPb"
	"github.com/Hayato360/go_shop/modules/player/playerUsecase"
)

type (
	playerGrpcHandler struct {
		playerUsecase playerusecase.PlayerUsecaseService
		playerPb.UnimplementedPlayerGrpcServiceServer
	}
)

func NewPlayerGrpcHandler(playerUsecase playerusecase.PlayerUsecaseService) *playerGrpcHandler{
	return &playerGrpcHandler{
		playerUsecase : playerUsecase,
	}
}

func (g *playerGrpcHandler) CredentialSearch(ctx context.Context, req *playerPb.CredentialSearchReq) (*playerPb.PlayerProfile, error) {

	return g.playerUsecase.FindOnePlayerCredential(ctx, req.Email, req.Password)
}

func (g *playerGrpcHandler) FindOnePlayerProfileToRefresh(ctx context.Context, req *playerPb.FindOnePlayerProfileToRefreshReq)  (*playerPb.PlayerProfile, error){
	return nil , nil
}

func (g *playerGrpcHandler) GetPlayerSavingAccount(ctx context.Context, req *playerPb.GetPlayerSavingAccountReq)  (*playerPb.GetPlayerSavingAccountRes, error){
	return nil , nil
}