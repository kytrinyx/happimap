package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

const longFormat = "2006-01-02 15:04:05 +0000 UTC"

func limit() (minutes int) {
	return configLimit(configFile())
}

func latest() time.Time {
	return configLatest(latestFile())
}

func configFile() (name string) {
	return fmt.Sprintf("%s/%s/%s", os.Getenv("HOME"), ".happimap", "config")
}

func latestFile() (name string) {
	return fmt.Sprintf("%s/%s/%s", os.Getenv("HOME"), ".happimap", "latest")
}

func logFile() (name string) {
	return fmt.Sprintf("%s/%s/%s", os.Getenv("HOME"), ".happimap", "log")
}

func configLimit(file string) (minutes int) {
	value, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}
	s := strings.Trim(string(value), "\n")
	minutes, err = strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func configLatest(file string) time.Time {
	value, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err)
	}

	s := strings.Trim(string(value), "\n")
	t, err := time.Parse(longFormat, s)
	if err != nil {
		fmt.Println(err)
	}
	return t
}

func updateLatest(file string, timestamp time.Time) {
	ioutil.WriteFile(file, []byte(timestamp.Format(longFormat)), 0644)
}

func logLatest(file string, timestamp time.Time, status string) {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	ts := timestamp.Format(longFormat)
	s := fmt.Sprintf("%s\t%s\n", ts, status)
	defer f.Close()
	_, err = f.Write([]byte(s))
	return
}

func status(allowed, forced bool) string {
	switch {
	default:
		return "WAIT"
	case allowed:
		return "OK"
	case forced:
		return "FORCE"
	}
}
