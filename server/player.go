package server

import (
	"log"

	"github.com/Hayato360/go_shop/modules/player/playerHandler"
	"github.com/Hayato360/go_shop/modules/player/playerRepository"
	"github.com/Hayato360/go_shop/modules/player/playerUsecase"
	"github.com/Hayato360/go_shop/pkg/grpccon"
	playerPb "github.com/Hayato360/go_shop/modules/player/playerPb"
)

func (s *server) playerService() {
	repo := playerrepository.NewPlayerRepository(s.db)
	usecase := playerusecase.NewPlayerUsecase(repo)
	httpHandler := playerhandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerhandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerhandler.NewPlayerQueueHandler(s.cfg, usecase)

	//gRPC
	go func() {
		grpcServer, lis := grpccon.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.PlayerUrl)

		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Player gRPC server listening on: %s", s.cfg.Grpc.PlayerUrl)
		grpcServer.Serve(lis)
	}()

	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// Health check
	player.GET("", s.healthCheckService)

	player.POST("/player/register", httpHandler.CreatePlayer)
	player.GET("/player/:player_id", httpHandler.FindOnePlayerProfile)
}
