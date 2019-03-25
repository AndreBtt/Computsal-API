package card

import (
	"database/sql"
	"fmt"
)

func GetCards(db *sql.DB) ([]PlayerCard, error) {
	statement := fmt.Sprintf(
		`SELECT 
			player.name, 
			team.name as team_name,
			player.id as player_id,
			COALESCE(sum(card.yellow),0) as yellow,
			COALESCE(sum(card.red),0) as red
		FROM 
			player
		INNER JOIN team
			ON player.fk_player_team = team.name
		LEFT JOIN card
			ON card.fk_card_player = player.id
		GROUP BY name, team_name, player_id
		ORDER BY 
			yellow DESC,
			red DESC`)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	playersCard := []PlayerCard{}

	for rows.Next() {
		var p PlayerCard
		if err := rows.Scan(&p.Player.Name, &p.Player.Team, &p.Player.ID, &p.Yellow, &p.Red); err != nil {
			return nil, err
		}
		playersCard = append(playersCard, p)
	}

	return playersCard, nil
}

func GetPlayerCard(db *sql.DB, matchKey int) ([]PlayerCard, error) {
	statement := fmt.Sprintf(
		`SELECT
			player.id,
			player.name,
			player.fk_player_team as team,
			card.yellow,
			card.red
		FROM
			card
		INNER JOIN
			player
		ON card.fk_card_player = player.id
		WHERE card.fk_card_match = %d
		ORDER BY 
			yellow DESC,
			red DESC`, matchKey)

	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	playersCard := []PlayerCard{}

	for rows.Next() {
		var p PlayerCard
		if err := rows.Scan(&p.Player.ID, &p.Player.Name, &p.Player.Team, &p.Yellow, &p.Red); err != nil {
			return nil, err
		}
		playersCard = append(playersCard, p)
	}

	return playersCard, nil
}
