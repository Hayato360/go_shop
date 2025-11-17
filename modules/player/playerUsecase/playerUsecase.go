package playerusecase

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Hayato360/go_shop/modules/player"
	playerPb "github.com/Hayato360/go_shop/modules/player/playerPb"
	playerrepository "github.com/Hayato360/go_shop/modules/player/playerRepository"
	"github.com/Hayato360/go_shop/pkg/utils"
	"golang.org/x/crypto/bcrypt"
)

type (
	PlayerUsecaseService interface {
		CreatePlayer(pctx context.Context, req *player.CreatePlayerReq) (*player.PlayerProfile, error)
		FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfile, error)
		AddPlayerMoney(pctx context.Context, req *player.CreatePlayerTransactionReq) (*player.PlayerSavingAccount, error) 
		GetPlayerSavingAccount(pctx context.Context, playerId string) (*player.PlayerSavingAccount, error)
		FindOnePlayerCredential(pctx context.Context, email, password string) (*playerPb.PlayerProfile, error)
	}


	playerUsecase struct {
		playerRepository playerrepository.PlayerRepositoryService
	}
)

func NewPlayerUsecase(playerRepository playerrepository.PlayerRepositoryService) PlayerUsecaseService {
	return &playerUsecase{playerRepository: playerRepository}
}

func (u *playerUsecase) CreatePlayer(pctx context.Context, req *player.CreatePlayerReq) (*player.PlayerProfile, error) {
	if !u.playerRepository.IsUniquePlayer(pctx, req.Email, req.Username) {
		return nil, errors.New("error: email or username already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("error: failed to hash password")
	}

	playerId, err := u.playerRepository.InsertOnePlayer(pctx, &player.Player{
		Email:     req.Email,
		Password:  string(hashedPassword),
		Username:  req.Username,
		CreatedAt: utils.LocalTime(),
		UpdatedAt: utils.LocalTime(),
		PlayerRole: []player.PlayerRole{
			{
				RoleTitle: "Player",
				RoleCode:  0,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return u.FindOnePlayerProfile(pctx, playerId.Hex())
}

func (u *playerUsecase) FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfile, error) {
	result, err := u.playerRepository.FindOnePlayerProfile(pctx, playerId)
	if err != nil {
		return nil, err
	}
	return &player.PlayerProfile{
		Id:       result.Id.Hex(),
		Email:    result.Email,
		Username: result.Username,
		CreateAt: result.CreatedAt,
		UpdateAt: result.UpdatedAt,
	}, nil
}


func (u *playerUsecase) AddPlayerMoney(pctx context.Context, req *player.CreatePlayerTransactionReq) (*player.PlayerSavingAccount, error) {
	if err := u.playerRepository.InsertOnePlayerTransaction(pctx, &player.PlayerTransaction{
		PlayerId: req.PlayerId,
		Amount: int64(req.Amount),
		CreatedAt: utils.LocalTime(),
	}); err != nil {
		return nil, err
	}

	return u.playerRepository.GetPlayerSavingAccount(pctx, req.PlayerId)
}

func (u *playerUsecase) GetPlayerSavingAccount(pctx context.Context, playerId string) (*player.PlayerSavingAccount, error){
	return u.playerRepository.GetPlayerSavingAccount(pctx, playerId)
}

func (u *playerUsecase) FindOnePlayerCredential(pctx context.Context, email, password string) (*playerPb.PlayerProfile, error) {
	result, err := u.playerRepository.FindOnePlayerCredential(pctx, email) 
	if err != nil {
		return nil , err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(result.Password)); err != nil {
		log.Printf("Error: FindOnePlayerCredential: %s", err.Error())
		return nil , errors.New("error: invalid credentials")
	}

	loc, _ := time.LoadLocation("Asia/Bankok")

	return &playerPb.PlayerProfile{
		Id:       result.Id.Hex(),
		Email:    result.Email,
		Username: result.Username,
		CreateAt: result.CreatedAt.In(loc).String(),
		UpdateAt: result.UpdatedAt.In(loc).String(),
	}, nil
}