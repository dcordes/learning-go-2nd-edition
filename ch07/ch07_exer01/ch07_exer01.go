package main

type Team struct {
	team_name    string
	player_names []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func main() {

}
