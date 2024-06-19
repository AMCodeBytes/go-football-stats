package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
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

// type Teams struct {
// 	Teams []Team
// }

// type Team struct {
// 	Name string
// }

func main() {
	file, err := os.Open("data.json")

	if err != nil {
		panic("Failed to open file")
	}

	byteValue, _ := io.ReadAll(file)

	var matches Matches
	var teams []string

	json.Unmarshal(byteValue, &matches)

	for i := 0; i < len(matches.Matches); i++ {
		// fmt.Printf("| Match Date: %s | Season Game: %d | Match Day: %d | Full Time Result: %s | Home Team: %s | Full Time Home Goals: %d | Away Team: %s | Full Time Away Goals: %d |\n", matches.Matches[i].MatchDate, matches.Matches[i].SeasonGame, matches.Matches[i].MatchDay, matches.Matches[i].FullTimeResult, matches.Matches[i].HomeTeam, matches.Matches[i].FullTimeHomeGoals, matches.Matches[i].AwayTeam, matches.Matches[i].FullTimeAwayGoals)
		teams = append(teams, matches.Matches[i].HomeTeam, matches.Matches[i].AwayTeam)
	}

	fmt.Println(removeDuplicates(teams))

	// fmt.Print(teams)

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
