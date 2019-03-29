package player

import (
	"database/sql"
	"fmt"
)

func CreatePlayers(db *sql.DB, players []PlayerCreate) error {
	statement := fmt.Sprintf("INSERT INTO player(name, fk_player_team) VALUES")
	for _, elem := range players {
		values := fmt.Sprintf(" ('%s', '%s'),", elem.Name, elem.Team)
		statement += values
	}

	statement = statement[:len(statement)-1]

	_, err := db.Exec(statement)
	return err
}

func DeletePlayers(db *sql.DB, players []PlayerID) error {
	if len(players) == 0 {
		return nil
	}

	statement := fmt.Sprintf("DELETE FROM player WHERE")
	for _, elem := range players {
		value := fmt.Sprintf(" id = %d OR", elem.ID)
		statement += value
	}
	statement = statement[:len(statement)-2]
	_, err := db.Exec(statement)
	return err
}

func UpdatePlayers(db *sql.DB, players []PlayerUpdate) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, elem := range players {
		statement := fmt.Sprintf("UPDATE player SET name='%s' WHERE id=%d", elem.Name, elem.ID)
		if _, err := tx.Exec(statement); err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return rollbackErr
			}
			return err
		}
	}

	return tx.Commit()
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

func (p *PlayerTable) CreatePlayer(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO player(name, fk_player_team) VALUES ('%s', '%s')", p.Name, p.Team)
	if _, err := db.Exec(statement); err != nil {
		return err
	}
	err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&p.ID)
	return err
}

func GetPlayersScore(db *sql.DB, teamName string) ([]PlayerTeamScore, error) {
	statement := fmt.Sprintf(`
		SELECT
			player.id,
			player.name,
			coalesce(sum(player_match.quantity),0) AS score,
			coalesce(sum(player_match.yellow), 0) AS yellow,
			coalesce(sum(player_match.red),0) AS red
		FROM 
			player
		LEFT JOIN
			player_match
				ON player_match.fk_score_player = player.id
		WHERE player.fk_player_team = '%s'
		GROUP BY player.id, player.name
		ORDER BY score DESC, yellow DESC, red DESC`, teamName)

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	players := []PlayerTeamScore{}

	for rows.Next() {
		var p PlayerTeamScore
		if err := rows.Scan(&p.ID, &p.Name, &p.Score, &p.YellowCard, &p.RedCard); err != nil {
			return nil, err
		}
		players = append(players, p)
	}

	return players, nil
}
