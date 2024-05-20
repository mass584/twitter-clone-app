package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mass584/twitter-clone-app/backend/database"
	"github.com/mass584/twitter-clone-app/backend/infrastructure/repository"
)

type CreateFollowableUserRequest struct {
	TargetUserId int64 `json:"target_user_id"`
}

func CreateFollowableUser(context *gin.Context) {
	user_id, err := repository.AuthenticateUnsafely(context)
	if err != nil {
		context.JSON(401, gin.H{"message": "unauthorized"})
	}

	var req CreateFollowableUserRequest
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	repo := repository.NewFollowableUserRepository(database.DB)

	err = repo.Create(user_id, req.TargetUserId)
	if err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	context.Status(204)
	return
}

func DeleteFollowableUser(context *gin.Context) {
	user_id, err := repository.AuthenticateUnsafely(context)
	if err != nil {
		context.JSON(401, gin.H{"message": "unauthorized"})
	}

	target_user_id, err := strconv.ParseInt(context.Param("target_user_id"), 10, 64)
	if err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	repo := repository.NewFollowableUserRepository(database.DB)

	if err := repo.Delete(user_id, target_user_id); err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	context.Status(204)
	return
}

func ListFollowableUsers(context *gin.Context) {
	user_id, err := repository.AuthenticateUnsafely(context)
	if err != nil {
		context.JSON(401, gin.H{"message": "unauthorized"})
	}

	display_name := context.Query("display_name")

	repo := repository.NewFollowableUserRepository(database.DB)

	followable_users, _ := repo.FindListByQuery(user_id, display_name)

	context.JSON(200, gin.H{"followable_users": followable_users})
	return
}
