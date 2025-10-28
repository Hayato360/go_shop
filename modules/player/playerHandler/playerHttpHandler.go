package playerhandler

import (
	"context"
	"net/http"

	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/modules/player"
	"github.com/Hayato360/go_shop/modules/player/playerUsecase"
	"github.com/Hayato360/go_shop/pkg/request"
	"github.com/Hayato360/go_shop/pkg/response"
	"github.com/labstack/echo/v4"
)

type (
	PlayerHttpHandlerService interface{
		CreatePlayer(c echo.Context) error
	}

	playerHttpHandler struct {
		cfg           *config.Config
		playerUsecase playerusecase.PlayerUsecaseService
	}
)

func NewPlayerHttpHandler(cfg *config.Config,playerUsecase playerusecase.PlayerUsecaseService) PlayerHttpHandlerService {
	return &playerHttpHandler{playerUsecase: playerUsecase}
}


func (h *playerHttpHandler) CreatePlayer(c echo.Context) error {
	ctx := context.Background()

	wrapper := request.ContextWrapper(c)

	req := new(player.CreatePlayerReq)

	if err := wrapper.Bind(req); err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	res , err := h.playerUsecase.CreatePlayer(ctx, req)
	if err != nil {
		return response.ErrResponse(c, http.StatusBadRequest, err.Error())
	}

	
	return response.SuccessResponse(c, http.StatusCreated, res)
}