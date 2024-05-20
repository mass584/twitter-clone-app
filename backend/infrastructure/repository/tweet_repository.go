package repository

import (
	"database/sql"

	"github.com/mass584/twitter-clone-app/backend/domain/entity"
	"github.com/mass584/twitter-clone-app/backend/domain/repository_if"
)

type TweetRepository struct {
	database *sql.DB
}

func NewTweetRepository(db *sql.DB) repository_if.TweetRepository {
	return &TweetRepository{database: db}
}

func (r *TweetRepository) Store(tweet *entity.Tweet) error {
	if tweet.Id == 0 {
		statement, _ := r.database.Prepare("insert into tweets(user_id, text_contents, tweet_at) values(?, ?, ?)")
		defer statement.Close()
		result, err := statement.Exec(tweet.UserId, tweet.TextContents, tweet.TweetAt)
		if err != nil {
			return err
		}
		newTweetId, _ := result.LastInsertId()
		tweet.Id = newTweetId
	} else {
		panic("Update action is not implemented.")
	}

	return nil
}
