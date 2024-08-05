package handlers

import (
    "net/http"
    "html/template"
    "log"

    "github.com/Appliance-Surge/Appliance-Surge/internal/storage"
    "github.com/Appliance-Surge/Appliance-Surge/internal/assets"
    "github.com/Appliance-Surge/Appliance-Surge/internal/utils"
)

var tmpl = template.Must(template.New("").Funcs(utils.FuncMap()).ParseFiles(
    "templates/layouts/main_layout.html",
    "templates/home.html",
    "templates/components/navigation.html",
    "templates/components/primary-link.html",
    "templates/components/responsive-nav-link.html",
))

func HomeHandler(db *storage.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        postsWithUsers, err := db.FetchPostsWithUsers()
        if err != nil {
            log.Printf("Failed to fetch posts with users: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
            return
        }

        data := map[string]interface{}{
            "PostsWithUsers": postsWithUsers,
            "CSSPaths":       assets.CSSPaths("resources/js/main.js"),
            "JSPath":         assets.AssetPath("resources/js/main.js"),
        }

        if err := tmpl.ExecuteTemplate(w, "main_layout", data); err != nil {
            log.Printf("Failed to execute template: %v", err)
            http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        }
    }
}
