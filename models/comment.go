package models

import (
    "time"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Comment struct {
    ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
    PostID    string             `bson:"post_id" json:"post_id"`
    Author    string             `bson:"author" json:"author"`
    Content   string             `bson:"content" json:"content"`
    CreatedAt time.Time          `bson:"created_at" json:"created_at"`
    UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
}