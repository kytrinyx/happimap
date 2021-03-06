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
	app.Flags = []cli.Flag{
		cli.BoolFlag{"force, f", "Override wait injunction"},
	}
	app.Action = func(c *cli.Context) {
		g := latestGuard()

		now := time.Now().UTC()
		status := status(g.mayFetch(), c.Bool("force"))
		logLatest(logFile(), now, status)

		if g.mayFetch() || c.Bool("force") {
			fmt.Printf("Checking email.\nLast checked %d minutes ago.\n", g.minutesAgo())
			accounts := fetchEmail()
			updateLatest(latestFile(), now)

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
