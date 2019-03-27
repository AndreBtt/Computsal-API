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
			player_match.quantity as score,
			player_match.yellow as yellow,
			player_match.red as red
		FROM
			player_match
		INNER JOIN
			player
		ON player_match.fk_score_player = player.id
		WHERE player_match.fk_score_match = %d
		ORDER BY 
			score DESC,
			yellow ASC,
			red ASC
		`, matchKey)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	playersScore := []PlayerScore{}

	for rows.Next() {
		var p PlayerScore
		if err := rows.Scan(&p.Player.ID, &p.Player.Name, &p.Player.Team, &p.Score, &p.Yellow, &p.Red); err != nil {
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
			COALESCE(sum(player_match.quantity), 0) as score,
			COALESCE(sum(player_match.yellow), 0) as yellow,
			COALESCE(sum(player_match.red), 0) as red
		FROM 
			player
		INNER JOIN team
			ON player.fk_player_team = team.name
		LEFT JOIN player_match
			ON player_match.fk_score_player = player.id
		GROUP BY name, team_name, player_id
		ORDER BY 
			score DESC,
			red ASC,
			yellow ASC`)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	playersScore := []PlayerScore{}

	for rows.Next() {
		var p PlayerScore
		if err := rows.Scan(&p.Player.Name, &p.Player.Team, &p.Player.ID, &p.Score, &p.Yellow, &p.Red); err != nil {
			return nil, err
		}
		playersScore = append(playersScore, p)
	}

	return playersScore, nil
}

func DeleteScores(db *sql.DB, matchID int, scores []PlayerIDScore) error {
	statement := fmt.Sprintf(`
		DELETE FROM
			player_match
		WHERE `)
	for _, elem := range scores {
		s := fmt.Sprintf("(fk_score_player = %d AND fk_score_match = %d) OR ", elem.ID, matchID)
		statement += s
	}

	statement = statement[:len(statement)-3]
	_, err := db.Exec(statement)
	return err
}

func UpdateScores(db *sql.DB, matchID int, scores []PlayerIDScore) error {
	deleteScores := []PlayerIDScore{}
	updateScores := []PlayerIDScore{}

	for _, elem := range scores {
		// if the player does not have data in a match he should be deleted
		if elem.Score == 0 && elem.Yellow == 0 && elem.Red == 0 {
			deleteScores = append(deleteScores, elem)
		} else {
			updateScores = append(updateScores, elem)
		}
	}

	if len(deleteScores) > 0 {
		if err := DeleteScores(db, matchID, deleteScores); err != nil {
			return err
		}
	}

	if len(updateScores) > 0 {
		tx, err := db.Begin()
		if err != nil {
			return err
		}

		for _, elem := range updateScores {
			statement := fmt.Sprintf(`
				UPDATE
					player_match
				SET
					quantity = %d,
					yellow = %d,
					red = %d
				WHERE
					fk_score_player = %d and
					fk_score_match = %d
				`, elem.Score, elem.Yellow, elem.Red, elem.ID, matchID)

			if _, err := tx.Exec(statement); err != nil {
				if rollbackErr := tx.Rollback(); rollbackErr != nil {
					return rollbackErr
				}
				return err
			}
		}

		if err := tx.Commit(); err != nil {
			return err
		}
	}

	return nil
}

func (ps *PlayerScoreTable) CreateScore(db *sql.DB) error {
	// if the player does not have any data in a match he should not be added
	if ps.Quantity == 0 && ps.Red == 0 && ps.Yellow == 0 {
		return nil
	}

	statement := fmt.Sprintf(`
		INSERT INTO 
			player_match(fk_score_player, fk_score_match, quantity, yellow, red)
		VALUES(%d, %d, %d, %d, %d)
		`, ps.PlayerID, ps.MatchID, ps.Quantity, ps.Yellow, ps.Red)

	if _, err := db.Exec(statement); err != nil {
		return err
	}

	err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&ps.ID)
	return err
}
