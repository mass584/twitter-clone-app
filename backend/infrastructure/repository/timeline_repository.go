package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/mass584/twitter-clone-app/backend/domain/entity"
	"github.com/mass584/twitter-clone-app/backend/domain/repository_if"
)

type TimelineRepository struct {
	database *sql.DB
}

func NewTimelineRepository(db *sql.DB) repository_if.TimelineRepository {
	return &TimelineRepository{database: db}
}

func (r *TimelineRepository) Get(user_id int64, per int64, page int64) (*entity.Timeline, error) {
	statement, _ := r.database.Prepare(
		"select tweets.id, tweets.user_id, tweets.text_contents, tweets.tweet_at, users.display_name from tweets " +
			"join users on tweets.user_id = users.id " +
			"where tweets.user_id in (select followed_user_id from follows where follower_user_id = ?) " +
			"order by tweets.tweet_at desc " +
			"limit ? offset ? ",
	)
	defer statement.Close()

	limit := per
	offset := per * (page - 1)
	rows, _ := statement.Query(user_id, limit, offset)
	defer rows.Close()

	var tweets []entity.Tweet = []entity.Tweet{}

	for rows.Next() {
		var id int64
		var user_id int64
		var text_contents string
		var tweet_at_str string
		var display_name string

		err := rows.Scan(&id, &user_id, &text_contents, &tweet_at_str, &display_name)
		if err != nil {
			log.Fatal(err)
		}

		location, _ := time.LoadLocation("UTC")
		tweet_at, _ := time.ParseInLocation("2006-01-02 15:04:05", tweet_at_str, location)

		user, _ := entity.NewUser(user_id, display_name)
		tweet, _ := entity.NewTweet(id, user_id, user, text_contents, tweet_at)
		tweets = append(tweets, *tweet)
	}

	timeline, _ := entity.NewTimeline(per, page, tweets)

	return timeline, nil
}
