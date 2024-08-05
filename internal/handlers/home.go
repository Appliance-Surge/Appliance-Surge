package handlers

import (
    "net/http"
    "html/template"
    "log"

    "github.com/Appliance-Surge/Appliance-Surge/internal/storage"
)

var tmpl = template.Must(template.ParseFiles(
    "templates/layouts/main_layout.html",
    "templates/home.html",
))

func HomeHandler(db *storage.DB, assetsData map[string]interface{}) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postsWithUsers, err := db.FetchPostsWithUsers()
        if err != nil {
            log.Printf("Failed to fetch posts with users: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        data := map[string]interface{}{
            "PostsWithUsers": postsWithUsers,
            "CSSPaths":       assetsData["CSSPaths"],
            "JSPath":         assetsData["JSPath"],
        }

        if err := tmpl.ExecuteTemplate(w, "main_layout", data); err != nil {
            log.Printf("Failed to execute template: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        }
    }
}
