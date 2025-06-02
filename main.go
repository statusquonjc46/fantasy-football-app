package main

import (
	"fmt"
	"github.com/statusquonjc46/fantasy-football-app/internal/api"
)

type returnError struct {
	Error string `json:"error"`
}

func main() {
	var year string
	var leagueID string
	var swid string
	var espnS2 string

	fmt.Println("> Beginning of fantasy football app")
	fmt.Println("> Please enter your league id: ")
	fmt.Println("> You can find this information if you go to your leagues home page and checking the url for https://fantasy.espn.com/football/team?leagueid={copy this number here}")
	fmt.Scanln(&leagueID)
	fmt.Println("Please enter the year you would like to query: ")
	fmt.Scanln(&year)
	fmt.Printf("League ID: %s | Season: %s\n", leagueID, year)
	fmt.Println("Enter the SWID Cookie value: Ex.{ABCD1234-ABCD-1234-EF56-1234567890AB}")
	fmt.Scanln(&swid)
	fmt.Println("Enter the espn_s2 Cookie value:")
	fmt.Scanln(&espnS2)

	callData := api.ProcessUserInput(leagueID, year, swid, espnS2)
	api.MakeCall(callData)

}
