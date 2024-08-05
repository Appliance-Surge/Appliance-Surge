package storage

import (
    "github.com/Appliance-Surge/Appliance-Surge/internal/models"
)

// CreateUser inserts a new user into the database
func (db *DB) CreateUser(user *models.User) error {
    query := `
        INSERT INTO users (
            username, name, image, banner, score, location, about, website, social_id, social_type, email, email_verified_at, password, remember_token, created_at, updated_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16
        )
        RETURNING id
    `
    err := db.QueryRow(
        query,
        user.Username, user.Name, user.Image, user.Banner, user.Score, user.Location, user.About, user.Website, user.SocialID, user.SocialType, user.Email, user.EmailVerifiedAt, user.Password, user.RememberToken, user.CreatedAt, user.UpdatedAt,
    ).Scan(&user.ID)
    return err
}

