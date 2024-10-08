package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/rjrocha19/semana-tech-go-react-server/internal/api"
	"github.com/rjrocha19/semana-tech-go-react-server/internal/store/pgstore"
)

func main() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, fmt.Sprintf(
		"user=%s password=%s host=%s dbname=%s",
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"),
		os.Getenv("PGDATABASE"),
	))

	if err != nil {
		panic(err)
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err)	
	}

	handler := api.NewHandler(pgstore.New(pool))

	go func ()  {
		if err := http.ListenAndServe(":8080", handler); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	
}
