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
