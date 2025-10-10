package server

import (
	"github.com/Hayato360/go_shop/modules/item/itemHandler"
	"github.com/Hayato360/go_shop/modules/item/itemRepository"
	"github.com/Hayato360/go_shop/modules/item/itemUsecase"
)

func (s *server) itemService() {
	repo := itemrepository.NewItemRepository(s.db)
	usecase := itemusecase.NewItemUsecase(repo)
	httpHandler := itemhandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemhandler.NewItemGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// Health check
	item.GET("", s.healthCheckService)
}
