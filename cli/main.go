package main

import (
	"os"
	"sort"
	"strconv"
	"winwin"

	"github.com/olekukonko/tablewriter"
)

func main() {
	teams := getTeams()

	league := winwin.NewLeague(teams)

	for _, m := range league.Matches {
		winwin.SimulateMatch(m)
	}

	fixture := winwin.NewFixture(*league)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"", "Team", "OM", "G", "B", "M", "A", "P"})

	sort.SliceStable(fixture, func(i, j int) bool {
		return fixture[i].Score > fixture[j].Score
	})

	for i, s := range fixture {
		table.Append([]string{
			strconv.Itoa(i + 1),
			s.Team.Name,
			strconv.Itoa(s.Games),
			strconv.Itoa(s.Wins),
			strconv.Itoa(s.Ties),
			strconv.Itoa(s.Loses),
			strconv.Itoa(s.Avg),
			strconv.Itoa(s.Score),
		})
	}
	table.Render() // Send output
}

func getTeams() []*winwin.Team {
	tt := make([]*winwin.Team, 0)

	tt = append(tt, &winwin.Team{Name: "Galatasaray", Strength: 12})
	tt = append(tt, &winwin.Team{Name: "Fenerbahçe", Strength: 10})
	tt = append(tt, &winwin.Team{Name: "Başakşehir", Strength: 11})
	tt = append(tt, &winwin.Team{Name: "Beşiktaş", Strength: 10})
	tt = append(tt, &winwin.Team{Name: "Trabzonspor", Strength: 8})
	tt = append(tt, &winwin.Team{Name: "Göztepe", Strength: 7})
	tt = append(tt, &winwin.Team{Name: "Sivasspor", Strength: 7})
	tt = append(tt, &winwin.Team{Name: "Kasımpaşa", Strength: 7})
	tt = append(tt, &winwin.Team{Name: "Kayserispor", Strength: 6})
	tt = append(tt, &winwin.Team{Name: "Yeni Malatya", Strength: 6})
	tt = append(tt, &winwin.Team{Name: "Akhisar", Strength: 6})
	tt = append(tt, &winwin.Team{Name: "Alanyaspor", Strength: 6})
	tt = append(tt, &winwin.Team{Name: "Bursaspor", Strength: 6})
	tt = append(tt, &winwin.Team{Name: "Antalyaspor", Strength: 5})
	tt = append(tt, &winwin.Team{Name: "Konyaspor", Strength: 5})
	tt = append(tt, &winwin.Team{Name: "Osmanlıspor", Strength: 4})
	tt = append(tt, &winwin.Team{Name: "Gençlerbirliği", Strength: 4})
	tt = append(tt, &winwin.Team{Name: "Karabükspor", Strength: 2})

	return tt
}
