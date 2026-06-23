package app

import (
	"log"

	"github.com/your-org/your-app/config"
	"github.com/your-org/your-app/internal/server"
)

func Run() {
	cfg := config.Load()
	log.Println("Server running in port :", cfg.Port)
	if err := server.Run(cfg); err != nil {
		log.Fatal(err)
	}
}
