package server

import (
	"log"

	"github.com/judoassistant/judoassistant-meta-service/controllers"
	"github.com/judoassistant/judoassistant-meta-service/db"
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
	router := NewRouter(tournamentController)

	router.Run("localhost:8080")
}
