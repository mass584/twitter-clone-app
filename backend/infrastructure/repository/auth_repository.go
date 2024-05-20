package repository

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AuthenticateUnsafely(context *gin.Context) (int64, error) {
	var userId int64 = 0
	cookies := context.Request.Cookies()

	for _, cookie := range cookies {
		if cookie.Name == "user_id" {
			userId, _ = strconv.ParseInt(cookie.Value, 10, 64)
		}
	}

	if userId == 0 {
		return userId, fmt.Errorf("Unauthenticated.")
	}

	return userId, nil
}
