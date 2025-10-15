package server

import (
	"log"

	"github.com/Hayato360/go_shop/modules/inventory/inventoryHandler"
	inventoryPb "github.com/Hayato360/go_shop/modules/inventory/inventoryPb"
	"github.com/Hayato360/go_shop/modules/inventory/inventoryRepository"
	"github.com/Hayato360/go_shop/modules/inventory/inventoryUsecase"
	"github.com/Hayato360/go_shop/pkg/grpccon"
)

func (s *server) inventoryService() {
	repo := inventoryrepository.NewInventoryRepository(s.db)
	usecase := inventoryusecase.NewInventoryUsecase(repo)
	httpHandler := inventoryhandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryhandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryhandler.NewInventoryQueueHandler(s.cfg, usecase)

		//gRPC 
	go func() {
		grpcServer , lis := grpccon.NewGrpcServer(&s.cfg.Jwt , s.cfg.Grpc.InventoryUrl)

		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Inventory gRPC server listening on: %s", s.cfg.Grpc.InventoryUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	// Health check
	inventory.GET("", s.healthCheckService)
}
