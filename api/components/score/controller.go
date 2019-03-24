package score

import (
	"database/sql"
	"fmt"
)

func GetPlayerScore(db *sql.DB, matchKey int) ([]PlayerScore, error) {
	statement := fmt.Sprintf(
		`SELECT
			player.id,
			player.name,
			player.fk_player_team as team,
			player_score.quantity as score
		FROM
			player_score
		INNER JOIN
			player
		ON player_score.fk_score_player = player.id
		WHERE player_score.fk_score_match = %d
		ORDER BY score DESC`, matchKey)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	playersScore := []PlayerScore{}

	for rows.Next() {
		var p PlayerScore
		if err := rows.Scan(&p.Player.ID, &p.Player.Name, &p.Player.Team, &p.Score); err != nil {
			return nil, err
		}
		playersScore = append(playersScore, p)
	}

	return playersScore, nil
}

func GetScore(db *sql.DB) ([]PlayerScore, error) {
	statement := fmt.Sprintf(
		`SELECT 
			player.name, 
			team.name as team_name,
			player.id as player_id,
			COALESCE(sum(player_score.quantity), 0) as score
		FROM 
			player
		INNER JOIN team
			ON player.fk_player_team = team.name
		LEFT JOIN player_score
			ON player_score.fk_score_player = player.id
		GROUP BY name, team_name, player_id
		ORDER BY score DESC`)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	playersScore := []PlayerScore{}

	for rows.Next() {
		var p PlayerScore
		if err := rows.Scan(&p.Player.Name, &p.Player.Team, &p.Player.ID, &p.Score); err != nil {
			return nil, err
		}
		playersScore = append(playersScore, p)
	}

	return playersScore, nil
}
