package team

import (
	"database/sql"
	"fmt"

	"github.com/AndreBtt/Computsal/api/components/captain"
	"github.com/AndreBtt/Computsal/api/components/player"
)

func CreateTeam(db *sql.DB, t TeamCreate) error {
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
