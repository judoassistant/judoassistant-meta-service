package server

import (
	"log"

	"github.com/judoassistant/judoassistant-meta-service/controllers"
	"github.com/judoassistant/judoassistant-meta-service/db"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/repositories"
	"github.com/judoassistant/judoassistant-meta-service/services"
)

func Init() {
	db, err := db.Init()
	if err != nil {
		log.Fatalln(err)
	}

	tournamentRepository := repositories.NewTournamentRepository(db)
	tournamentService := services.NewTournamentService(tournamentRepository)
	tournamentController := controllers.NewTournamentController(tournamentService)

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	authMiddleware := middleware.BasicAuthMiddleware(userService)
	router := NewRouter(authMiddleware, tournamentController, userController)

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run("localhost:8080")
}
