package minion

import "testing"

func TestNewDefaults(t *testing.T) {
	m := New("")
	if m.Name != "Bob" {
		t.Errorf("default name = %q, want Bob", m.Name)
	}
	if m.Stats.Energy != 60 || m.Stats.Hunger != 30 {
		t.Errorf("unexpected starting stats: %+v", m.Stats)
	}
}

func TestFeedClamps(t *testing.T) {
	m := &Minion{Stats: Stats{Energy: 95, Hunger: 10}}
	m.Feed()
	if m.Stats.Energy != 100 {
		t.Errorf("energy should clamp at 100, got %d", m.Stats.Energy)
	}
	if m.Stats.Hunger != 0 {
		t.Errorf("hunger should clamp at 0, got %d", m.Stats.Hunger)
	}
}

func TestTickActivityFeedsEnergy(t *testing.T) {
	m := &Minion{Stats: Stats{Energy: 50, Hunger: 20}}
	m.Tick(10) // -5 decay + 10 activity = +5
	if m.Stats.Energy != 55 {
		t.Errorf("energy = %d, want 55", m.Stats.Energy)
	}
	if m.Stats.Hunger != 30 {
		t.Errorf("hunger = %d, want 30", m.Stats.Hunger)
	}
}

func TestAgeIncrementsEachTick(t *testing.T) {
	m := New("")
	if m.Age != 0 {
		t.Errorf("new minion age = %d, want 0", m.Age)
	}
	activities := []int{0, 5, -3, 100, 0}
	for i, a := range activities {
		m.Tick(a)
		if m.Age != i+1 {
			t.Errorf("after %d ticks age = %d, want %d", i+1, m.Age, i+1)
		}
	}
}

func TestAgeStringPluralization(t *testing.T) {
	cases := []struct {
		age  int
		want string
	}{
		{0, "0 ticks"},
		{1, "1 tick"}, // singular
		{2, "2 ticks"},
		{42, "42 ticks"},
	}
	for _, c := range cases {
		m := &Minion{Age: c.age}
		if got := m.AgeString(); got != c.want {
			t.Errorf("AgeString(age=%d) = %q, want %q", c.age, got, c.want)
		}
	}
}

func TestIsHungry(t *testing.T) {
	cases := []struct {
		hunger int
		want   bool
	}{
		{69, false}, // just below threshold
		{70, true},  // boundary
		{71, true},  // above threshold
	}
	for _, c := range cases {
		m := &Minion{Stats: Stats{Hunger: c.hunger}}
		if got := m.IsHungry(); got != c.want {
			t.Errorf("IsHungry(hunger=%d) = %v, want %v", c.hunger, got, c.want)
		}
	}
}

func TestMood(t *testing.T) {
	cases := []struct {
		energy, hunger int
		want           string
	}{
		{80, 20, "happy"},
		{50, 50, "content"},
		{15, 20, "grumpy"},
		{50, 85, "sad"},
	}
	for _, c := range cases {
		m := &Minion{Stats: Stats{Energy: c.energy, Hunger: c.hunger}}
		if got := m.Mood(); got != c.want {
			t.Errorf("mood(e=%d,h=%d) = %q, want %q", c.energy, c.hunger, got, c.want)
		}
	}
}
