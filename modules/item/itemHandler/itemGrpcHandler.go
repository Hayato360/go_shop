package itemhandler

import (
	"context"

	"github.com/Hayato360/go_shop/modules/item/itemUsecase"
	itemPb "github.com/Hayato360/go_shop/modules/item/itemPb"

	
)

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

func(g *itemGrpcHandler) FindItemsInIds(ctx context.Context, req *itemPb.FindItemsInIdsReq) (*itemPb.FindItemsInIdsRes, error){
	return nil , nil
}
