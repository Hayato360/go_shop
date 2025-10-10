package server

import (
	"github.com/Hayato360/go_shop/modules/player/playerHandler"
	"github.com/Hayato360/go_shop/modules/player/playerRepository"
	"github.com/Hayato360/go_shop/modules/player/playerUsecase"
)

func (s *server) playerService() {
	repo := playerrepository.NewPlayerRepository(s.db)
	usecase := playerusecase.NewPlayerUsecase(repo)
	httpHandler := playerhandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerhandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerhandler.NewPlayerQueueHandler(s.cfg, usecase)

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// Health check
	player.GET("", s.healthCheckService)
}
