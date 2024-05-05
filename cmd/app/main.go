package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/elgntt/notes/internal/api/handlers"
	"github.com/elgntt/notes/internal/config"
	"github.com/elgntt/notes/internal/pkg/db"
	"github.com/elgntt/notes/internal/pkg/logger"
	"github.com/elgntt/notes/internal/repository"
	taskservice "github.com/elgntt/notes/internal/service"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalln(err)
	}

	pool, err := db.OpenDB(context.Background(), cfg.DBConfig)
	if err != nil {
		log.Fatalln(err)
	}

	myLogger, err := logger.New(cfg.ServerConfig.LogFilePath)
	if err != nil {
		log.Fatalln(err)
	}

	r, err := handlers.New(
		taskservice.New(
			repository.New(pool),
		),
		myLogger,
	)
	if err != nil {
		log.Fatalln(err)
	}

	srv := &http.Server{
		Addr:    ":" + cfg.ServerConfig.ServerPort,
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	waitForShutdown(srv)
}

func waitForShutdown(srv *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
	case <-ctx.Done():
		log.Println("Done")
	}
	log.Println("Server exiting")
}
