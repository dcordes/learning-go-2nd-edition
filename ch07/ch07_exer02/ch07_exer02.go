package main

import (
	"fmt"
	"sort"
)

type Team struct {
	team_name    string
	player_names []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(
	// using a pointer receiver here because we are modifying the elements of the receiver
	//	---> but note that we get indentical results if we make this a value receiver
	//	--->	I did not have time to figure out why
	first_team_name string,
	first_team_score int,
	second_team_name string,
	second_team_score int,
) {
	// increment the wins for a team based on which team did best, ignoring ties
	if first_team_score > second_team_score {
		l.Wins[first_team_name]++
	} else {
		l.Wins[second_team_name]++
	}
}

func (l League) Ranking() []string {
	// using a value receiver here because we can get away with it,
	//	since we are not modifying the elements of the receiver, as a test

	// returns a slice of the team names in order of wins
	// ~~~~~~~~~~~~~

	// note that the author of Learning Go has a much more
	//	elegant way of sorting here that uses "sort.Slice" with
	//	a function comparison, such that he does not have to create
	//	all these slices

	scores := map[int][]string{}

	for _, team := range l.Teams {
		wins, ok := l.Wins[team.team_name]
		if !ok {
			scores[0] = append(scores[0], team.team_name)
		} else {
			scores[wins] = append(scores[wins], team.team_name)
		}
	}

	//fmt.Println("scores map!: ", scores)

	// turn the scores into a slice of ints
	var raw_scores []int
	for score := range scores {
		raw_scores = append(raw_scores, score)
	}
	// sort the scores
	//	https://pkg.go.dev/sort#Reverse
	//fmt.Println("before sort", raw_scores)
	sort.Sort(sort.Reverse(sort.IntSlice(raw_scores)))
	//fmt.Println("after sort", raw_scores)

	// populate team names
	//	couldn't use RANGE here since it seemed to start at the end the slice
	var names []string
	for i := 0; i < len(raw_scores); i++ {
		teams := scores[raw_scores[i]]
		//fmt.Println("for raw score", raw_scores[i], " found these teams: ", teams)
		for _, team := range teams {
			names = append(names, team)
		}
	}

	return names
}

func main() {
	aaa := Team{
		team_name:    "The AAA Team",
		player_names: []string{"A", "AA", "AAA"},
	}
	bbb := Team{
		team_name:    "The BBB Team",
		player_names: []string{"B", "BB", "BBB"},
	}
	ccc := Team{
		team_name:    "The CCC Team",
		player_names: []string{"C", "CC", "CCC"},
	}
	ddd := Team{
		team_name:    "The DDD Team",
		player_names: []string{"D", "DD", "DDD"},
	}

	// declare and "construct" the struct
	the_league := League{
		Teams: []Team{aaa, bbb, ccc, ddd},
		Wins:  make(map[string]int),
	}

	// populate with matches
	the_league.MatchResult(
		aaa.team_name,
		5,
		bbb.team_name,
		9,
	)
	the_league.MatchResult(
		ccc.team_name,
		12,
		ddd.team_name,
		3,
	)
	the_league.MatchResult(
		bbb.team_name,
		7,
		ccc.team_name,
		4,
	)
	the_league.MatchResult(
		aaa.team_name,
		7,
		ddd.team_name,
		7,
	)
	the_league.MatchResult(
		aaa.team_name,
		7,
		ddd.team_name,
		7,
	)

	fmt.Println("Teams in order of greatest to least wins: ", the_league.Ranking())

}
