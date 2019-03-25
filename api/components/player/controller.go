package player

import (
	"database/sql"
	"fmt"
)

func (p *PlayerTable) CreatePlayer(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO player(name, fk_player_team) VALUES('%s', '%s')", p.Name, p.Team)
	if _, err := db.Exec(statement); err != nil {
		return err
	}

	if err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&p.ID); err != nil {
		return err
	}

	return nil
}

func (p *PlayerTable) DeletePlayer(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM player WHERE id=%d", p.ID)
	_, err := db.Exec(statement)
	return err
}

func (p *PlayerTable) UpdatePlayer(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE player SET name='%s' WHERE id=%d", p.Name, p.ID)
	_, err := db.Exec(statement)
	return err
}

func (p *Player) GetPlayer(db *sql.DB) error {
	statement := fmt.Sprintf(
		`SELECT
			player.id,
			player.name,
			team.name AS team_name,
			team.photo AS team_photo,
			(SELECT COUNT(player.id) 
				FROM player 
					INNER JOIN captain
						ON fk_captain_player = player.id
				WHERE player.id = %d) AS captain,
			COALESCE(SUM(quantity), 0) AS score,
			COALESCE(SUM(yellow), 0) AS yellow,
			COALESCE(SUM(red), 0) AS red
		FROM player
			INNER JOIN player_match
				ON player_match.fk_score_player = player.id
			INNER JOIN team
				ON player.fk_player_team = team.name
		WHERE
			player.id = %d
		GROUP BY player.id, player.name, team.name, captain`, p.ID, p.ID)
	if err := db.QueryRow(statement).Scan(&p.ID, &p.Name, &p.Team, &p.TeamPhotoURL, &p.Captain, &p.Score, &p.YellowCard, &p.RedCard); err != nil {
		return err
	}

	return nil
}

func GetPlayers(db *sql.DB) ([]PlayerTable, error) {
	statement := fmt.Sprintf("SELECT * FROM player")
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	players := []PlayerTable{}

	for rows.Next() {
		var p PlayerTable
		if err := rows.Scan(&p.ID, &p.Name, &p.Team); err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	return players, nil
}
