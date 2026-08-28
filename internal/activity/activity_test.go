package activity

import (
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestCountRegularFiles(t *testing.T) {
	cases := []struct {
		name  string
		files int
	}{
		{"empty", 0},
		{"one", 1},
		{"several", 5},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			dir := t.TempDir()
			for i := 0; i < c.files; i++ {
				path := filepath.Join(dir, "session-"+strconv.Itoa(i))
				if err := os.WriteFile(path, []byte("session"), 0o644); err != nil {
					t.Fatalf("writing file: %v", err)
				}
			}
			// A subdirectory must not be counted.
			if err := os.Mkdir(filepath.Join(dir, "subdir"), 0o755); err != nil {
				t.Fatalf("mkdir: %v", err)
			}
			if got := Count(dir); got != c.files {
				t.Errorf("Count() = %d, want %d", got, c.files)
			}
		})
	}
}

func TestCountMissingDir(t *testing.T) {
	missing := filepath.Join(t.TempDir(), "does-not-exist")
	if got := Count(missing); got != 0 {
		t.Errorf("Count(missing) = %d, want 0", got)
	}
}
