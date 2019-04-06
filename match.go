package winwin

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type Match struct {
	Home   *Team
	Away   *Team
	result *MatchResult
}

func (m *Match) String() string {
	var scores string

	teams := fmt.Sprintf("%v vs %v", m.Home.Name, m.Away.Name)
	if m.result != nil {
		scores = fmt.Sprintf("%v-%v", m.result.HomeGoals, m.result.AwayGoals)
	}

	return strings.Join([]string{teams, scores}, ":")
}

type MatchResult struct {
	HomeGoals int
	AwayGoals int
}

func SimulateMatch(m *Match) {
	hg := simulateGoals(*m.Home, *m.Away)
	ag := simulateGoals(*m.Away, *m.Home)

	result := &MatchResult{HomeGoals: hg, AwayGoals: ag}

	m.result = result
}

func simulateGoals(attacker Team, defenser Team) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	diff := float64(attacker.Strength - defenser.Strength)
	max := math.Max(diff, 2.0)

	goals := r.Intn(int(max))

	return goals
}
