package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/mass584/twitter-clone-app/backend/config"
	"github.com/mass584/twitter-clone-app/backend/controller"
	_ "github.com/mass584/twitter-clone-app/backend/database"
)

func main() {
	router := gin.Default()

	config := config.New()

	cors_config := cors.DefaultConfig()
	cors_config.AllowOrigins = []string{config.ClientHost()}
	cors_config.AllowCredentials = true
	router.Use(cors.New(cors_config))

	router.POST("/login", controller.CreateSession)
	router.POST("/tweets", controller.CreateTweet)
	router.GET("/timeline", controller.GetTimeline)
	router.GET("/followable_users", controller.ListFollowableUsers)
	router.POST("/followable_users", controller.CreateFollowableUser)
	router.DELETE("/followable_users/:target_user_id", controller.DeleteFollowableUser)
	router.Run()
}
