package main

import (
	"encoding/json"
	"fmt"
	"github.com/statusquonjc46/fantasy-football-app/internal/api/api.go"
	"io"
	"net/http"
	"os"
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

	league := &api.leagueInfo{
		LeagueID: leagueID,
		Year:     year,
		V2Api:    "https://lm-api-reads.fantasy.espn.com/apis/v3/games/ffl/leagueHistory/{leagueid}?seasonId={year}",
		V3Api:    "https://lm-api-reads.fantasy.espn.com/apis/v3/games/ffl/seasons/{year}/segments/0/leagues/{leagueid}",
		SWID:     swid,
		S2:       espnS2,
	}

	url, err := league.formatApiCall()
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{}

	// Create a new HTTP GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Add a cookie to the request
	cookieSWID := &http.Cookie{
		Name:  "swid",
		Value: league.SWID,
	}
	req.AddCookie(cookieSWID)

	cookieS2 := &http.Cookie{
		Name:  "espn_s2",
		Value: league.S2,
	}
	req.AddCookie(cookieS2)

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		err := fmt.Errorf("[Error] - Failed to make request to ESPN api: %w\n", err)
		fmt.Println(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err := fmt.Errorf("[Error] - Request failed with status code: %d\n", resp.StatusCode)
		fmt.Println(err)
		os.Exit(1)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		err := fmt.Errorf("[Error] - Failed to read response body: %w", err)
		fmt.Println(err)
	}

	var printJson string
	err = json.Unmarshal(body, &printJson)
	if err != nil {
		err := fmt.Errorf("[Error] - Error unmarshalling JSON: %w", err)
		fmt.Println(err)
	}

	fmt.Println(printJson)
}
