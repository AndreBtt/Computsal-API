package player

import (
	"database/sql"
	"fmt"
)

func GetPlayers(db *sql.DB) ([]Player, error) {
	statement := fmt.Sprintf("SELECT id, nome FROM player")
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	players := []Player{}

	for rows.Next() {
		var u Player
		if err := rows.Scan(&u.ID, &u.Name); err != nil {
			return nil, err
		}
		players = append(players, u)
	}

	return players, nil
}
