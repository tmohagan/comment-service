package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    PostID    string             `bson:"post_id" json:"postId"`
    UserID    string             `bson:"user_id" json:"userId"`
    Username  string             `bson:"username" json:"username"`
    Content   string             `bson:"content" json:"content"`
    CreatedAt time.Time          `bson:"created_at" json:"createdAt"`
    UpdatedAt time.Time          `bson:"updated_at" json:"updatedAt"`
}