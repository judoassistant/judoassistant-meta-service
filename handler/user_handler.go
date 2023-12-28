package handler

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/judoassistant/judoassistant-meta-service/dto"
	"github.com/judoassistant/judoassistant-meta-service/errors"
	"github.com/judoassistant/judoassistant-meta-service/middleware"
	"github.com/judoassistant/judoassistant-meta-service/service"
	"go.uber.org/zap"
)

const _minPasswordLength = 8

type UserHandler interface {
	Create(c *gin.Context) error
	Get(c *gin.Context) error
	Index(c *gin.Context) error
	Update(c *gin.Context) error
	UpdatePassword(c *gin.Context) error
	Authenticate(c *gin.Context) error
}

type userHandler struct {
	userService service.UserService
	logger      *zap.Logger
}

func NewUserHandler(userService service.UserService, logger *zap.Logger) UserHandler {
	return &userHandler{
		userService: userService,
		logger:      logger,
	}
}

func (handler *userHandler) Index(c *gin.Context) error {
	users, err := handler.userService.GetAll()
	if err != nil {
		return errors.Wrap(err, "unable to get users")
	}

	c.JSON(http.StatusOK, users)
	return nil
}

func (handler *userHandler) Create(c *gin.Context) error {
	// Parse input
	request := dto.UserRegistrationRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	// Validate input
	if err := validateEmail(request.Email); err != nil {
		return err
	}
	if err := validatePassword(request.Password); err != nil {
		return err
	}

	// Create
	response, err := handler.userService.Register(&request)
	if err != nil {
		return errors.Wrap(err, "unable to register user")
	}

	c.JSON(http.StatusOK, response)
	return nil
}

func (handler *userHandler) Get(c *gin.Context) error {
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)
	if query.ID != authorizedUser.ID {
		return errors.New("cannot access other users", errors.CodeForbidden)
	}

	user, err := handler.userService.GetById(query.ID)
	if err != nil {
		return errors.Wrap(err, "unable to get user")
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (handler *userHandler) UpdatePassword(c *gin.Context) error {
	// Parse input
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	request := dto.UserPasswordUpdateRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	// Validate input
	if err := validatePassword(request.Password); err != nil {
		return err
	}

	// Update password
	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)
	if query.ID != authorizedUser.ID {
		return errors.New("cannot access other users", errors.CodeForbidden)
	}

	user, err := handler.userService.UpdatePassword(query.ID, request.Password)
	if err != nil {
		return errors.Wrap(err, "unable to update password")
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (handler *userHandler) Update(c *gin.Context) error {
	// Parse input
	query := dto.UserQueryDTO{}
	if err := c.ShouldBindQuery(&query); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	request := dto.UserUpdateRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	// Validate input
	if err := validateEmail(request.Email); err != nil {
		return err
	}

	// Update
	authorizedUser := c.MustGet(middleware.AuthUserKey).(*dto.UserResponseDTO)
	if query.ID != authorizedUser.ID {
		return errors.New("cannot access other users", errors.CodeForbidden)
	}

	user, err := handler.userService.Update(query.ID, &request)
	if err != nil {
		return errors.Wrap(err, "unable to update user")
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func (handler *userHandler) Authenticate(c *gin.Context) error {
	// Parse input
	request := dto.UserAuthenticationRequestDTO{}
	if err := c.ShouldBindJSON(&request); err != nil {
		return errors.WrapWithCode(err, "unable to map request", errors.CodeBadRequest)
	}

	// Validate input
	if err := validateEmail(request.Email); err != nil {
		return err
	}
	if err := validatePassword(request.Password); err != nil {
		return err
	}

	// Authenticate
	user, err := handler.userService.Authenticate(&request)
	if err != nil {
		return errors.Wrap(err, "unable to get user")
	}

	c.JSON(http.StatusOK, user)
	return nil
}

func validateEmail(email string) error {
	address, err := mail.ParseAddress(email)
	if err != nil {
		return errors.WrapWithCode(err, "invalid email", errors.CodeBadRequest)
	}
	if address.Address != email {
		// ParseAddress accepts any RFC5322 format. We accept only the subset, where
		// the entire string is the email.
		return errors.New("invalid email", errors.CodeBadRequest)
	}

	return nil
}

func validatePassword(password string) error {
	if len(password) < _minPasswordLength {
		return errors.New(fmt.Sprintf("invalid password; passwords must be at least %d characters", _minPasswordLength), errors.CodeBadRequest)
	}
	return nil
}
