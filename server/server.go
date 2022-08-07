package server

import (
	"log"

	"github.com/judoassistant/judoassistant-meta-service/controllers"
	"github.com/judoassistant/judoassistant-meta-service/db"
	"github.com/judoassistant/judoassistant-meta-service/dto"
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
		log.Fatalln(err.Error())
	}

	tournamentRepository := repositories.NewTournamentRepository(database)
	tournamentService := services.NewTournamentService(tournamentRepository)
	tournamentController := controllers.NewTournamentController(tournamentService)

	userRepository := repositories.NewUserRepository(database)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	if err := InitScaffoldingData(userService); err != nil {
		log.Fatalln(err)
	}

	authMiddleware := middleware.BasicAuthMiddleware(userService)
	adminAreaMiddleware := middleware.AdminAreaMiddleware()
	router := NewRouter(authMiddleware, adminAreaMiddleware, tournamentController, userController)

	router.SetTrustedProxies([]string{"127.0.0.1"})
	router.Run(":8080")
}

func InitScaffoldingData(userService *services.UserService) error {
	exists, err := userService.ExistsByEmail("svendcs@svendcs.com")
	if err != nil {
		log.Fatalln(err.Error())
	}

	if exists {
		return nil
	}

	user := dto.UserRegistrationRequestDTO{
		Email:     "svendcs@svendcs.com",
		Password:  "password",
		FirstName: "Svend Christian",
		LastName:  "Svendsen",
	}
	if _, err := userService.Register(&user); err != nil {
		return err
	}

	return nil
}
