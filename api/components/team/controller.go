package team

import (
	"database/sql"
	"fmt"

	player "github.com/AndreBtt/Computsal/api/components/player"
)

func (t *TeamTable) CreateTeam(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO team(name, photo, group_number) VALUES('%s', '%s', %d)", t.Name, t.PhotoURL, t.Group)
	if _, err := db.Exec(statement); err != nil {
		return err
	}
	if err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&t.ID); err != nil {
		return err
	}
	return nil
}

func (t *TeamTable) UpdateTeam(db *sql.DB) error {
	statement := fmt.Sprintf(`
		UPDATE 
			team
		SET
			name = '%s',
			photo = '%s',
			group_number = %d
		WHERE
			id = %d`, t.Name, t.PhotoURL, t.Group, t.ID)
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

func (t *TeamTable) GetTeam(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id, photo, group_number FROM team WHERE name = '%s'", t.Name)
	return db.QueryRow(statement).Scan(&t.ID, &t.PhotoURL, &t.Group)
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

func GetPlayers(db *sql.DB, teamName string) ([]player.PlayerTable, error) {
	statement := fmt.Sprintf("SELECT * FROM player where fk_player_team = '%s'", teamName)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	players := []player.PlayerTable{}

	for rows.Next() {
		var p player.PlayerTable
		if err := rows.Scan(&p.ID, &p.Name, &p.Team); err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	return players, nil
}
