package storage

import (
    "database/sql"
    "log"

    _ "github.com/lib/pq"

    "github.com/Appliance-Surge/Appliance-Surge/internal/models"
)

type DB struct {
    *sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
    db, err := sql.Open("postgres", dataSourceName)
    if err != nil {
        return nil, err
    }

    // Verify the connection
    if err := db.Ping(); err != nil {
        return nil, err
    }

    log.Println("Successfully connected to the database")
    return &DB{db}, nil
}

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

// CreatePost inserts a new post into the database
func (db *DB) CreatePost(post *models.Post) error {
    query := `
        INSERT INTO posts (
            user_id, title, type, brand, model, content, created_at, updated_at, deleted_at
        ) VALUES (
            $1, $2, $3, $4, $5, $6, $7, $8, $9
        )
        RETURNING id
    `
    err := db.QueryRow(
        query,
        post.UserID, post.Title, post.Type, post.Brand, post.Model, post.Content, post.CreatedAt, post.UpdatedAt, post.DeletedAt,
    ).Scan(&post.ID)
    return err
}

// FetchPostsWithUsers fetches all posts along with their associated users
func (db *DB) FetchPostsWithUsers() ([]models.PostWithUser, error) {
    query := `
        SELECT
            posts.id AS post_id,
            posts.user_id,
            users.username,
            posts.title,
            posts.type,
            posts.brand,
            posts.model,
            posts.content,
            posts.created_at,
            posts.updated_at
        FROM
            posts
        JOIN
            users ON posts.user_id = users.id
        ORDER BY
            posts.created_at DESC
    `

    rows, err := db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var postsWithUsers []models.PostWithUser
    for rows.Next() {
        var postWithUser models.PostWithUser
        if err := rows.Scan(
            &postWithUser.PostID,
            &postWithUser.UserID,
            &postWithUser.Username,
            &postWithUser.Title,
            &postWithUser.Type,
            &postWithUser.Brand,
            &postWithUser.Model,
            &postWithUser.Content,
            &postWithUser.CreatedAt,
            &postWithUser.UpdatedAt,
        ); err != nil {
            return nil, err
        }
        postsWithUsers = append(postsWithUsers, postWithUser)
    }

    return postsWithUsers, nil
}
