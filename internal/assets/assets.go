package assets

import (
    "encoding/json"
    "os"
    "log"
    "path/filepath"
)

type ManifestEntry struct {
    File string   `json:"file"`
    CSS  []string `json:"css"`
}

var manifest map[string]ManifestEntry

func LoadManifest() {
    manifestFile := filepath.Join("static", ".vite", "manifest.json")
    data, err := os.ReadFile(manifestFile)
    if err != nil {
        log.Fatalf("Failed to read manifest file: %v", err)
    }

    err = json.Unmarshal(data, &manifest)
    if err != nil {
        log.Fatalf("Failed to parse manifest file: %v", err)
    }
}

func AssetPath(asset string) string {
    if entry, exists := manifest[asset]; exists {
        return entry.File
    }
    return ""
}

func CSSPaths(asset string) []string {
    if entry, exists := manifest[asset]; exists {
        return entry.CSS
    }
    return nil
}

