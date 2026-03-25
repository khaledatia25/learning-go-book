package main

import (
	"fmt"
	"slices"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if _, ok := l.Teams[team1]; !ok {
		return
	}
	if _, ok := l.Teams[team2]; !ok {
		return
	}
	var winner string
	if score1 == score2 {
		return
	}
	if score1 > score2 {
		winner = team1
	} else {
		winner = team2
	}
	l.Wins[winner]++
}

func (l *League) Ranking() []string {
	ranking := make([]string, 0, len(l.Teams))
	for team := range l.Teams {
		ranking = append(ranking, team)
	}
	slices.SortFunc(ranking, func(i, j string) int {
		return l.Wins[j] - l.Wins[i]
	})
	return ranking
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 50, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)
	results := l.Ranking()
	fmt.Println(results)
	for l, r := range l.Wins {
		fmt.Println(l, r)
	}
}
