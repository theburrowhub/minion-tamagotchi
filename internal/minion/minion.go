// Package minion is the tamagotchi model: stats, mood, and rendering. It is
// pure (no filesystem/clock access) so it is fully unit-testable.
package minion

import "fmt"

// Stats are the minion's vital signs, each clamped to 0–100.
type Stats struct {
	Energy int `json:"energy"`
	Hunger int `json:"hunger"`
}

// Minion is the pet.
type Minion struct {
	Name  string `json:"name"`
	Stats Stats  `json:"stats"`
	// Age is the number of ticks the minion has lived through. Unlike the
	// vital signs in Stats, it is an unbounded lifetime counter (never clamped).
	Age int `json:"age"`
}

// New returns a fresh minion with balanced stats.
func New(name string) *Minion {
	if name == "" {
		name = "Bob"
	}
	return &Minion{Name: name, Stats: Stats{Energy: 60, Hunger: 30}}
}

func clamp(v int) int {
	if v < 0 {
		return 0
	}
	if v > 100 {
		return 100
	}
	return v
}

// Feed raises energy and lowers hunger.
func (m *Minion) Feed() {
	m.Stats.Energy = clamp(m.Stats.Energy + 15)
	m.Stats.Hunger = clamp(m.Stats.Hunger - 25)
}

// Tick advances time one step: agent activity feeds mood (energy), while hunger
// creeps up and energy decays. activity is a non-negative count of recent
// agent events.
func (m *Minion) Tick(activity int) {
	if activity < 0 {
		activity = 0
	}
	m.Stats.Energy = clamp(m.Stats.Energy - 5 + activity)
	m.Stats.Hunger = clamp(m.Stats.Hunger + 10)
	m.Age++
}

// Mood is derived from energy and hunger.
func (m *Minion) Mood() string {
	switch {
	case m.Stats.Hunger >= 80:
		return "sad"
	case m.Stats.Energy >= 70 && m.Stats.Hunger <= 40:
		return "happy"
	case m.Stats.Energy <= 25:
		return "grumpy"
	default:
		return "content"
	}
}

var faces = map[string]string{
	"happy":   "(^_^)",
	"content": "(-_-)",
	"grumpy":  "(>_<)",
	"sad":     "(T_T)",
}

// Render returns an ASCII view of the minion plus its stats.
func (m *Minion) Render() string {
	mood := m.Mood()
	return fmt.Sprintf("%s  %s\n  energy: %d/100\n  hunger: %d/100\n  mood:   %s",
		m.Name, faces[mood], m.Stats.Energy, m.Stats.Hunger, mood)
}
