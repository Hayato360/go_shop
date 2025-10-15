package server

import (
	"log"

	authhandler "github.com/Hayato360/go_shop/modules/auth/authHandler"
	authPb "github.com/Hayato360/go_shop/modules/auth/authPb"
	authrepository "github.com/Hayato360/go_shop/modules/auth/authRepository"
	authusecase "github.com/Hayato360/go_shop/modules/auth/authUsecase"
	"github.com/Hayato360/go_shop/pkg/grpccon"
)

func (s *server) authService() {
	repo := authrepository.NewAuthRepository(s.db)
	usecase := authusecase.NewAuthUsecase(repo)
	httpHandler := authhandler.NewAuthHttpHandler(s.cfg , usecase)
	grpcHandler := authhandler.NewAuthGrpcHandler(usecase)

	//gRPC 
	go func() {
		grpcServer , lis := grpccon.NewGrpcServer(&s.cfg.Jwt , s.cfg.Grpc.AuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Auth gRPC server listening on: %s", s.cfg.Grpc.AuthUrl)
		grpcServer.Serve(lis)
	}()

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// Health check
	auth.GET("", s.healthCheckService)
}