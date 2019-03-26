package match

import (
	"database/sql"
	"fmt"

	score "github.com/AndreBtt/Computsal/api/components/score"
)

func GetPreviousMatches(db *sql.DB) ([]PreviousMatchList, error) {
	statement := fmt.Sprintf(`
		SELECT
			previous_match.match_type,
			previous_match.phase,
			previous_match.fk_match_team1 AS team1,
			previous_match.fk_match_team2 AS team2,
			COALESCE(player.fk_player_team, "flag_no_score") AS team,
			previous_match.id AS match_id,
			COALESCE(sum(player_match.quantity), 0) AS team_score
		FROM previous_match
			LEFT JOIN player_match
				ON previous_match.id = player_match.fk_score_match
			LEFT JOIN player
				ON player.id = player_match.fk_score_player
		GROUP BY match_type, phase, team1, team2, team, match_id
		ORDER BY match_id
		`)

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	matchQuery := []PreviousMatchesQuery{}

	for rows.Next() {
		var p PreviousMatchesQuery
		if err := rows.Scan(&p.Type, &p.Phase, &p.Team1, &p.Team2, &p.Team, &p.ID, &p.Score); err != nil {
			return nil, err
		}
		matchQuery = append(matchQuery, p)
	}

	matchList := []PreviousMatchList{}

	for i := 0; i < len(matchQuery); i++ {
		var currentMatch PreviousMatchList
		if matchQuery[i].Team == "flag_no_score" {
			// draw
			currentMatch = drawMatch(matchQuery[i])
		} else if (i == len(matchQuery)-1) || (matchQuery[i+1].ID != matchQuery[i].ID) {
			// one of the teams has score 0
			currentMatch = oneScoreMatch(matchQuery[i])
		} else {
			// both teams scored
			currentMatch = bothScoreMatch(matchQuery[i], matchQuery[i+1])
			i++
		}
		matchList = append(matchList, currentMatch)
	}

	return matchList, nil
}

func bothScoreMatch(match1, match2 PreviousMatchesQuery) PreviousMatchList {
	var currentMatch PreviousMatchList

	currentMatch.ID = match1.ID
	currentMatch.Phase = match1.Phase
	currentMatch.Type = match1.Type
	currentMatch.Team1 = match1.Team1
	currentMatch.Team2 = match1.Team2
	if currentMatch.Team1 == match1.Team {
		currentMatch.Score1 = match1.Score
		currentMatch.Score2 = match2.Score
	} else {
		currentMatch.Score1 = match2.Score
		currentMatch.Score2 = match1.Score
	}

	return currentMatch
}

func oneScoreMatch(match PreviousMatchesQuery) PreviousMatchList {
	var currentMatch PreviousMatchList

	currentMatch.ID = match.ID
	currentMatch.Phase = match.Phase
	currentMatch.Type = match.Type
	currentMatch.Team1 = match.Team1
	currentMatch.Team2 = match.Team2
	if currentMatch.Team1 == match.Team {
		currentMatch.Score1 = match.Score
		currentMatch.Score2 = 0
	} else {
		currentMatch.Score1 = 0
		currentMatch.Score2 = match.Score
	}

	return currentMatch
}

func drawMatch(match PreviousMatchesQuery) PreviousMatchList {
	var currentMatch PreviousMatchList

	currentMatch.ID = match.ID
	currentMatch.Phase = match.Phase
	currentMatch.Type = match.Type
	currentMatch.Team1 = match.Team1
	currentMatch.Team2 = match.Team2
	currentMatch.Score1 = 0
	currentMatch.Score2 = 0

	return currentMatch
}

func (matchDetails *PreviousMatch) GetPreviousMatch(db *sql.DB) error {
	statement := fmt.Sprintf(`
		SELECT 
			previous_match.id AS match_id,
			previous_match.fk_match_team1 AS team1,
			previous_match.fk_match_team2 AS team2,
			previous_match.match_type,
			previous_match.phase,
			COALESCE(player_match.quantity, 0) AS score,
			COALESCE(player_match.red, 0) AS red,
			COALESCE(player_match.yellow, 0) AS yellow,
			COALESCE(player.name, "flag_no_player") AS playerName,
			COALESCE(player.id, -1) AS playerID,
			COALESCE(player.fk_player_team, "flag_no_team") AS team
		FROM 
			previous_match
		LEFT JOIN
			player_match
				ON player_match.fk_score_match = previous_match.id
		LEFT JOIN
			player
				ON player.id = player_match.fk_score_player
		WHERE previous_match.id = %d
	`, matchDetails.ID)
	rows, err := db.Query(statement)

	if err != nil {
		return err
	}

	defer rows.Close()

	matchQuery := []PreviousMatchQuery{}

	for rows.Next() {
		var p PreviousMatchQuery
		if err := rows.Scan(&p.ID, &p.Team1, &p.Team2,
			&p.Type, &p.Phase, &p.Score, &p.Red,
			&p.Yellow, &p.PlayerName, &p.PlayerID, &p.Team); err != nil {
			return err
		}
		matchQuery = append(matchQuery, p)
	}

	matchDetails.Team1 = matchQuery[0].Team1
	matchDetails.Team2 = matchQuery[0].Team2
	matchDetails.Phase = matchQuery[0].Phase
	matchDetails.Type = matchQuery[0].Type

	for _, elem := range matchQuery {
		if elem.Team == "flag_no_team" {
			continue
		}

		playerScore := score.PlayerScore{}
		playerScore.Player.ID = elem.PlayerID
		playerScore.Player.Name = elem.PlayerName
		playerScore.Player.Team = elem.Team
		playerScore.Score = elem.Score
		playerScore.Yellow = elem.Yellow
		playerScore.Red = elem.Red
		matchDetails.PlayerScore = append(matchDetails.PlayerScore, playerScore)

		if elem.Team == elem.Team1 {
			matchDetails.Score1 += elem.Score
			matchDetails.YellowCard1 += elem.Yellow
			matchDetails.RedCard1 += elem.Red
		} else {
			matchDetails.Score2 += elem.Score
			matchDetails.YellowCard2 += elem.Yellow
			matchDetails.RedCard2 += elem.Red
		}
	}

	return nil
}

func DeletePreviousMatch(db *sql.DB, matchID int) error {
	statement := fmt.Sprintf("DELETE FROM previous_match WHERE id=%d", matchID)
	_, err := db.Exec(statement)
	return err
}
