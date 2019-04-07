package league

import (
	"winwin/match"
	"winwin/team"
)

type League struct {
	Teams   []*team.Team
	Matches []*match.Match
}

type standings map[*team.Team]*Standing

func New(tt []*team.Team) *League {
	m := generateMatches(tt)

	return &League{Teams: tt, Matches: m}
}

func (l League) Standings() standings {
	ss := make(standings)

	for _, t := range l.Teams {
		ss[t] = &Standing{}
	}

	for _, m := range l.Matches {
		hs, as := matchStandings(*m)

		ss[m.Home].Apply(*hs)
		ss[m.Away].Apply(*as)
	}

	return ss
}

func generateMatches(teams []*team.Team) []*match.Match {
	matches := make([]*match.Match, 0)
	for _, home := range teams {
		for _, away := range teams {
			if home == away {
				continue
			}

			matches = append(matches, &match.Match{Home: home, Away: away})
		}
	}

	return matches
}
