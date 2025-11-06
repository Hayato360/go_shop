package playerrepository

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Hayato360/go_shop/modules/player"
	"github.com/Hayato360/go_shop/pkg/utils"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type (
	PlayerRepositoryService interface {
		IsUniquePlayer(pctx context.Context, email, username string) bool
		InsertOnePlayer(pctx context.Context, req *player.Player) (bson.ObjectID, error)
		FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfileBson, error)
		InsertOnePlayerTransaction(pctx context.Context, req *player.PlayerTransaction) error
		GetPlayerSavingAccount(pctx context.Context, playerId string) (*player.PlayerSavingAccount, error)
	}

	playerRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(db *mongo.Client) PlayerRepositoryService {
	return &playerRepository{db: db}
}

func (r *playerRepository) playerDbConn(pctx context.Context) *mongo.Database {
	return r.db.Database("player_db")
}

func (r *playerRepository) IsUniquePlayer(pctx context.Context, email, username string) bool {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	// Use a simple struct to check existence without decoding the full player
	var result bson.M
	if err := col.FindOne(
		ctx,
		bson.M{"$or": []bson.M{
			{"username": username},
			{"email": email},
		}},
	).Decode(&result); err != nil {
		log.Printf("Error: IsUniquePlayer: %s", err.Error())
		return true
	}
	return false
}

func (r *playerRepository) InsertOnePlayer(pctx context.Context, req *player.Player) (bson.ObjectID, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	playerId, err := col.InsertOne(ctx, req)
	if err != nil {
		log.Printf("Error: InsertOnePlayer: %s", err.Error())
		return bson.NilObjectID, errors.New("error: failed to insert player")
	}

	return playerId.InsertedID.(bson.ObjectID), nil
}

func (r *playerRepository) FindOnePlayerProfile(pctx context.Context, playerId string) (*player.PlayerProfileBson, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("players")

	result := new(player.PlayerProfileBson)

	if err := col.FindOne(
		ctx,
		bson.M{"_id": utils.ConvertToObjectId(playerId)},
		options.FindOne().SetProjection(
			bson.M{
				"_id":        1,
				"email":      1,
				"username":   1,
				"created_at": 1,
				"updated_at": 1,
			},
		),
	).Decode(result); err != nil {
		log.Printf("error: FindOnePlayerProfile: %s", err.Error())
		return nil, errors.New("error:player profile not found")
	}

	return result, nil
}


func (r *playerRepository) InsertOnePlayerTransaction(pctx context.Context, req *player.PlayerTransaction) error {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("player_transactions")

	result,err := col.InsertOne(ctx,req)
	if err != nil {
		log.Printf("Error: InsertOnePlayerTransaction: %s", err.Error())
		return errors.New("error: insert one player transaction failed")
	}
	log.Printf("Result: InsertOnePlayerTransaction: %v", result.InsertedID)
	
	return nil
}

func (r *playerRepository) GetPlayerSavingAccount(pctx context.Context, playerId string) (*player.PlayerSavingAccount, error) {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	db := r.playerDbConn(ctx)
	col := db.Collection("player_transactions")

	result := new(player.PlayerSavingAccount)

	filter := bson.A{
		bson.D{{Key: "$match", Value: bson.D{{Key: "player_id", Value: playerId}}}},
		bson.D{{Key: "$group", Value: bson.D{
			{Key: "_id", Value: "$player_id"},
			{Key: "balance", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
		}}},
		bson.D{{Key: "$project", Value: bson.D{
			{Key: "player_id", Value: "$_id"},
			{Key: "_id", Value: 0},
			{Key: "balance", Value: 1},
		}}},
	}

	cursors, err := col.Aggregate(ctx, filter)
	if err != nil {
		log.Printf("Error: GetPlayerSavingAccount: %s", err.Error())
		return nil , errors.New("error: failed to aggregate player saving account")
	}

	for cursors.Next(ctx) {
		if err := cursors.Decode(result); err != nil {
			log.Printf("Error: GetPlayerSavingAccount Decode: %s", err.Error())
			_ = cursors.Close(ctx)
			return nil , errors.New("error: failed to decode player saving account")
		}

		// we expect a single aggregation result; close cursor and return
		_ = cursors.Close(ctx)
		return result, nil
	}

	// check for any errors encountered during iteration
	if err := cursors.Err(); err != nil {
		log.Printf("Error: GetPlayerSavingAccount Cursor: %s", err.Error())
		return nil, errors.New("error: cursor encountered an error")
	}

	// no aggregation result found
	return nil, errors.New("error: no player saving account found")
}
