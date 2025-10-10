package database

import (
	"context"
	"log"
	"time"

	"github.com/Hayato360/go_shop/config"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"go.mongodb.org/mongo-driver/v2/mongo/readpref"
)

func DbConn(pctx context.Context, cfg *config.Config) *mongo.Client {
	ctx, cancel := context.WithTimeout(pctx, 10*time.Second)
	defer cancel()

	// Create client options
	clientOptions := options.Client().ApplyURI(cfg.Db.Url)

	// Connect to MongoDB (v2 doesn't take context in Connect)
	client, err := mongo.Connect(clientOptions)
	if err != nil {
		log.Fatal("MongoDB connection error:", err)
	}

	// Test the connection with context
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("error test ping database" , err)
	}

	return client
}
