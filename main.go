package main

import (
	"context"
	"log"
	"os"

	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/pkg/database"
	"github.com/Hayato360/go_shop/server"
)

func main() {
	ctx := context.Background()

	// Initialize Config
	cfg := config.LoadConfig(func() string{
		if len(os.Args) < 2 {
			log.Fatal("Please provide the path to the .env file as a command line argument")
		}
		return os.Args[1]
	}())

	db := database.DbConn(ctx , &cfg)
	defer db.Disconnect(ctx)

	server.Start(ctx, &cfg, db)

}