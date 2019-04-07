package score

import (
	"io"
	"os"
	"sort"
	"strconv"
	"winwin/league"
	"winwin/team"

	"github.com/olekukonko/tablewriter"
)

// Table is reprents league scores
type Table []*Row

type Row struct {
	Team *team.Team
	*league.Standing
}

func NewTable(l league.League) Table {
	tt := make(Table, 0)
	standings := l.Standings()

	for team, s := range standings {
		tt = append(tt, &Row{team, s})
	}

	return tt
}

func (t Table) OrderByScore() {
	sort.SliceStable(t, func(i, j int) bool {
		return t[i].Score > t[j].Score
	})
}

func (t Table) Render(w io.Writer) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Team", "OM", "G", "B", "M", "A", "P"})

	for i, s := range t {
		table.Append(tableRow(i+1, s))
	}

	table.Render() // Send output
}

func tableRow(order int, s *Row) []string {
	return []string{
		strconv.Itoa(order),
		s.Team.Name,
		strconv.Itoa(s.Games),
		strconv.Itoa(s.Wins),
		strconv.Itoa(s.Ties),
		strconv.Itoa(s.Loses),
		strconv.Itoa(s.Avg),
		strconv.Itoa(s.Score),
	}
}
