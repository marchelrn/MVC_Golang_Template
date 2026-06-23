package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/your-org/your-app/config"
	"github.com/your-org/your-app/repository"
	"github.com/your-org/your-app/routes"
	"github.com/your-org/your-app/service"
)

func Run(cfg config.Config) error {
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetOutput(os.Stdout)

	cfg = config.Load()

	repo := repository.New()
	svc, err := service.New(repo)
	if err != nil {
		log.Fatal(err)
	}

	r := routes.NewRouter(svc)

	serv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: r,
	}

	if err := serv.ListenAndServe(); err != nil {
		return nil
	}

	return nil
}
