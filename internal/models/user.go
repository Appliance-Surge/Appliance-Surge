package models

import "time"

// User represents a user in the application
type User struct {
	ID              int        `db:"id"`
	Username        string     `db:"username"`
	Name            string     `db:"name"`
	Image           string     `db:"image"`
	Banner          string     `db:"banner"`
	Score           int        `db:"score"`
	Location        string     `db:"location"`
	About           string     `db:"about"`
	Website         string     `db:"website"`
	SocialID        string     `db:"social_id"`
	SocialType      string     `db:"social_type"`
	Email           string     `db:"email"`
	EmailVerifiedAt *time.Time `db:"email_verified_at"`
	Password        string     `db:"password"`
	RememberToken   string     `db:"remember_token"`
	CreatedAt       *time.Time `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
}
