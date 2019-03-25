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

func GetScores(db *sql.DB) ([]PlayerScore, error) {
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

func DeleteScore(db *sql.DB, scoreID int) error {
	statement := fmt.Sprintf(`
		DELETE FROM
			player_score
		WHERE 
			id = %d
		`, scoreID)

	_, err := db.Exec(statement)
	return err
}

// func (ps *PlayerScoreTable) UpdateScore(db *sql.DB) error {
// 	// if the player does not score in a match he should be deleted
// 	if ps.Quantity == 0 {
// 		err := ps.DeleteScore(db)
// 		return err
// 	}

// 	statement := fmt.Sprintf(`
// 		UPDATE
// 			player_score
// 		SET
// 			quantity = %d
// 		WHERE
// 			id = %d
// 		`, ps.Quantity, ps.ID)

// 	_, err := db.Exec(statement)
// 	return err
// }

func (ps *PlayerScoreTable) CreateScore(db *sql.DB) error {
	// if the player does not score in a match he should not be added
	if ps.Quantity == 0 {
		return nil
	}

	statement := fmt.Sprintf(`
		INSERT INTO 
			player_score
		VALUES
			(%d, %d, %d)
		`, ps.PlayerID, ps.MatchID, ps.Quantity)

	if _, err := db.Exec(statement); err != nil {
		return err
	}

	err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&ps.ID)
	return err
}
