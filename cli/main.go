package main

import (
	"os"
	"sort"
	"strconv"
	"winwin/league"
	"winwin/match"
	"winwin/team"

	"github.com/olekukonko/tablewriter"
)

type row struct {
	t *team.Team
	*league.Standing
}

func main() {
	teams := getTeams()
	league := league.New(teams)

	for _, m := range league.Matches {
		match.Simulate(m)
	}

	renderStandings(league)
}

func renderStandings(l *league.League) {
	rows := leagueRows(l)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Team", "OM", "G", "B", "M", "A", "P"})

	sort.SliceStable(rows, func(i, j int) bool {
		return rows[i].Score > rows[j].Score
	})

	for i, s := range rows {
		r := tableRow(i+1, s)
		table.Append(r)
	}

	table.Render() // Send output
}

func leagueRows(l *league.League) []*row {
	rows := make([]*row, 0)
	standings := l.Standings()

	for t, s := range standings {
		rows = append(rows, &row{t, s})
	}

	return rows
}

func tableRow(order int, s *row) []string {
	return []string{
		strconv.Itoa(order),
		s.t.Name,
		strconv.Itoa(s.Games),
		strconv.Itoa(s.Wins),
		strconv.Itoa(s.Ties),
		strconv.Itoa(s.Loses),
		strconv.Itoa(s.Avg),
		strconv.Itoa(s.Score),
	}
}
