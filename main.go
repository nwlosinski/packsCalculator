package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nwlosinski/packsCalculator/calculator"
	"github.com/nwlosinski/packsCalculator/config"
	"github.com/nwlosinski/packsCalculator/handlers"
)

func main() {
	cfg, err := config.Load("config.json")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	repo := calculator.NewMemoryRepo(cfg.DefaultPackSizes)
	service := calculator.NewService(repo)
	h := handlers.NewHandler(service)

	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir("./ui")))

	// API
	h.Register(mux)

	// Start server
	addr := fmt.Sprintf(":%d", cfg.Port)
	log.Println("Server running on", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}
