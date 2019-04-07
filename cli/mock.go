package main

import "winwin/team"

func getTeams() []*team.Team {
	tt := make([]*team.Team, 0)

	tt = append(tt, &team.Team{Name: "Galatasaray", Strength: 12})
	tt = append(tt, &team.Team{Name: "Fenerbahçe", Strength: 10})
	tt = append(tt, &team.Team{Name: "Başakşehir", Strength: 11})
	tt = append(tt, &team.Team{Name: "Beşiktaş", Strength: 10})
	tt = append(tt, &team.Team{Name: "Trabzonspor", Strength: 8})
	tt = append(tt, &team.Team{Name: "Göztepe", Strength: 7})
	tt = append(tt, &team.Team{Name: "Sivasspor", Strength: 7})
	tt = append(tt, &team.Team{Name: "Kasımpaşa", Strength: 7})
	tt = append(tt, &team.Team{Name: "Kayserispor", Strength: 6})
	tt = append(tt, &team.Team{Name: "Yeni Malatya", Strength: 6})
	tt = append(tt, &team.Team{Name: "Akhisar", Strength: 6})
	tt = append(tt, &team.Team{Name: "Alanyaspor", Strength: 6})
	tt = append(tt, &team.Team{Name: "Bursaspor", Strength: 6})
	tt = append(tt, &team.Team{Name: "Antalyaspor", Strength: 5})
	tt = append(tt, &team.Team{Name: "Konyaspor", Strength: 5})
	tt = append(tt, &team.Team{Name: "Osmanlıspor", Strength: 4})
	tt = append(tt, &team.Team{Name: "Gençlerbirliği", Strength: 4})
	tt = append(tt, &team.Team{Name: "Karabükspor", Strength: 2})

	return tt
}
