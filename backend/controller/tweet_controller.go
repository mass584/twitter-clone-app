package controller

import (
	"time"

	"github.com/gin-gonic/gin"

	"github.com/mass584/twitter-clone-app/backend/database"
	"github.com/mass584/twitter-clone-app/backend/domain/entity"
	"github.com/mass584/twitter-clone-app/backend/infrastructure/repository"
)

type CreateTweetRequest struct {
	TextContents string `json:"text_contents"`
}

func CreateTweet(context *gin.Context) {
	user_id, err := repository.AuthenticateUnsafely(context)
	if err != nil {
		context.JSON(401, gin.H{"message": "unauthorized"})
	}

	var req CreateTweetRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	repo := repository.NewTweetRepository(database.DB)

	tweet, err := entity.NewTweet(0, user_id, nil, req.TextContents, time.Now().UTC())

	if err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	if err := repo.Store(tweet); err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	context.JSON(201, gin.H{"tweet": tweet})
	return
}
