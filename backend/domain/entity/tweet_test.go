package entity_test

import (
	"testing"
	"time"

	"github.com/mass584/twitter-clone-app/backend/domain/entity"
)

func TestNewTweet(t *testing.T) {
	var error error

	_, error = entity.NewTweet(0, 0, nil, "", time.Now())
	if error.Error() != "TextContents is empty." {
		t.Errorf("Expected return empty error, but not return")
	}

	_, error = entity.NewTweet(0, 0, nil, "1", time.Now())
	if error != nil {
		t.Errorf("Expected not return error, but return error")
	}

	long_text := ""
	for i := 0; i < 256; i++ {
		long_text += "A"
	}
	_, error = entity.NewTweet(0, 0, nil, long_text, time.Now())
	if error.Error() != "TextContents is too long." {
		t.Errorf("Expected return too long error, but not return")
	}
}
