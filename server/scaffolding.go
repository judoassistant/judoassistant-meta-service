package server

import (
	"time"

	"github.com/benbjohnson/clock"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/service"
	"github.com/pkg/errors"
)

func InitScaffoldingData(userService service.UserService, tournamentService service.TournamentService, clock clock.Clock) error {
	userRequest := dto.UserRegistrationRequestDTO{
		Email:     "svendcs@svendcs.com",
		Password:  "password",
		FirstName: "Svend Christian",
		LastName:  "Svendsen",
	}

	exists, err := userService.ExistsByEmail(userRequest.Email)
	if err != nil {
		return errors.Wrap(err, "unable to get user")
	}

	if exists {
		return nil
	}

	user, err := userService.Register(&userRequest)
	if err != nil {
		return errors.Wrap(err, "unable to register user")
	}

	tournament := dto.TournamentCreationRequestDTO{Name: "Bjergkøbing Grand Prix", Location: "Bjergkøbing", Date: time.Now()}

	if _, err := tournamentService.Create(user, &tournament); err != nil {
		return errors.Wrap(err, "unable to create tournament")
	}

	return nil
}
