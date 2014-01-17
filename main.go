package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "happimap"
	app.Usage = "Control the email reflex"
	app.Action = func(c *cli.Context) {
		g := latestGuard()
		if g.mayFetch() {
			fmt.Printf("Checking email.\nLast checked %d minutes ago.\n", g.minutesAgo())
			accounts := fetchEmail()
			now := time.Now().UTC()
			updateLatest(latestFile(), now)
			logLatest(logFile(), now)

			if len(accounts) == 0 {
				fmt.Println("No new emails.")
			} else {
				fmt.Printf("You have mail in %v.\n", accounts)
			}

		} else {
			min := "minutes"
			if g.next() == 1 {
				min = "minute"
			}

			fmt.Printf("Wait another %d %s.\n", g.next(), min)
		}
	}

	app.Run(os.Args)
}

func latestGuard() guard {
	return guard{limit: limit(), latest: latest()}
}
