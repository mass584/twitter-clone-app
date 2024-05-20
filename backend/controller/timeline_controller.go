package controller

import (
	"github.com/gin-gonic/gin"

	"github.com/mass584/twitter-clone-app/backend/database"
	"github.com/mass584/twitter-clone-app/backend/infrastructure/repository"
)

type GetTimelineRequest struct {
	Per  int64 `form:"per" binding:"required"`
	Page int64 `form:"page" binding:"required"`
}

func GetTimeline(context *gin.Context) {
	user_id, err := repository.AuthenticateUnsafely(context)
	if err != nil {
		context.JSON(401, gin.H{"message": "unauthorized"})
	}

	var req GetTimelineRequest

	if err := context.ShouldBindQuery(&req); err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	repo := repository.NewTimelineRepository(database.DB)
	timeline, err := repo.Get(user_id, req.Per, req.Page)

	if err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	context.JSON(200, gin.H{"timeline": timeline})
	return
}
