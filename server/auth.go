package server

import (
	authhandler "github.com/Hayato360/go_shop/modules/auth/authHandler"
	authrepository "github.com/Hayato360/go_shop/modules/auth/authRepository"
	authusecase "github.com/Hayato360/go_shop/modules/auth/authUsecase"
)

func (s *server) authService() {
	repo := authrepository.NewAuthRepository(s.db)
	usecase := authusecase.NewAuthUsecase(repo)
	httpHandler := authhandler.NewAuthHttpHandler(s.cfg , usecase)
	grpcHandler := authhandler.NewAuthGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health check
	auth.GET("", s.healthCheckService)
}