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
			player.name,
			team.name as team_name,
			team.photo as team_photo,
				(SELECT COUNT(player.id) 
				FROM player 
					INNER JOIN captain
						ON fk_captain_player = player.id
				WHERE player.id = %d) as captain,
			COALESCE((SELECT SUM(quantity) 
				FROM player_score 
					LEFT JOIN player
						ON player.id = fk_score_player),0) as score,
			COALESCE((SELECT SUM(yellow) 
				FROM card 
					LEFT JOIN player 
						ON player.id = fk_card_player),0) as yellow,
			COALESCE((SELECT SUM(red) 
				FROM card 
					LEFT JOIN player 
						ON player.id = fk_card_player),0) as red
			FROM player
				INNER JOIN team
					ON player.fk_player_team = team.name
			WHERE 
				player.id = %d`, p.ID, p.ID)
	if err := db.QueryRow(statement).Scan(&p.Name, &p.Team, &p.TeamPhotoURL, &p.Captain, &p.Score, &p.YellowCard, &p.RedCard); err != nil {
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
