package main

import (
	"io/ioutil"
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

func TestLogLatest(t *testing.T) {
	file := "./.testLogLatest"
	t1 := time.Date(2014, 1, 2, 3, 45, 0, 0, time.UTC)
	t2 := time.Date(2014, 2, 3, 4, 56, 0, 0, time.UTC)
	logLatest(file, t1, true)
	logLatest(file, t2, false)
	defer os.Remove(file)

	actual, err := ioutil.ReadFile(file)
	if err != nil {
		t.Errorf("Unable to read %s", file)
	}

	expected := "2014-01-02 03:45:00 +0000 UTC\tOK\n2014-02-03 04:56:00 +0000 UTC\tWAIT\n"
	if expected != string(actual) {
		t.Errorf("Expected: %s\n Got: %s\n", expected, actual)
	}
}
