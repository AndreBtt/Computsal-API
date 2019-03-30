package team

import (
	"database/sql"
	"fmt"

	"github.com/AndreBtt/Computsal/api/components/match"

	"github.com/AndreBtt/Computsal/api/components/captain"
	"github.com/AndreBtt/Computsal/api/components/player"
)

func CreateTeam(db *sql.DB, t TeamCreate) error {
	if len(t.Players) == 0 {
		return fmt.Errorf("Insert at least captain player")
	}

	// create team table
	statement := fmt.Sprintf(`
		INSERT INTO 
			team(name, photo, group_number) 
		VALUES('%s', '%s', %d)`, t.Name, t.PhotoURL, -1)
	if _, err := db.Exec(statement); err != nil {
		return err
	}

	// create team's captain
	if err := createTeamCaptain(db, t); err != nil {
		return err
	}

	// move captain out of players
	t.Players = t.Players[1:]

	// create players
	return createTeamPlayers(db, t)
}

func createTeamPlayers(db *sql.DB, t TeamCreate) error {
	teamName := t.Name
	var players []player.PlayerCreate
	for _, elem := range t.Players {
		p := player.PlayerCreate{Name: elem.Name, Team: teamName}
		players = append(players, p)
	}
	return player.CreatePlayers(db, players)
}

func createTeamCaptain(db *sql.DB, t TeamCreate) error {
	// first need to create captain player and get the ID
	captainName := t.Players[0].Name
	captainPlayer := player.PlayerTable{Team: t.Name, Name: captainName}
	if err := captainPlayer.CreatePlayer(db); err != nil {
		return err
	}

	// create captain
	cap := captain.CaptainCreate{PlayerID: captainPlayer.ID, Team: captainPlayer.Team, Email: t.CaptainEmail}
	return cap.CreateCaptain(db)
}

func (t *TeamUpdate) UpdateTeam(db *sql.DB) error {
	statement := fmt.Sprintf(`
		UPDATE 
			team
		SET
			name = '%s',
			photo = '%s'
		WHERE
			id = %d`, t.Name, t.PhotoURL, t.ID)
	_, err := db.Exec(statement)
	return err
}

func (t *TeamTable) DeleteTeam(db *sql.DB) error {
	statement := fmt.Sprintf(`
		DELETE FROM
			team 
		WHERE 
			id = %d`, t.ID)
	_, err := db.Exec(statement)
	return err
}

func GetTeams(db *sql.DB) ([]TeamTable, error) {
	statement := fmt.Sprintf("SELECT id, name, photo, group_number FROM team")
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := []TeamTable{}

	for rows.Next() {
		var t TeamTable
		if err := rows.Scan(&t.ID, &t.Name, &t.PhotoURL, &t.Group); err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}

	return teams, nil
}

func (teamDetails *Team) GetTeam(db *sql.DB) error {
	statement := fmt.Sprintf(`
		SELECT
			team.name,
			team.photo,
			team.group_number,
			team.id,
			COALESCE(next_match.fk_next_team1, "flag_no_team") AS team1,
			COALESCE(next_match.fk_next_team2, "flag_no_team") AS team2,
			COALESCE(time.time, "00:00:00") AS time,
			player.name AS captain
		FROM team
			LEFT JOIN
				next_match
					ON 	next_match.fk_next_team1 = team.name OR
						next_match.fk_next_team2 = team.name
			LEFT JOIN
				time
					ON time.id = next_match.time
			INNER JOIN
				captain
					ON captain.fk_captain_team = team.name
			INNER JOIN
				player
					ON player.id = captain.fk_captain_player
		WHERE 
			team.name = '%s'`, teamDetails.Name)

	var team1, team2 string
	if err := db.QueryRow(statement).Scan(&teamDetails.Name, &teamDetails.PhotoURL,
		&teamDetails.Group, &teamDetails.ID, &team1, &team2,
		&teamDetails.NextMatch.Time, &teamDetails.CaptainName); err != nil {
		return err
	}

	if team1 == teamDetails.Name {
		teamDetails.NextMatch.Name = team2
	}
	if team2 == teamDetails.Name {
		teamDetails.NextMatch.Name = team1
	}

	// get team's previous matches to calculate win lose draw and goals
	var err error
	if teamDetails.PreviousMatches, err = match.GetTeamPreviousMatches(db, teamDetails.Name); err != nil {
		return err
	}
	for _, elem := range teamDetails.PreviousMatches {
		if elem.Team1 == teamDetails.Name {
			if elem.Score1 > elem.Score2 {
				teamDetails.Win++
			} else if elem.Score1 < elem.Score2 {
				teamDetails.Lose++
			} else {
				teamDetails.Draw++
			}
			teamDetails.GoalsPro += elem.Score1
			teamDetails.GoalsAgainst += elem.Score2
		} else {
			if elem.Score2 > elem.Score1 {
				teamDetails.Win++
			} else if elem.Score2 < elem.Score1 {
				teamDetails.Lose++
			} else {
				teamDetails.Draw++
			}
			teamDetails.GoalsPro += elem.Score2
			teamDetails.GoalsAgainst += elem.Score1
		}
	}

	if teamDetails.Players, err = player.GetPlayersScore(db, teamDetails.Name); err != nil {
		return err
	}

	return nil
}
