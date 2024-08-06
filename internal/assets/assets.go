package assets

import (
    "encoding/json"
    "os"
    "log"
    "path/filepath"
)

// Represents single entry in the manifest file
//
// Fields:
// - File (string): Path to the main asset file
// - CSS ([]string): A list of associated CSS files
//
// Since: 0.0.1
type ManifestEntry struct {
    File string   `json:"file"`
    CSS  []string `json:"css"`
}

// Map that holds the contents of the manifest file
//
// Keys are asset names, and the values are ManifestEntry structs.
//
// Since: 0.0.1
var manifest map[string]ManifestEntry

// Reads the manifest file located in the specified directory
// and parses the JSON content into the manifest map.
//
// Returns: void
//
// Since: 0.0.1
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

// Looks up the given asset in the manifest map and returns
// the path to its main file.
//
// Params:
// - asset (string): Name of the asset
//
// Returns:
// - string: Path to the asset file, or an empty string if the asset
// is not found
//
// Since: 0.0.1
func AssetPath(asset string) string {
    if entry, exists := manifest[asset]; exists {
        return entry.File
    }
    return ""
}

// Looks up the given asset in the manifest map and returns
// a list of its associated css files
//
// Params:
// - asset (string): Name of the asset
//
// Returns:
// []string: List of paths to associated CSS files, or nil if the
// asset is not found
//
// Since: 0.0.1
func CSSPaths(asset string) []string {
    if entry, exists := manifest[asset]; exists {
        return entry.CSS
    }
    return nil
}

