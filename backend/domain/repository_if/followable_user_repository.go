package repository_if

import "github.com/mass584/twitter-clone-app/backend/domain/entity"

type FollowedUserRepository interface {
	Create(user_id int64, target_user_id int64) error
	Delete(user_id int64, target_user_id int64) error
	FindListByQuery(user_id int64, query string) (*[]entity.FollowableUser, error)
}
