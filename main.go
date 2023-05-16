package main

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
	"github.com/maxwelbm/transinterdigital/internal/config"
	"github.com/maxwelbm/transinterdigital/internal/domain/usecases"
	"github.com/maxwelbm/transinterdigital/internal/gateway/database"
	"github.com/maxwelbm/transinterdigital/internal/gateway/handlers"
	"github.com/maxwelbm/transinterdigital/internal/gateway/handlers/middleware"
	"github.com/maxwelbm/transinterdigital/pkg/logger"
	"net/http"
	"os"
)

func main() {
	log := logger.New()
	defer log.Sync()

	cfg := config.LoadConfig()

	r := chi.NewRouter()
	r.Use(middleware.Logger(log))

	urlDB := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	conn, err := pgx.Connect(context.Background(), urlDB)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed connect database")
		os.Exit(1)
	}
	logger.Info("starting database")
	defer conn.Close(context.Background())

	accountRepo := database.NewAccountRepository(conn)
	transferRepo := database.NewTransfersRepository(conn)
	var repo usecases.Repository
	repo.SetAccount(accountRepo)
	repo.SetTransfer(transferRepo)

	useCase := usecases.New(repo)
	handler := handlers.NewHandlers(useCase, cfg)
	handler.Handlers(r)

	logger.Info("starting server")
	http.ListenAndServe(":3000", r)
}
