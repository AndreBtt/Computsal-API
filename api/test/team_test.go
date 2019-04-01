package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"testing"

	"github.com/AndreBtt/Computsal/api/components/match"
	"github.com/AndreBtt/Computsal/api/components/player"
	"github.com/AndreBtt/Computsal/api/components/team"

	_ "github.com/go-sql-driver/mysql"
)

func TestMain(m *testing.M) {
	code := m.Run()

	os.Exit(code)
}

func TestTeamAPI(t *testing.T) {

	/* -------------  CREATE TEAM -------------------- */

	url := "http://localhost:8080"

	payload := []byte(`
	{
        "name":     "flag_test_team",
        "photo":    "www.flag_test_url.com.br",
        "players": 
            [
                {
                    "name" : "flag_test_captain"		
                },
				{
					"name" : "flag_test_player2"		
				}
            ],
        "captain_email" : "flag_test_email@email.com.br"
	}`)

	req, _ := http.NewRequest("POST", url+"/teams", bytes.NewBuffer(payload))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	checkResponseCode(t, http.StatusCreated, resp)
	resp.Body.Close()

	/* -------------  RETRIVE TEAM ------------------- */

	tExpected := team.Team{
		Name:         "flag_test_team",
		PhotoURL:     "www.flag_test_url.com.br",
		CaptainName:  "flag_test_captain",
		Group:        -1,
		Win:          0,
		Lose:         0,
		Draw:         0,
		GoalsPro:     0,
		GoalsAgainst: 0,
		NextMatch: team.TeamNextMatch{
			Name: "",
			Time: "00:00:00",
		},
		PreviousMatches: []match.PreviousMatchList{},
	}

	players := []player.PlayerTeamScore{}
	players = append(players, player.PlayerTeamScore{
		Name:       "flag_test_captain",
		Score:      0,
		YellowCard: 0,
		RedCard:    0,
	})
	players = append(players, player.PlayerTeamScore{
		Name:       "flag_test_player2",
		Score:      0,
		YellowCard: 0,
		RedCard:    0,
	})

	tExpected.Players = players

	req, _ = http.NewRequest("GET", url+"/teams/flag_test_team", bytes.NewBuffer(payload))
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	var tRetrive team.Team

	if err := json.NewDecoder(resp.Body).Decode(&tRetrive); err != nil {
		t.Fatal(err)
	}
	checkResponseCode(t, http.StatusOK, resp)
	resp.Body.Close()

	checkTeam(tExpected, tRetrive, t)

	/* -------------  UPDATE TEAM -------------------- */

	payload = []byte(
		fmt.Sprintf(`
		{  
			"id":        %d,
			"name":      "flag_test_team_update",
			"photo":     "www.flag_test_url_update.com.br"
		}`, tRetrive.ID))

	req, _ = http.NewRequest("PUT", url+"/teams", bytes.NewBuffer(payload))

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	checkResponseCode(t, http.StatusCreated, resp)
	resp.Body.Close()

	/* -------------  RETRIVE UPDATED TEAM ----------- */

	req, _ = http.NewRequest("GET", url+"/teams/flag_test_team_update", bytes.NewBuffer(payload))
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	var tRetriveUpdate team.Team
	if err := json.NewDecoder(resp.Body).Decode(&tRetriveUpdate); err != nil {
		t.Fatal(err)
	}

	checkResponseCode(t, http.StatusOK, resp)
	resp.Body.Close()

	tExpected.Name = "flag_test_team_update"
	tExpected.PhotoURL = "www.flag_test_url_update.com.br"
	checkTeam(tExpected, tRetriveUpdate, t)

	/* -------------  RETRIVE TEAMS ----------- */

	req, _ = http.NewRequest("GET", url+"/teams", bytes.NewBuffer(payload))
	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	var teams []team.TeamTable
	if err := json.NewDecoder(resp.Body).Decode(&teams); err != nil {
		t.Fatal(err)
	}

	checkResponseCode(t, http.StatusOK, resp)
	resp.Body.Close()

	found := false
	for _, elem := range teams {
		if elem.Name == tExpected.Name {
			found = true
		}
	}

	if !found {
		t.Errorf("get teams request did not find created team")
	}

	/* -------------  DELETE TEAM -------------------- */

	req, _ = http.NewRequest("DELETE", url+"/teams/"+strconv.Itoa(tRetriveUpdate.ID), nil)

	client = &http.Client{}
	resp, err = client.Do(req)
	if err != nil {
		panic(err)
	}
	checkResponseCode(t, http.StatusOK, resp)
	resp.Body.Close()
}

func checkTeam(tExpected team.Team, tGot team.Team, t *testing.T) {
	if tGot.Name != tExpected.Name {
		t.Errorf("Name is incorrect, got: %v, want: %v", tGot.Name, tExpected.Name)
	}

	if tGot.PhotoURL != tExpected.PhotoURL {
		t.Errorf("PhotoURL is incorrect, got: %v, want: %v", tGot.PhotoURL, tExpected.PhotoURL)
	}

	if tGot.CaptainName != tExpected.CaptainName {
		t.Errorf("Captain name is incorrect, got: %v, want: %v", tGot.CaptainName, tExpected.CaptainName)
	}

	if tGot.Group != tExpected.Group {
		t.Errorf("Group number is incorrect, got: %d, want: %d", tGot.Group, tExpected.Group)
	}

	if tGot.Win != tExpected.Win {
		t.Errorf("Win is incorrect, got: %d, want: %d", tGot.Win, tExpected.Win)
	}

	if tGot.Lose != tExpected.Lose {
		t.Errorf("Lose is incorrect, got: %d, want: %d", tGot.Lose, tExpected.Lose)
	}

	if tGot.Draw != tExpected.Draw {
		t.Errorf("Draw is incorrect, got: %d, want: %d", tGot.Draw, tExpected.Draw)
	}

	if tGot.GoalsPro != tExpected.GoalsPro {
		t.Errorf("GoalsPro is incorrect, got: %d, want: %d", tGot.GoalsPro, tExpected.GoalsPro)
	}

	if tGot.GoalsAgainst != tExpected.GoalsAgainst {
		t.Errorf("GoalsAgainst is incorrect, got: %d, want: %d", tGot.GoalsAgainst, tExpected.GoalsAgainst)
	}

	if tGot.NextMatch.Name != tExpected.NextMatch.Name {
		t.Errorf("NextMatch name is incorrect, got: %v, want: %v", tGot.NextMatch.Name, tExpected.NextMatch.Name)
	}

	if tGot.NextMatch.Time != tExpected.NextMatch.Time {
		t.Errorf("NextMatch time is incorrect, got: %v, want: %v", tGot.NextMatch.Time, tExpected.NextMatch.Time)
	}

	comparePlayers(tExpected.Players, tGot.Players, t)

}

func comparePlayers(tExpected []player.PlayerTeamScore, tGot []player.PlayerTeamScore, t *testing.T) {
	for _, elem := range tGot {
		if elem.Score != 0 {
			t.Errorf("Player %v score should be 0 instead got %d", elem.Name, elem.Score)
		}
		if elem.YellowCard != 0 {
			t.Errorf("Player %v YellowCard should be 0 instead got %d", elem.Name, elem.YellowCard)
		}
		if elem.RedCard != 0 {
			t.Errorf("Player %v RedCard should be 0 instead got %d", elem.Name, elem.RedCard)
		}
	}

	if !((tGot[0].Name == tExpected[0].Name && tGot[1].Name == tExpected[1].Name) ||
		(tGot[1].Name == tExpected[0].Name && tGot[0].Name == tExpected[1].Name)) {
		t.Errorf("Players name are not equal")
	}
}

func checkResponseCode(t *testing.T, expected int, actual *http.Response) {
	if expected != actual.StatusCode {
		fmt.Println(actual.Status)
		t.Errorf("Expected response code %d. Got %d\n", expected, actual.StatusCode)
	}
}
