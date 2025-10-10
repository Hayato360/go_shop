package inventoryhandler

import (
	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/modules/inventory/inventoryUsecase"
)

type (
	InventoryHttpHandlerService interface{}

	inventoryHttpHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler(cfg *config.Config, inventoryUsecase inventoryusecase.InventoryUsecaseService) InventoryHttpHandlerService {
	return &inventoryHttpHandler{
		cfg:              cfg,
		inventoryUsecase: inventoryUsecase,
	}
}
