package main

import (
	"context"
	"log"

	"github.com/elgntt/notes/internal/api"
	"github.com/elgntt/notes/internal/config"
	"github.com/elgntt/notes/internal/pkg/db"
	"github.com/elgntt/notes/internal/pkg/logger"
	"github.com/elgntt/notes/internal/repository/note"
	"github.com/elgntt/notes/internal/service"
)

func main() {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatalln(err)
	}

	ctx := context.Background()

	pool, err := db.OpenDB(ctx, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	logger, err := logger.New(cfg.LogFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	r := api.New(
		service.New(
			note.New(pool),
		),
		logger,
	)

	log.Fatal(r.Run(cfg.ServerPort))

}
