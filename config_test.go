package main

import (
	"os"
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
	file := "./.testUpdateLatest"
	t1 := time.Date(2014, 1, 2, 3, 45, 0, 0, time.UTC)
	updateLatest(file, t1)
	defer os.Remove(file)

	t2 := configLatest(file)
	if t1 != t2 {
		t.Errorf("Expected: %v, got: %v")
	}
}
