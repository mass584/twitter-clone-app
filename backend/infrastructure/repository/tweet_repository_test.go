package repository_test

import (
	"testing"

	"github.com/mass584/twitter-clone-app/backend/database"
	"github.com/mass584/twitter-clone-app/backend/domain/entity"
	"github.com/mass584/twitter-clone-app/backend/infrastructure/repository"
)

func TestStore(t *testing.T) {
	db := database.DB
	repo := repository.NewTweetRepository(db)
	tweet := entity.Tweet{Id: 0, TextContents: "This is a tweet text."}
	repo.Store(&tweet)
	if tweet.Id == 0 {
		t.Errorf("Expected not zero, but got zero")
	}
}
