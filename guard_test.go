package main

import (
	"testing"
	"time"
)

func TestMinutesAgo(t *testing.T) {
	g := guard{limit: 5, latest: time.Now().UTC().Add(-6 * time.Minute)}
	expected := 6
	actual := g.minutesAgo()
	if expected != actual {
		t.Errorf("Expected: %d, got: %v", expected, actual)
	}
}

func TestMayFetch(t *testing.T) {
	g := guard{limit: 5, latest: time.Now().UTC().Add(-5 * time.Minute)}
	if !g.mayFetch() {
		t.Errorf("It's been %d minutes. Should be able to fetch", g.minutesAgo())
	}
}

