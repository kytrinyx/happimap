package main

import (
	"testing"
	"time"
)

func TestConfigLimit(t *testing.T) {
	expected := 5
	actual := configLimit("./.testConfig")
	if expected != actual {
		t.Errorf("Expected: %d, got: %v", expected, actual)
	}
}

func TestConfigLatest(t *testing.T) {
	expected := time.Date(2014, 1, 2, 3, 45, 0, 0, time.UTC)
	actual := configLatest("./.testLatest")
	if expected != actual {
		t.Errorf("Expected: %v, got: %v", expected, actual)
	}
}

func TestUpdateLatest(t *testing.T) {
	updateLatest("./.testUpdateLatest", time.Now().UTC())
}
