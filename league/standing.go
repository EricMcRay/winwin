package league

import (
	"winwin/match"
)

type Standing struct {
	Games int
	Wins  int
	Ties  int
	Loses int
	Avg   int
	Score int
}

func (s *Standing) Apply(a Standing) {
	s.Games += a.Games
	s.Wins += a.Wins
	s.Ties += a.Ties
	s.Loses += a.Loses
	s.Avg += a.Avg
	s.Score += a.Score
}

func matchStandings(m match.Match) (*Standing, *Standing) {
	hs := &Standing{}
	as := &Standing{}

	hs.Games++
	as.Games++

	hs.Avg += m.Result.HomeGoals - m.Result.AwayGoals
	as.Avg += m.Result.AwayGoals - m.Result.HomeGoals

	switch true {
	case m.Result.HomeGoals == m.Result.AwayGoals:
		hs.Ties++
		as.Ties++

		hs.Score++
		as.Score++
	case m.Result.HomeGoals > m.Result.AwayGoals:
		hs.Wins++
		as.Loses++

		hs.Score += 3
	case m.Result.HomeGoals < m.Result.AwayGoals:
		hs.Loses++
		as.Wins++

		as.Score += 3
	}

	return hs, as
}
