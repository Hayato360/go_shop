package itemhandler

import (
	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/modules/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface{}

	itemHttpHandler struct {
		cfg         *config.Config
		itemUsecase itemusecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(cfg *config.Config, itemUsecase itemusecase.ItemUsecaseService) ItemHttpHandlerService {
	return &itemHttpHandler{
		cfg:         cfg,
		itemUsecase: itemUsecase,
	}
}
