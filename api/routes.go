package api

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"os"
	"strings"
)

func SetupRoutes(router *gin.Engine, client *mongo.Client) {
	commentsCollection := client.Database("comment-service").Collection("comments")

	config := cors.DefaultConfig()
	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins != "" {
		config.AllowOrigins = strings.Split(allowedOrigins, ",")
	} else {
		config.AllowAllOrigins = true
	}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour
	router.Use(cors.New(config))
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})
	router.POST("/comments", createComment(commentsCollection))
	router.GET("/comments", getAllComments(commentsCollection))
	router.GET("/comments/:id", getComment(commentsCollection))
	router.PUT("/comments/:id", updateComment(commentsCollection))
	router.DELETE("/comments/:id", deleteComment(commentsCollection))
}