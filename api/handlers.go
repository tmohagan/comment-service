package api

import (
	"comment-service/models"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func createComment(collection *mongo.Collection) gin.HandlerFunc {
    return func(c *gin.Context) {
        var comment models.Comment
        if err := c.BindJSON(&comment); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        
        // Log the received comment data
        fmt.Printf("Received comment: %+v\n", comment)
        
        // Ensure postId is not empty
        if comment.PostID == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "postId is required"})
            return
        }
        
        comment.CreatedAt = time.Now()
        comment.UpdatedAt = time.Now()
        
        result, err := collection.InsertOne(context.Background(), comment)
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating comment"})
            return
        }
        
        c.JSON(http.StatusCreated, gin.H{"id": result.InsertedID})
    }
}

func getAllComments(collection *mongo.Collection) gin.HandlerFunc {
    return func(c *gin.Context) {
        postId := c.Query("postId")
        if postId == "" {
            c.JSON(http.StatusBadRequest, gin.H{"error": "postId is required"})
            return
        }

        cursor, err := collection.Find(context.Background(), bson.M{"postId": postId})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching comments"})
            return
        }
        defer cursor.Close(context.Background())
        
        var comments []models.Comment
        if err = cursor.All(context.Background(), &comments); err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding comments"})
            return
        }
        
        c.JSON(http.StatusOK, comments)
    }
}

func getComment(collection *mongo.Collection) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := primitive.ObjectIDFromHex(c.Param("id"))
        var comment models.Comment
        err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&comment)
        if err != nil {
            c.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
            return
        }
        c.JSON(http.StatusOK, comment)
    }
}

func updateComment(collection *mongo.Collection) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := primitive.ObjectIDFromHex(c.Param("id"))
        var comment models.Comment
        if err := c.BindJSON(&comment); err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        
        comment.UpdatedAt = time.Now()
        
        _, err := collection.UpdateOne(
            context.Background(),
            bson.M{"_id": id},
            bson.D{
                {"$set", bson.D{
                    {"content", comment.Content},
                    {"updated_at", comment.UpdatedAt},
                }},
            },
        )
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error updating comment"})
            return
        }
        
        c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
    }
}

func deleteComment(collection *mongo.Collection) gin.HandlerFunc {
    return func(c *gin.Context) {
        id, _ := primitive.ObjectIDFromHex(c.Param("id"))
        _, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "Error deleting comment"})
            return
        }
        c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
    }
}