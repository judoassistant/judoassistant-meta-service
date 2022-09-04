package server

import (
	"time"

	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/service"
)

func InitScaffoldingData(userService service.UserService, tournamentService service.TournamentService) error {
	userRequest := dto.UserRegistrationRequestDTO{
		Email:     "svendcs@svendcs.com",
		Password:  "password",
		FirstName: "Svend Christian",
		LastName:  "Svendsen",
	}

	exists, err := userService.ExistsByEmail(userRequest.Email)
	if err != nil {
		return err
	}

	if exists {
		return nil
	}

	user, err := userService.Register(&userRequest)
	if err != nil {
		return err
	}

	tournament := dto.TournamentCreationRequestDTO{Name: "Bjergkøbing Grand Prix", Location: "Bjergkøbing", Date: time.Now()}

	if _, err := tournamentService.Create(user, &tournament); err != nil {
		return err
	}

	return nil
}
