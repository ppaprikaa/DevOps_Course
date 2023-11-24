package main

import (
	"app/internal/cfg"
	"app/internal/db"
	"app/internal/handlers"
	"app/internal/server"
	"app/internal/storage"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := cfg.Load()

	dbCtx, dbCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer dbCancel()
	db, err := db.Connect(dbCtx, cfg.DB.DSN)
	must(err)
	defer db.Close(context.Background())

	userStorage := storage.NewUser(db)
	userHandler := handlers.NewUserServiceHandler(userStorage)

	router := chi.NewMux()

	router.Get("/users", userHandler.GetAll())
	router.Delete("/users/{id}", userHandler.Delete())
	router.Post("/users/{username}", userHandler.Insert())

	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	srv := server.New(addr, router)

	log.Printf("started server at %s\n", addr)
	must(srv.Run())
	fmt.Println(db)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
