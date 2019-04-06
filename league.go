package winwin

type League struct {
	Teams   []*Team
	Matches []*Match
}

type Standing struct {
	Matches uint
	Wins    uint
	Ties    uint
	Loses   uint
	Average uint
	Score   uint
}

func NewLeague(tt []*Team) *League {
	m := generateMatches(tt)

	return &League{Teams: tt, Matches: m}
}

func generateMatches(teams []*Team) []*Match {
	matches := make([]*Match, 0)
	for _, home := range teams {
		for _, away := range teams {
			if home == away {
				continue
			}

			matches = append(matches, &Match{Home: home, Away: away})
		}
	}

	return matches
}
