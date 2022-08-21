package server

import (
	"time"

	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/services"
)

func InitScaffoldingData(userService *services.UserService, tournamentService *services.TournamentService) error {
	user := dto.UserRegistrationRequestDTO{
		Email:     "svendcs@svendcs.com",
		Password:  "password",
		FirstName: "Svend Christian",
		LastName:  "Svendsen",
	}

	exists, err := userService.ExistsByEmail(user.Email)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	if _, err := userService.Register(&user); err != nil {
		return err
	}

	tournament := dto.TournamentCreationRequestDTO{Name: "Bjergkøbing Grand Prix", Location: "Bjergkøbing", Date: time.Now()}

	if _, err := tournamentService.Create(&tournament); err != nil {
		return err
	}

	return nil
}
