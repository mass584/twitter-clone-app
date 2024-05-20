package repository_if

import (
	"github.com/mass584/twitter-clone-app/backend/domain/entity"
)

type TweetRepository interface {
	Store(tweet *entity.Tweet) error
}
