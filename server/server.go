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
	database, err := db.Init()
	if err != nil {
		log.Fatalln(err.Error())
	}

	if err := db.Migrate(database); err != nil {
		log.Println("Test1")
		log.Fatalln(err.Error())
	}
	log.Println("Test")

	tournamentRepository := repositories.NewTournamentRepository(database)
	tournamentService := services.NewTournamentService(tournamentRepository)
	tournamentController := controllers.NewTournamentController(tournamentService)

	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	authMiddleware := middleware.BasicAuthMiddleware(userService)
	router := NewRouter(authMiddleware, tournamentController, userController)

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run("localhost:8080")
}
