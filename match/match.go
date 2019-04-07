package match

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
	"winwin/team"
)

type Match struct {
	Home   *team.Team
	Away   *team.Team
	Result *Result
}

func (m *Match) String() string {
	var scores string

	teams := fmt.Sprintf("%v vs %v", m.Home.Name, m.Away.Name)
	if m.Result != nil {
		scores = fmt.Sprintf("%v-%v", m.Result.HomeGoals, m.Result.AwayGoals)
	}

	return strings.Join([]string{teams, scores}, ":")
}

type Result struct {
	HomeGoals int
	AwayGoals int
}

func Simulate(m *Match) {
	hg := simulateGoals(*m.Home, *m.Away)
	ag := simulateGoals(*m.Away, *m.Home)

	m.Result = &Result{HomeGoals: hg, AwayGoals: ag}
}

func simulateGoals(attacker team.Team, defenser team.Team) int {
	diff := float64(attacker.Strength - defenser.Strength)
	max := math.Max(diff, 2.0)

	return randomGoals(int(max))
}

func randomGoals(max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return r.Intn(max)
}
