package main

import (
	"io/ioutil"
	"os/exec"
	"regexp"
)

func fetchEmail() []string {
	return detectAccounts(detectEmails(offlineimap()))
}

func offlineimap() []byte {
	imapCmd := exec.Command("offlineimap")
	imapOut, _ := imapCmd.StderrPipe()
	imapCmd.Start()
	imapBytes, _ := ioutil.ReadAll(imapOut)
	imapCmd.Wait()
	return imapBytes
}

func detectEmails(imapBytes []byte) []byte {
	grepCmd := exec.Command("grep", "Remote:INBOX ->")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write(imapBytes)
	grepIn.Close()
	grepBytes, _ := ioutil.ReadAll(grepOut)
	grepCmd.Wait()
	return grepBytes
}

func detectAccounts(grepBytes []byte) []string {
	accounts := make([]string, 0)
	s := string(grepBytes)
	r, _ := regexp.Compile(`([^\ ]*)-Remote:INBOX\ ->`)
	for _, matches := range r.FindAllStringSubmatch(s, -1) {
		account := matches[len(matches)-1]
		if !include(account, accounts) {
			accounts = append(accounts, account)
		}
	}
	return accounts
}

func include(element string, collection []string) bool {
	for _, e := range collection {
		if e == element {
			return true
		}
	}
	return false
}
