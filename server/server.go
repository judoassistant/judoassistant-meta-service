package server

import (
	"fmt"

	"github.com/benbjohnson/clock"
	"github.com/judoassistant/judoassistant-meta-service/config"
	"github.com/judoassistant/judoassistant-meta-service/handler"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/repository"
	"github.com/judoassistant/judoassistant-meta-service/service"
	"go.uber.org/zap"
)

func Init() {
	clock := clock.New()
	logger, _ := zap.NewDevelopment() // TODO: Use production
	defer logger.Sync()

	config, err := config.NewConfig()
	if err != nil {
		logger.Fatal("Unable to read config", zap.Error(err))
		return
	}

	database, err := repository.NewDatabase(config)
	if err != nil {
		logger.Fatal("Unable to initialize database", zap.Error(err))
		return
	}

	tournamentRepository := repository.NewTournamentRepository(database)
	tournamentService := service.NewTournamentService(tournamentRepository, clock)
	tournamentHandler := handler.NewTournamentHandler(tournamentService, logger)

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, logger)

	if err := InitScaffoldingData(userService, tournamentService, config, logger, clock); err != nil {
		logger.Fatal("Unable to scaffold database", zap.Error(err))
		return
	}

	loggingMiddleware := middleware.NewLoggingMiddleware(logger, clock)
	authMiddleware := middleware.NewBasicAuthMiddleware(userService, logger)
	adminAreaMiddleware := middleware.NewAdminAreaMiddleware(logger)
	router, err := NewRouter(config, loggingMiddleware, authMiddleware, adminAreaMiddleware, tournamentHandler, userHandler, logger)
	if err != nil {
		logger.Fatal("Unable to initialize router", zap.Error(err))
		return
	}

	router.SetTrustedProxies(config.URL)
	router.Run(fmt.Sprintf(":%d", config.Port))
}
