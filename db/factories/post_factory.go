package factories

import (
    "math/rand"
    "time"

    "github.com/go-faker/faker/v4"

    "github.com/Appliance-Surge/Appliance-Surge/internal/storage"
    "github.com/Appliance-Surge/Appliance-Surge/internal/models"
)

func CreatePostFactory(db *storage.DB) error {
    now := time.Now()

    // Generate fake data for a post
    post := models.Post{
        UserID: rand.Intn(10) + 1,
        Title: faker.Word(),
        Type: rand.Intn(10),
        Brand: faker.Word(),
        Model: faker.Word(),
        Content: faker.Paragraph(),
        CreatedAt: &now,
        UpdatedAt: &now,
    }

    // Insert the user into the database
    return db.CreatePost(&post)
}
