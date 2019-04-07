package main

import (
	"os"
	"winwin/league"
	"winwin/match"
	"winwin/score"
)

func main() {
	tt := getTeams()
	l := league.New(tt)

	for _, m := range l.Matches {
		match.Simulate(m)
	}

	scores := score.NewTable(*l)
	scores.OrderByScore()

	scores.Render(os.Stdout)
}
