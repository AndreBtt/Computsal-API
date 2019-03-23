package team

import (
	"database/sql"
	"fmt"

	player "github.com/AndreBtt/Computsal/api/components/player"
)

func (t *Team) CreateTeam(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO team(name, photo, group_number) VALUES('%s', '%s', %d)", t.Name, t.PhotoURL, t.Group)
	if _, err := db.Exec(statement); err != nil {
		return err
	}
	if err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&t.ID); err != nil {
		return err
	}
	return nil
}

func (t *Team) UpdateTeam(db *sql.DB, key string) error {
	statement := fmt.Sprintf("UPDATE team SET name='%s', photo='%s', group_number=%d WHERE name='%s'", t.Name, t.PhotoURL, t.Group, key)
	_, err := db.Exec(statement)
	return err
}

func (t *Team) DeleteTeam(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM team WHERE name='%s'", t.Name)
	_, err := db.Exec(statement)
	return err
}

func (t *Team) GetTeam(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT id, photo, group_number FROM team WHERE name = '%s'", t.Name)
	return db.QueryRow(statement).Scan(&t.ID, &t.PhotoURL, &t.Group)
}

func GetTeams(db *sql.DB) ([]Team, error) {
	statement := fmt.Sprintf("SELECT * FROM team")
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := []Team{}

	for rows.Next() {
		var t Team
		if err := rows.Scan(&t.ID, &t.Name, &t.PhotoURL, &t.Group); err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}

	return teams, nil
}

func GetPlayers(db *sql.DB, teamName string) ([]player.Player, error) {
	statement := fmt.Sprintf("SELECT * FROM player where fk_player_team = '%s'", teamName)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	players := []player.Player{}

	for rows.Next() {
		var p player.Player
		if err := rows.Scan(&p.ID, &p.Name, &p.Team); err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	return players, nil
}
