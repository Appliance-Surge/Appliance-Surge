package assets

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

const testManifestContent = `
{
    "main.js": {
        "file": "assets/main.js",
        "css": ["assets/main.css"]
    },
    "another.js": {
        "file": "assets/another.js",
        "css": ["assets/another.css"]
    }
}
`

func setupTestManifest(t *testing.T) {
	t.Helper()

	manifestFile := filepath.Join("static", ".vite", "manifest.json")
	err := os.MkdirAll(filepath.Dir(manifestFile), 0755)
	if err != nil {
		t.Fatalf("Failed to create test manifest directory: %v", err)
	}

	err = os.WriteFile(manifestFile, []byte(testManifestContent), 0644)
	if err != nil {
		t.Fatalf("Failed to write test manifest file: %v", err)
	}
}

func teardownTestManifest(t *testing.T) {
	t.Helper()
	manifestFile := filepath.Join("static", ".vite", "manifest.json")
	err := os.RemoveAll(filepath.Dir(manifestFile))
	if err != nil {
		t.Fatalf("Failed to remove test manifest directory: %v", err)
	}
}

func TestLoadManifest(t *testing.T) {
	setupTestManifest(t)
	defer teardownTestManifest(t)

	LoadManifest()

	expected := map[string]ManifestEntry{
		"main.js": {
			File: "assets/main.js",
			CSS:  []string{"assets/main.css"},
		},
		"another.js": {
			File: "assets/another.js",
			CSS:  []string{"assets/another.css"},
		},
	}

	for key, entry := range expected {
		if manifestEntry, exists := manifest[key]; exists {
			if manifestEntry.File != entry.File || !reflect.DeepEqual(manifestEntry.CSS, entry.CSS) {
				t.Errorf("For key %s, expected %v, got %v", key, entry, manifestEntry)
			}
		} else {
			t.Errorf("Expected key %s to exist", key)
		}
	}
}

func TestAssetPath(t *testing.T) {
	setupTestManifest(t)
	defer teardownTestManifest(t)

	LoadManifest()

	tests := []struct {
		asset    string
		expected string
	}{
		{"main.js", "assets/main.js"},
		{"another.js", "assets/another.js"},
		{"nonexistent.js", ""},
	}

	for _, tt := range tests {
		t.Run(tt.asset, func(t *testing.T) {
			got := AssetPath(tt.asset)
			if got != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, got)
			}
		})
	}
}

func TestCSSPaths(t *testing.T) {
	setupTestManifest(t)
	defer teardownTestManifest(t)

	LoadManifest()

	tests := []struct {
		asset    string
		expected []string
	}{
		{"main.js", []string{"assets/main.css"}},
		{"another.js", []string{"assets/another.css"}},
		{"nonexistent.js", nil},
	}

	for _, tt := range tests {
		t.Run(tt.asset, func(t *testing.T) {
			got := CSSPaths(tt.asset)
			if !equal(got, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, got)
			}
		})
	}
}

// Helper function to compare slices
func equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
