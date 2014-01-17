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

func TestNext(t *testing.T) {
	g := guard{limit: 5, latest: time.Now().UTC().Add(-1 * time.Minute)}
	expected := 4
	actual := g.next()
	if expected != actual {
		t.Errorf("Expected to wait 4 more minutes, but must wait %d more minutes.", actual)
	}
}
