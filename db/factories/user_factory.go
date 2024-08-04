package factories

import (
    "math/rand"
    "time"

    "github.com/go-faker/faker/v4"

    "github.com/Appliance-Surge/Appliance-Surge/internal/storage"
    "github.com/Appliance-Surge/Appliance-Surge/internal/models"
)

func CreateUserFactory(db *storage.DB) error {
    now := time.Now()

    // Generate fake data for a user
    user := models.User{
        Username: faker.Username(),
        Name: faker.Name(),
        Image: faker.URL(),
        Banner: faker.URL(),
        Score: rand.Intn(100),
        Location: faker.Word(),
        About: faker.Paragraph(),
        Website: faker.URL(),
        SocialID: faker.UUIDDigit(),
        SocialType: "Google",
        Email: faker.Email(),
        Password: faker.Password(),
        RememberToken: faker.UUIDHyphenated(),
        CreatedAt: &now,
        UpdatedAt: &now,
    }

    // Insert the user into the database
    return db.CreateUser(&user)
}
