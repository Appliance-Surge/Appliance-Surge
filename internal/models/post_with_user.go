package models

import "time"

// Represents Posts with Associated Authors
//
// Fields:
// - POSTID (int): The post ID
// - UserID (int): The associated User ID
// - Username (string): The author's username
// - UserScore (int): The user's score
// - Title (string): The post title
// - Type (int): The type of appliance
// - Brand (string): The brand of appliance
// - Model (string): The model of appliance
// - Content (string): The text content of the post
// - CreatedAt (*time.Time): When the post was created
// - UpdatedAt (*time.Time): When the post was last
// updated
//
// Since: 0.0.1
type PostWithUser struct {
    PostID int `db:"post_id"`
    UserID int `db:"user_id"`
    Username string `db:"username"`
    UserScore int `db:"score"`
    Title string `db:"title"`
    Type int `db:"type"`
    Brand string `db:"brand"`
    Model string `db:"model"`
    Content string `db:"content"`
    CreatedAt *time.Time `db:"created_at"`
    UpdatedAt *time.Time `db:"updated_at"`
}
