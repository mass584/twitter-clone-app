package repository_if

import (
	"github.com/mass584/twitter-clone-app/backend/domain/entity"
)

type TimelineRepository interface {
	Get(user_id int64, per int64, page int64) (*entity.Timeline, error)
}
