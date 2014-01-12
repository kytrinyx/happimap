package main

import "time"

type guard struct {
	limit  int
	latest time.Time
}

func (g guard) mayFetch() (ok bool) {
	return g.minutesAgo() > g.limit
}

func (g guard) minutesAgo() int {
	return int(time.Since(g.latest).Minutes())
}

func (g guard) next() (minutes int) {
	if g.mayFetch() {
		return 0
	}
	return g.limit - g.minutesAgo()
}
