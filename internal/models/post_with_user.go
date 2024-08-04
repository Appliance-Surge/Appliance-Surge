package models

import "time"

// PostWithUser represents a post along with the associated user
type PostWithUser struct {
    PostID int `db:"post_id"`
    UserID int `db:"user_id"`
    Username string `db:"username"`
    Title string `db:"title"`
    Type int `db:"type"`
    Brand string `db:"brand"`
    Model string `db:"model"`
    Content string `db:"content"`
    CreatedAt *time.Time `db:"created_at"`
    UpdatedAt *time.Time `db:"updated_at"`
}
