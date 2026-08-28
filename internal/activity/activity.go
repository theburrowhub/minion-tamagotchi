// Package activity measures recent agent activity from the filesystem so
// `minion tick` can feed the minion with real work as its energy source.
package activity

import "os"

// Count returns the number of regular files directly under dir, used as a
// stand-in for recent Claude Code session files. Subdirectories, symlinks and
// other non-regular entries are ignored, and it does not recurse. A missing or
// unreadable dir yields 0 rather than an error.
func Count(dir string) int {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return 0
	}
	n := 0
	for _, entry := range entries {
		if entry.Type().IsRegular() {
			n++
		}
	}
	return n
}
