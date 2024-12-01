package main

import (
	"fmt"
	"github.com/SemenVologdin/go-advanced/config"
	"github.com/SemenVologdin/go-advanced/internal/services"
	"github.com/SemenVologdin/go-advanced/internal/verify"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatalf("Error config.New:%v\n", err)
	}

	srv := services.New(cfg)

	router := http.NewServeMux()
	verify.NewHandler(router, verify.HandlerDeps{
		SenderService:  srv.Sender,
		StorageService: srv.Storage,
		HashService:    srv.Hash,
	})

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("server.ListenAndServe: %v\n", err)
	}
}
