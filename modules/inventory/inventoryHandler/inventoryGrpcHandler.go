package inventoryhandler

import (
	"context"

	inventoryPb "github.com/Hayato360/go_shop/modules/inventory/inventoryPb"
	"github.com/Hayato360/go_shop/modules/inventory/inventoryUsecase"
)

type (
	inventoryGrpcHandler struct {
		inventoryUsecase inventoryusecase.InventoryUsecaseService
		inventoryPb.UnimplementedInventoryGrpcServiceServer
	}
)

func NewInventoryGrpcHandler( inventoryUsecase inventoryusecase.InventoryUsecaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{
		inventoryUsecase: inventoryUsecase,
	}
}

func (g *inventoryGrpcHandler) IsAvailableToSell(ctx context.Context, req *inventoryPb.IsAvailableToSellReq) (*inventoryPb.IsAvailableToSellRes, error){
	return nil , nil
}
