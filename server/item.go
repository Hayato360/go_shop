package server

import (
	"log"

	"github.com/Hayato360/go_shop/modules/item/itemHandler"
	itemPb "github.com/Hayato360/go_shop/modules/item/itemPb"
	"github.com/Hayato360/go_shop/modules/item/itemRepository"
	"github.com/Hayato360/go_shop/modules/item/itemUsecase"
	"github.com/Hayato360/go_shop/pkg/grpccon"
)

func (s *server) itemService() {
	repo := itemrepository.NewItemRepository(s.db)
	usecase := itemusecase.NewItemUsecase(repo)
	httpHandler := itemhandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemhandler.NewItemGrpcHandler(usecase)

	//gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.ItemUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Item gRPC server listening on: %s", s.cfg.Grpc.ItemUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// Health check
	item.GET("", s.healthCheckService)
}
