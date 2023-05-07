package server

import (
	"log"

	"github.com/benbjohnson/clock"
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
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	tournamentRepository := repository.NewTournamentRepository(database)
	tournamentService := service.NewTournamentService(tournamentRepository, clock)
	tournamentHandler := handler.NewTournamentHandler(tournamentService, logger)

	userRepository := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService, logger)

	if err := InitScaffoldingData(userService, tournamentService, clock); err != nil {
		log.Fatalln(err)
	}

	authMiddleware := middleware.BasicAuthMiddleware(userService)
	adminAreaMiddleware := middleware.AdminAreaMiddleware()
	router := NewRouter(authMiddleware, adminAreaMiddleware, tournamentHandler, userHandler)

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run(":8080")
}
