package main

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

	app "github.com/tarvarrs/transaction-blacklist-guard/internal/application/walletoperation"
	"github.com/tarvarrs/transaction-blacklist-guard/internal/config"
	domain "github.com/tarvarrs/transaction-blacklist-guard/internal/domain/walletoperation"
	"github.com/tarvarrs/transaction-blacklist-guard/internal/infrastructure/postgres"
	httpwallet "github.com/tarvarrs/transaction-blacklist-guard/internal/interfaces/http/walletoperation"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file provided, using system env")
	}

	cfg := config.Load()

	db, err := sql.Open("postgres", cfg.PGConnString)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	blacklistRepo := postgres.NewBlacklistRepository(db)
	decisionService := domain.NewDecisionService(blacklistRepo)
	appService := app.NewService(decisionService)
	handler := httpwallet.NewHandler(appService)

	router := gin.Default()
	handler.Register(router)

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("shutting down gracefully")
}
