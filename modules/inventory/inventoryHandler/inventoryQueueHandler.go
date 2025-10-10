package inventoryhandler

import (
	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/modules/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandlerService interface{}

	inventoryQueueHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryusecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandler(cfg *config.Config, inventoryUsecase inventoryusecase.InventoryUsecaseService) InventoryQueueHandlerService {
	return &inventoryQueueHandler{
		cfg:              cfg,
		inventoryUsecase: inventoryUsecase,
	}
}
