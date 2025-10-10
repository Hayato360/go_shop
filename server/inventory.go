package server

import (
	"github.com/Hayato360/go_shop/modules/inventory/inventoryHandler"
	"github.com/Hayato360/go_shop/modules/inventory/inventoryRepository"
	"github.com/Hayato360/go_shop/modules/inventory/inventoryUsecase"
)

func (s *server) inventoryService() {
	repo := inventoryrepository.NewInventoryRepository(s.db)
	usecase := inventoryusecase.NewInventoryUsecase(repo)
	httpHandler := inventoryhandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryhandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryhandler.NewInventoryQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	// Health check
	inventory.GET("", s.healthCheckService)
}
