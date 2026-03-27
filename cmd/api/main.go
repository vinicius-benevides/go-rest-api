package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/vinicius-benevides/go-rest-api/cmd/api/config"
	"github.com/vinicius-benevides/go-rest-api/cmd/api/factory"
	"github.com/vinicius-benevides/go-rest-api/src/interface/http/middlewares"
	"github.com/vinicius-benevides/go-rest-api/src/interface/http/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	container, err := factory.NewContainer(cfg)
	if err != nil {
		log.Fatalf("could not initialize application: %v", err)
	}
	defer container.Close()

	server := gin.Default()

	authMiddleware := middlewares.Authenticate(container.TokenService)
	handler := routes.NewHandler(container.EventService, container.RegistrationService, container.UserService)
	handler.Register(server, authMiddleware)

	addr := fmt.Sprintf(":%s", cfg.Server.Port)
	if err := server.Run(addr); err != nil {
		log.Fatalf("server stopped: %v", err)
	}
}
