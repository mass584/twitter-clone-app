package repository_if

import (
	"github.com/mass584/twitter-clone-app/backend/domain/entity"
)

type UserRepository interface {
	FindOneByEmail(email string) (*entity.User, error)
}
