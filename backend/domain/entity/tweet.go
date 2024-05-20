package entity

import (
	"fmt"
	"time"
	"unicode/utf8"
)

type Tweet struct {
	Id           int64
	UserId       int64
	User         *User
	TextContents string
	TweetAt      time.Time
}

func NewTweet(
	Id int64,
	UserId int64,
	User *User,
	TextContents string,
	TweetAt time.Time,
) (*Tweet, error) {
	if utf8.RuneCountInString(TextContents) > 255 {
		return nil, fmt.Errorf("TextContents is too long.")
	}

	if utf8.RuneCountInString(TextContents) == 0 {
		return nil, fmt.Errorf("TextContents is empty.")
	}

	tweet := Tweet{
		Id:           Id,
		UserId:       UserId,
		User:         User,
		TextContents: TextContents,
		TweetAt:      TweetAt,
	}

	return &tweet, nil
}
