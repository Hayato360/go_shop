package main

import (
	"context"
	"log"
	"os"

	"github.com/Hayato360/go_shop/config"
	"github.com/Hayato360/go_shop/pkg/database/migration"
)

func main() {
	ctx := context.Background()
	_ = ctx

	// Initialize config
	cfg := config.LoadConfig(func() string{
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is qeiored")
		}
		return os.Args[1]
	}())

	switch cfg.App.Name {
	case "player" : 
	case "auth" :
		migration.AuthMigrate(ctx, &cfg)
	case "item" :
	case "inventory" : 
	case "payment" :
	
	}
}