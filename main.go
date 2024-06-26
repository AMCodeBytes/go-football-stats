package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
)

type Matches struct {
	Matches []Match `json:"matches"`
}

type Match struct {
	MatchDate         string `json:"matchdate"`
	SeasonGame        int    `json:"seasongame"`
	MatchDay          int    `json:"matchday"`
	FullTimeResult    string `json:"fulltimeresult"`
	HomeTeam          string `json:"hometeam"`
	FullTimeHomeGoals int    `json:"fulltimehomegoals"`
	AwayTeam          string `json:"awayteam"`
	FullTimeAwayGoals int    `json:"fulltimeawaygoals"`
}

type UnfilteredTeams struct {
	Teams []string
}

type Teams struct {
	Teams []Team
}

type Team struct {
	Name       string
	Wins       int
	Draws      int
	Loses      int
	TotalGoals int
	Points     int
}

func main() {
	file, err := os.Open("data.json")

	if err != nil {
		panic("Failed to open file")
	}

	byteValue, _ := io.ReadAll(file)

	var matches Matches
	var unfilteredTeams UnfilteredTeams
	var teams Teams

	json.Unmarshal(byteValue, &matches)

	for i := 0; i < len(matches.Matches); i++ {
		unfilteredTeams.Teams = append(unfilteredTeams.Teams, matches.Matches[i].HomeTeam, matches.Matches[i].AwayTeam)
	}

	unfilteredTeams.Teams = removeDuplicates(unfilteredTeams.Teams)

	for i := 0; i < len(unfilteredTeams.Teams); i++ {
		var team Team

		team.Name = unfilteredTeams.Teams[i]

		for i := 0; i < len(matches.Matches); i++ {
			if matches.Matches[i].HomeTeam != team.Name && matches.Matches[i].AwayTeam != team.Name {
				continue
			}

			if matches.Matches[i].HomeTeam == team.Name {
				switch matches.Matches[i].FullTimeResult {
				case "H":
					team.Wins++
					team.Points += 3
				case "A":
					team.Loses++
				case "D":
					team.Draws++
					team.Points++
				}

				team.TotalGoals += matches.Matches[i].FullTimeHomeGoals
			} else {
				// If the above isn't true, we know the team is either home or away thus must be away
				switch matches.Matches[i].FullTimeResult {
				case "H":
					team.Loses++
				case "A":
					team.Wins++
					team.Points += 3
				case "D":
					team.Draws++
					team.Points++
				}

				team.TotalGoals += matches.Matches[i].FullTimeAwayGoals
			}
		}

		teams.Teams = append(teams.Teams, team)
	}

	_ = sortTeams(teams)

	for i := 0; i < len(teams.Teams); i++ {
		fmt.Printf("Name: %s, Points: %d, Wins: %d, Loses: %d, Draws: %d, Goals: %d,\n", teams.Teams[i].Name, teams.Teams[i].Points, teams.Teams[i].Wins, teams.Teams[i].Loses, teams.Teams[i].Draws, teams.Teams[i].TotalGoals)
	}

	defer file.Close()
}

func removeDuplicates(strString []string) []string {
	allKeys := make(map[string]bool)
	list := []string{}

	for _, item := range strString {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}

	return list
}

func sortTeams(unorderedTeams Teams) Teams {
	sort.Slice(unorderedTeams.Teams[:], func(i, j int) bool {
		return unorderedTeams.Teams[i].Points > unorderedTeams.Teams[j].Points
	})

	return unorderedTeams
}
