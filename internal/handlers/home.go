package handlers

import (
    "net/http"
    "html/template"
    "log"

    "github.com/Appliance-Surge/Appliance-Surge/internal/storage"
    "github.com/Appliance-Surge/Appliance-Surge/internal/assets"
    "github.com/Appliance-Surge/Appliance-Surge/internal/utils"
    "github.com/Appliance-Surge/Appliance-Surge/internal/models"
)

var tmpl *template.Template

// Parse the html templates required for the home page
// Load utility functions to be used within the templates
//
// @since 0.0.1
func fetchTemplates() {
    var err error
    tmpl, err = template.New("").Funcs(utils.FuncMap()).ParseFiles(
        "templates/layouts/main_layout.html",
        "templates/home.html",
    )
    if err != nil {
        log.Fatalf("Failed to parse base templates: %v", err)
    }

    tmpl, err = tmpl.ParseGlob("templates/components/*.html")
    if err != nil {
        log.Fatalf("Failed to parse component templates: %v", err)
    }
}

// Fetch data and render home view
//
// Params:
// - db (*storage.DB): The DB connection
//
// Returns:
// - http.HandlerFunc: http handler function
//
// Since: 0.0.1
func HomeHandler(db *storage.DB) http.HandlerFunc {
    fetchTemplates()
    return func(w http.ResponseWriter, r *http.Request) {
        postsWithUsers, err := fetchData(db, w)
        if err != nil {
            return
        }

        data := structureData(postsWithUsers)

        renderView(w, "main_layout", data)
    }
}

// Fetch Posts with Users
//
// Params:
// - db (*storage.DB): The DB connection
// - w (http.ResponseWriter): The response writer
//
// Returns:
// - []models.PostWithUser: If no error List of posts with their
// associated authors
// - error: If fetch fails, return error
//
// Since: 0.0.1
func fetchData(db *storage.DB, w http.ResponseWriter) ([]models.PostWithUser, error) {
    postsWithUsers, err := db.FetchPostsWithUsers()
    if err != nil {
        log.Printf("Failed to fetch posts with users: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return nil, err
    }
    return postsWithUsers, nil
}

// Structure the data to be sent to template
//
// Params:
// - postsWithUsers ([]models.PostWithUser): List of posts with users
//
// Returns:
// - map[string]interface{}: map of data including postsWithUsers, css paths, and js paths
//
// Since: 0.0.1
func structureData(postsWithUsers []models.PostWithUser) (map[string]interface{}) {
    data := map[string]interface{}{
        "PostsWithUsers": postsWithUsers,
        "CSSPaths": assets.CSSPaths("resources/js/main.js"),
        "JSPath": assets.AssetPath("resources/js/main.js"),
    }

    return data
}

// Render the HTML view
//
// Params:
// w (http.ResponseWriter): The response writer
// template (string): Base HTML template name
// data (map[string]interface{}): Data to be sent to templates
//
// Returns: void
//
// Since: 0.0.1
func renderView(w http.ResponseWriter, template string, data map[string]interface{}) {
    if err := tmpl.ExecuteTemplate(w, template, data); err != nil {
        log.Printf("Failed to execute template: %v", err)
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
    }
}
