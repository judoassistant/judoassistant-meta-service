package server

import (
	"time"

	"github.com/benbjohnson/clock"
	"github.com/judoassistant/judoassistant-meta-service/config"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/service"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func InitScaffoldingData(userService service.UserService, tournamentService service.TournamentService, config *config.Config, logger *zap.Logger, clock clock.Clock) error {
	exists, err := userService.ExistsByEmail(config.AdminEmail)
	if err != nil {
		return errors.Wrap(err, "unable to get user")
	}
	if exists {
		return nil
	}

	logger.Info("Scaffolding empty database")
	userRequest := dto.UserRegistrationRequestDTO{
		Email:     config.AdminEmail,
		Password:  config.AdminDefaultPassword,
		FirstName: config.AdminDefaultFirstName,
		LastName:  config.AdminDefaultLastName,
		IsAdmin:   true,
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
