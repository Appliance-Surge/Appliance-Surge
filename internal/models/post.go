package models

import "time"

// Post represents a post in the application
type Post struct {
	ID        int        `db:"id"`
	UserID    int        `db:"user_id"`
	Title     string     `db:"title"`
	Type      int        `db:"type"`
	Brand     string     `db:"brand"`
	Model     string     `db:"model"`
	Content   string     `db:"content"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
	DeletedAt *time.Time `db:"deleted_at"`
}
