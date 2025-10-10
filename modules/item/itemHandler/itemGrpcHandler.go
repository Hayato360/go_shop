package itemhandler

import "github.com/Hayato360/go_shop/modules/item/itemUsecase"

type (
	itemGrpcHandler struct {
		itemUsecase itemusecase.ItemUsecaseService
	}
)

func NewItemGrpcHandler(itemUsecase itemusecase.ItemUsecaseService) *itemGrpcHandler {
	return &itemGrpcHandler{
		itemUsecase: itemUsecase,
	}
}
