package server

import (
	"log"

	"github.com/benbjohnson/clock"
	"github.com/judoassistant/judoassistant-meta-service/config"
	"github.com/judoassistant/judoassistant-meta-service/db"
	"github.com/judoassistant/judoassistant-meta-service/handler"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/repository"
	"github.com/judoassistant/judoassistant-meta-service/service"
	"go.uber.org/zap"
)

func Init() {
	database, err := db.Init()
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := db.Migrate(database); err != nil {
		log.Fatalln(err.Error())
	}

	clock := clock.New()
	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	config, err := config.NewConfig()
	if err != nil {
		logger.Fatal("Unable to read config", zap.Error(err))
	}

	tournamentRepository := repository.NewTournamentRepository(database)
	tournamentService := service.NewTournamentService(tournamentRepository, clock)
	tournamentHandler := handler.NewTournamentHandler(tournamentService, logger)

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, logger)

	if err := InitScaffoldingData(userService, tournamentService, config, clock); err != nil {
		logger.Fatal("Unable to scaffold database", zap.Error(err))
	}

	authMiddleware := middleware.BasicAuthMiddleware(userService, logger)
	adminAreaMiddleware := middleware.AdminAreaMiddleware(logger)
	router := NewRouter(authMiddleware, adminAreaMiddleware, tournamentHandler, userHandler)

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run(":8080")
}
