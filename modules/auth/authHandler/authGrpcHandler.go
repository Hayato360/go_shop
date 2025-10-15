package authhandler

import (
	"context"
	authPb "github.com/Hayato360/go_shop/modules/auth/authPb"
	"github.com/Hayato360/go_shop/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authPb.UnimplementedAuthGrpcServiceServer
		authUsecase authusecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(authUsecase authusecase.AuthUsecaseService) *authGrpcHandler{
	return &authGrpcHandler{
		authUsecase: authUsecase,
	}
}

func (g *authGrpcHandler) CredentialSearch (ctx *context.Context,req *authPb.CredentialSearchReq) (*authPb.CredentialSearchRes, error) {
	return nil , nil
}

func  (g *authGrpcHandler) RolesCount (ctx context.Context, req *authPb.RolesCountReq) (*authPb.RolesCountRes, error) {
	return nil, nil
}