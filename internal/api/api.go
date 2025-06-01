package api

import (
	"fmt"
	"strconv"
	"strings"
)

func (l *leagueInfo) formatApiCall() (string, error) {
	yearAsInt, err := strconv.Atoi(l.Year)
	if err != nil {
		return "", fmt.Errorf("[Error] - Failed to convert year from string to int: %w\n", err)
	}

	var apiReqStr string
	if yearAsInt < 2018 {
		apiReqStr = strings.Replace(l.V2Api, "{year}", l.Year, 1)
		apiReqStr = strings.Replace(apiReqStr, "{leagueid}", l.LeagueID, 1)
	} else {
		apiReqStr = strings.Replace(l.V3Api, "{year}", l.Year, 1)
		apiReqStr = strings.Replace(apiReqStr, "{leagueid}", l.LeagueID, 1)
	}

	return apiReqStr, nil
}

type leagueInfo struct {
	LeagueID string `json:"league_id"`
	Year     string `json:"year"`
	V2Api    string `json:"v2Api"`
	V3Api    string `json:"v3Api"`
	SWID     string `json:"swid"`
	S2       string `json:"espn_s2"`
}

func main() {
	v3_years := []string{
		"2023",
		"2022",
		"2021",
		"2020",
		"2020",
		"2019",
		"2018",
	}

	v2_years := []string{
		"2017",
		"2016",
		"2015",
		"2014",
		"2013",
		"2012",
	}

}
