package factory

import (
	"database/sql"

	"github.com/vinicius-benevides/go-rest-api/cmd/api/config"
	infraDB "github.com/vinicius-benevides/go-rest-api/src/infrastructure/db"
	"github.com/vinicius-benevides/go-rest-api/src/infrastructure/repositories"
	"github.com/vinicius-benevides/go-rest-api/src/services"
)

type Container struct {
	Config              *config.Config
	DB                  *sql.DB
	EventService        services.EventService
	RegistrationService services.RegistrationService
	UserService         services.UserService
	TokenService        services.TokenService
}

func NewContainer(cfg *config.Config) (*Container, error) {
	dbConn, err := infraDB.NewConnection(infraDB.Config{
		Driver:       cfg.Database.Driver,
		DSN:          cfg.Database.DSN,
		MaxOpenConns: cfg.Database.MaxOpenConn,
		MaxIdleConns: cfg.Database.MaxIdleConn,
	})
	if err != nil {
		return nil, err
	}

	eventRepository := repositories.NewEventRepository(dbConn)
	registrationRepository := repositories.NewRegistrationRepository(dbConn)
	userRepository := repositories.NewUserRepository(dbConn)

	tokenService := services.NewTokenService(cfg.JWT.Secret, cfg.JWT.Expiration)
	eventService := services.NewEventService(eventRepository)
	registrationService := services.NewRegistrationService(eventRepository, registrationRepository)
	userService := services.NewUserService(userRepository, tokenService)

	return &Container{
		Config:              cfg,
		DB:                  dbConn,
		EventService:        eventService,
		RegistrationService: registrationService,
		UserService:         userService,
		TokenService:        tokenService,
	}, nil
}

func (c *Container) Close() error {
	if c.DB != nil {
		return c.DB.Close()
	}

	return nil
}
