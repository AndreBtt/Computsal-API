package player

import (
	"database/sql"
	"fmt"
)

func (p *Player) CreatePlayer(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO player(name, fk_player_team) VALUES('%s', '%s')", p.Name, p.Team)
	if _, err := db.Exec(statement); err != nil {
		return err
	}

	if err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&p.ID); err != nil {
		return err
	}

	return nil
}

func (p *Player) DeletePlayer(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM player WHERE id=%d", p.ID)
	_, err := db.Exec(statement)
	return err
}

func (p *Player) UpdatePlayer(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE player SET name='%s' WHERE id=%d", p.Name, p.ID)
	_, err := db.Exec(statement)
	return err
}

func (p *Player) GetPlayer(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT name, fk_player_team FROM player WHERE id='%d'", p.ID)
	return db.QueryRow(statement).Scan(&p.Name, &p.Team)
}

func GetPlayers(db *sql.DB) ([]Player, error) {
	statement := fmt.Sprintf("SELECT * FROM player")
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	players := []Player{}

	for rows.Next() {
		var p Player
		if err := rows.Scan(&p.ID, &p.Name, &p.Team); err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	return players, nil
}
