package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"task-manager/internal/api"
	"task-manager/internal/config"
	"task-manager/internal/pkg/db"
	"task-manager/internal/pkg/logger"
	catrepo "task-manager/internal/repository/categories"
	projrepo "task-manager/internal/repository/projects"
	taskrepo "task-manager/internal/repository/tasks"
	catservice "task-manager/internal/service/categories"
	projservice "task-manager/internal/service/projects"
	taskservice "task-manager/internal/service/tasks"
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

	categoryRepo := catrepo.New(pool)
	projectRepo := projrepo.New(pool)

	r, err := api.New(
		myLogger,
		taskservice.New(
			taskrepo.New(pool),
			categoryRepo,
			projectRepo,
		),
		catservice.New(
			categoryRepo,
			projectRepo,
		),
		projservice.New(
			projectRepo,
		),
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

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
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
