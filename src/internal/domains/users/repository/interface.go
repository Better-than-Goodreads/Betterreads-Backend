package repository

import (
	"errors"

	"github.com/betterreads/internal/domains/users/models"
	"github.com/google/uuid"
)

var (
	ErrUserNotFound         = errors.New("user not found")
	ErrUserStageNotFound    = errors.New("user stage not found")
	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrEmailAlreadyTaken    = errors.New("email already taken")
)

type UsersDatabase interface {
	CreateStageUser(user *models.UserStageRequest) (*models.UserStageRecord, error)
	JoinAndCreateUser(userAddtional *models.UserAdditionalRequest, id uuid.UUID) (*models.UserRecord, error)
	GetUser(id uuid.UUID) (*models.UserRecord, error)
	GetUsers() ([]*models.UserRecord, error)
	GetStageUser(id uuid.UUID) (*models.UserStageRecord, error)
	GetUserByUsername(username string) (*models.UserRecord, error)
	GetUserByEmail(email string) (*models.UserRecord, error)
	GetUserPicture(id uuid.UUID) ([]byte, error)
	CheckUserExistsForRegister(user *models.UserStageRequest) error
	CheckUserExists(id uuid.UUID) bool
	SaveUserPicture(id uuid.UUID, picture []byte) error
	SearchUsers(username string, isAuthor bool) ([]*models.UserRecord, error)
}
