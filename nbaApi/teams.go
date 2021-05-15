package nbaApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Team struct {
	League struct {
		Standard []struct {
			Isnbafranchise bool   `json:"isNBAFranchise"`
			Isallstar      bool   `json:"isAllStar"`
			City           string `json:"city"`
			Altcityname    string `json:"altCityName"`
			Fullname       string `json:"fullName"`
			Tricode        string `json:"tricode"`
			Teamid         string `json:"teamId"`
			Nickname       string `json:"nickname"`
			Urlname        string `json:"urlName"`
			Teamshortname  string `json:"teamShortName"`
			Confname       string `json:"confName"`
			Divname        string `json:"divName"`
		} `json:"standard"`
		Africa     []interface{} `json:"africa"`
		Sacramento []interface{} `json:"sacramento"`
		Vegas      []interface{} `json:"vegas"`
		Utah       []interface{} `json:"utah"`
	} `json:"league"`
}

type nbaTeamsResponse struct {
	Teams []Team `json:"standard"`
}

func getAllTheTeams() ([]Team, error) {
	res, err := http.Get(fmt.Sprintf("%s/standard", baseURL))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response nbaTeamsResponse
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return response.Teams, err
}
