package inventoryhandler

import (
	"github.com/Hayato360/go_shop/modules/inventory/inventoryUsecase"
)

type (
	inventoryGrpcHandler struct {
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler( inventoryUsecase inventoryusecase.InventoryUsecaseService) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{
		inventoryUsecase: inventoryUsecase,
	}
}
