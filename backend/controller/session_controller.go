package controller

import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/mass584/twitter-clone-app/backend/config"
	"github.com/mass584/twitter-clone-app/backend/database"
	"github.com/mass584/twitter-clone-app/backend/infrastructure/repository"
)

type CreateSessionRequest struct {
	Email string `json:"email" binding:"required"`
}

func CreateSession(context *gin.Context) {
	var req CreateSessionRequest

	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	repo := repository.NewUserRepository(database.DB)
	user, err := repo.FindOneByEmail(req.Email)

	if err != nil {
		context.JSON(401, gin.H{"message": err.Error()})
		return
	}

	// TODO 認証後に発行されたセッションを保存すべきだが、一旦直接user_idを保存しておく。
	// user_idさえ知っていれば簡単になりすましできてしまうので、認証を実装すること。
	userId := strconv.FormatInt(user.Id, 10)
	config := config.New()
	// プロトコルがhttpsならcookieのsecure属性はonにしておく
	secure := config.ClientProto == "https"
	// cookieの期限は無制限にしておく
	max_age := math.MaxInt32
	// jsからはcookieをさわれないようにしておく
	http_only := true
	context.SetCookie("user_id", userId, max_age, "/", config.ClientDomain, secure, http_only)
	context.Status(204)
	return
}
