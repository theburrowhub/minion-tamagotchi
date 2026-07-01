// Command minion is the tamagotchi CLI. It renders a pet that reacts to your
// AI-agent activity. See README.md for the specs.
package main

import (
	"fmt"
	"os"

	"github.com/theburrowhub/minion-tamagotchi/internal/minion"
)

func main() {
	cmd := "status"
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	m := minion.New("Bob")

	switch cmd {
	case "status":
		fmt.Println(m.Render())
	case "feed":
		m.Feed()
		fmt.Println("Yum!")
		fmt.Println(m.Render())
	case "tick":
		// Activity wiring lands in internal/activity; 0 for now.
		m.Tick(0)
		fmt.Println(m.Render())
	case "age":
		fmt.Printf("age: %d ticks\n", m.Age)
	case "help", "-h", "--help":
		usage(os.Stdout)
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n\n", cmd)
		usage(os.Stderr)
		os.Exit(2)
	}
}

func usage(w *os.File) {
	fmt.Fprint(w, `minion — a CLI tamagotchi that reacts to your AI-agent activity

Usage:
  minion status   Show the minion and its stats
  minion feed     Feed the minion
  minion tick     Advance time one step
  minion age      Show how many ticks the minion has lived
`)
}
