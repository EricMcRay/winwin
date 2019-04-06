package winwin

type fixture []*standing

type standing struct {
	Team  Team
	Games int
	Wins  int
	Ties  int
	Loses int
	Avg   int
	Score int
}

func NewFixture(l League) fixture {
	f := make(map[string]*standing)
	for _, t := range l.Teams {
		f[t.Name] = &standing{Team: *t}
	}

	for _, m := range l.Matches {
		hs := f[m.Home.Name]
		as := f[m.Away.Name]

		hs.Games++
		as.Games++

		hs.Avg += m.result.HomeGoals - m.result.AwayGoals
		as.Avg += m.result.AwayGoals - m.result.HomeGoals

		switch true {
		case m.result.HomeGoals == m.result.AwayGoals:
			hs.Ties++
			as.Ties++

			hs.Score++
			as.Score++
		case m.result.HomeGoals > m.result.AwayGoals:
			hs.Wins++
			as.Loses++

			hs.Score += 3
		case m.result.HomeGoals < m.result.AwayGoals:
			hs.Loses++
			as.Wins++

			as.Score += 3
		}
	}

	fixture := make(fixture, 0)
	for _, s := range f {
		fixture = append(fixture, s)
	}

	return fixture
}
