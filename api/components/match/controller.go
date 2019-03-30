package match

import (
	"database/sql"
	"fmt"
	"sort"
	"strconv"

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

func GetTeamPreviousMatches(db *sql.DB, teamName string) ([]PreviousMatchList, error) {
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
		WHERE fk_match_team1 = '%s' OR fk_match_team2 = '%s'
		GROUP BY match_type, phase, team1, team2, team, match_id
		ORDER BY match_id`, teamName, teamName)

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
	matchDetails.PlayerScore = []score.PlayerScore{}

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

func (match *NewMatch) CreateMatch(db *sql.DB) error {
	/*
		get match details
			Team1
			Team2
			Match type
	*/
	var matchDetails NextMatch
	statement := fmt.Sprintf(`
		SELECT 
			fk_next_team1, fk_next_team2, type
		FROM
			next_match
		WHERE id = %d 
		`, match.NextMatchID)
	if err := db.QueryRow(statement).Scan(&matchDetails.Team1, &matchDetails.Team2, &matchDetails.Type); err != nil {
		return err
	}

	match.createMatchGroup(db, matchDetails)

	if matchDetails.Type != 0 {
		return match.createMatchElimination(db, matchDetails)
	}

	return nil

}

func (match *NewMatch) createMatchElimination(db *sql.DB, matchDetails NextMatch) error {
	// get the winner team
	score1 := 0
	score2 := 0

	// get players from each team
	statement := fmt.Sprintf(`
		SELECT
			id,
			fk_player_team
		FROM
			player
		WHERE 
			fk_player_team = '%s' OR
			fk_player_team = '%s'`, matchDetails.Team1, matchDetails.Team2)

	rows, err := db.Query(statement)
	if err != nil {
		return err
	}

	playerIDteam := make(map[int]string)

	for rows.Next() {
		var id int
		var team string
		if err := rows.Scan(&id, &team); err != nil {
			return err
		}
		playerIDteam[id] = team
	}
	rows.Close()

	// get score from each team
	for _, elem := range match.PlayerScore {
		if playerIDteam[elem.ID] == matchDetails.Team1 {
			score1 += elem.Score
		} else {
			score2 += elem.Score
		}
	}

	// check who win
	var winningTeam string
	if score1 > score2 {
		winningTeam = matchDetails.Team1
	} else {
		winningTeam = matchDetails.Team2
	}

	// search in elimination_match table for this match type
	statement = fmt.Sprintf(`
		SELECT 
			id,
			team1,
			team2,
			type
		FROM
			elimination_match
		WHERE
			team1 = "flag_type_%d" OR
			team2 = "flag_type_%d"`, matchDetails.Type, matchDetails.Type)

	rows, err = db.Query(statement)
	if err != nil {
		return err
	}
	var eliminationMatch EliminationMatchTable
	for rows.Next() {
		if err := rows.Scan(&eliminationMatch.ID, &eliminationMatch.Team1,
			&eliminationMatch.Team2, &eliminationMatch.Type); err != nil {
			return err
		}
	}
	rows.Close()

	// update the winning team in the correct field
	if eliminationMatch.Team1 == fmt.Sprintf("flag_type_%d", matchDetails.Type) {
		eliminationMatch.Team1 = winningTeam
	} else {
		eliminationMatch.Team2 = winningTeam
	}

	// check if this match is ok to transfer to next_match table
	// any of the teams have 'flag_type_' int their names
	// it means that this match is ok to move forward
	if ((len(eliminationMatch.Team1) < 10) || (eliminationMatch.Team1[:10] != "flag_type_")) &&
		((len(eliminationMatch.Team2) < 10) || (eliminationMatch.Team2[:10] != "flag_type_")) {

		// delete from elimination round
		statement := fmt.Sprintf(`
			DELETE FROM
				elimination_match
			WHERE
				id = %d`, eliminationMatch.ID)
		if _, err := db.Exec(statement); err != nil {
			return err
		}

		// add to next match
		statement = fmt.Sprintf(`
			INSERT INTO
				next_match (fk_next_team1, fk_next_team2, time, type)
			VALUES('%s', '%s', %d, %d)`, eliminationMatch.Team1, eliminationMatch.Team2, -1, eliminationMatch.Type)
		_, err := db.Exec(statement)
		return err
	}
	// not ready to move
	// update elimination round
	statement = fmt.Sprintf(`
			UPDATE
				elimination_match
			SET 
				team1 = '%s',
				team2 = '%s'
			WHERE id = %d`, eliminationMatch.Team1, eliminationMatch.Team2, eliminationMatch.ID)
	_, err = db.Exec(statement)
	return err
}

func (match *NewMatch) createMatchGroup(db *sql.DB, matchDetails NextMatch) error {
	if err := match.createPreviousMatch(db, matchDetails); err != nil {
		return err
	}

	// delete the next match related to the previous match
	return DeleteNextMatch(db, match.NextMatchID)
}

func (match *NewMatch) createPreviousMatch(db *sql.DB, matchDetails NextMatch) error {
	// get the phase with the highest number between the two teams
	phase, err := getMatchPhase(matchDetails, db)
	if err != nil {
		return err
	}

	// create the previous match with the data we got
	statement := fmt.Sprintf(`
		INSERT INTO 
			previous_match
				(fk_match_team1, fk_match_team2, match_type, phase) 
		VALUES
			('%s', '%s', %d, %d);
		`,
		matchDetails.Team1, matchDetails.Team2, matchDetails.Type, phase)
	if _, err := db.Exec(statement); err != nil {
		return err
	}

	var matchID int
	if err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&matchID); err != nil {
		return err
	}

	if len(match.PlayerScore) == 0 {
		return nil
	}

	// add all players score and card to the new previous match
	statement = generateStatement(match, matchID)
	_, err = db.Exec(statement)
	return err
}

func getMatchPhase(match NextMatch, db *sql.DB) (int, error) {
	var phase int
	statement := fmt.Sprintf(`
		SELECT
			coalesce(max(phase),0) AS phase
		FROM
			previous_match
		WHERE 
			fk_match_team1 = '%s' OR 
			fk_match_team2 = '%s' OR
			fk_match_team1 = '%s' OR 
			fk_match_team2 = '%s'`,
		match.Team1, match.Team1, match.Team2, match.Team2)

	if err := db.QueryRow(statement).Scan(&phase); err != nil {
		return 0, err
	}

	return phase + 1, nil
}

func generateStatement(match *NewMatch, matchID int) string {
	statement := fmt.Sprintf(`
		INSERT INTO
			player_match
				(fk_score_player, fk_score_match, quantity, yellow, red)
		VALUES`)

	for _, elem := range match.PlayerScore {
		values := "(" + strconv.Itoa(elem.ID) + ", " +
			strconv.Itoa(matchID) + ", " + strconv.Itoa(elem.Score) +
			", " + strconv.Itoa(elem.Yellow) + ", " + strconv.Itoa(elem.Red) + "),"
		statement += values
	}

	statement = statement[:len(statement)-1]
	return statement
}

func DeleteNextMatch(db *sql.DB, matchID int) error {
	statement := fmt.Sprintf(`
		DELETE FROM
			next_match
		WHERE 
			id = %d
		`,
		matchID)
	_, err := db.Exec(statement)
	return err
}

func UpdateNextMatches(db *sql.DB, matches []NextMatchUpdate) error {
	if matches[0].Type == 0 {
		// group phase round
		err := updateGroupPhase(db, matches)
		return err

	} else {
		// elimination round
		err := updateEliminationPhase(db, matches)
		return err
	}
}

func updateEliminationPhase(db *sql.DB, matches []NextMatchUpdate) error {
	// Elimination phase can only update time
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, elem := range matches {
		statement := fmt.Sprintf(`
			UPDATE 
				next_match 
			SET 
				time = %d	
			WHERE id = %d`, elem.Time, elem.ID)
		if _, err := tx.Exec(statement); err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return rollbackErr
			}
			return err
		}
	}

	return tx.Commit()
}

func updateGroupPhase(db *sql.DB, matches []NextMatchUpdate) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, elem := range matches {
		statement := fmt.Sprintf(`
			UPDATE 
				next_match 
			SET 
				fk_next_team1 = '%s',
				fk_next_team2 = '%s',
				time = %d	
			WHERE id = %d`, elem.Team1, elem.Team2, elem.Time, elem.ID)
		if _, err := tx.Exec(statement); err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return rollbackErr
			}
			return err
		}
	}

	return tx.Commit()
}

func GetNextMatches(db *sql.DB) ([]NextMatchList, error) {
	statement := fmt.Sprintf(`
		SELECT 
			next_match.id,
			next_match.fk_next_team1,
			next_match.fk_next_team2,
			coalesce(time.time, "") time,
			next_match.type
		FROM 
			next_match
			LEFT JOIN
				time
					ON time.id = next_match.time`)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	matches := []NextMatchList{}

	for rows.Next() {
		var newMatch NextMatchList
		if err := rows.Scan(&newMatch.ID, &newMatch.Team1, &newMatch.Team2, &newMatch.Time, &newMatch.Type); err != nil {
			return nil, err
		}
		matches = append(matches, newMatch)
	}

	// group phase matches
	if len(matches) == 0 || matches[0].Type == 0 {
		sort.Slice(matches, func(i, j int) bool {
			if matches[i].Time == "" {
				return false
			}
			if matches[j].Time == "" {
				return true
			}
			return matches[i].Time < matches[j].Time
		})
		return matches, nil
	}

	// elimination phase matches
	statement = fmt.Sprintf(`
		SELECT 
			elimination_match.id,
			elimination_match.team1,
			elimination_match.team2,
			elimination_match.type
		FROM 
			elimination_match`)
	rows, err = db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var newMatch NextMatchList
		newMatch.Time = ""
		if err := rows.Scan(&newMatch.ID, &newMatch.Team1, &newMatch.Team2, &newMatch.Type); err != nil {
			return nil, err
		}
		matches = append(matches, newMatch)
	}

	return matches, nil
}

func CreateNextMatches(db *sql.DB, nextMatches []NextMatchCreate) error {
	nextMatchesGenerate := make([]NextMatchTable, 0)
	nextMatchesQueue := make([]NextMatchTable, 0)

	typeNumber := 1
	for _, elem := range nextMatches {
		var next NextMatchTable
		next.Team1 = elem.Team1
		next.Team2 = elem.Team2
		next.Time = -1
		next.Type = typeNumber
		nextMatchesGenerate = append(nextMatchesGenerate, next)
		nextMatchesQueue = append(nextMatchesQueue, next)
		typeNumber++
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// delete all next matches
	if _, err := tx.Exec(`TRUNCATE TABLE next_match`); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	statement := fmt.Sprintf(`
		INSERT INTO 
			next_match(fk_next_team1, fk_next_team2, time, type)
		VALUES `)
	for _, elem := range nextMatchesGenerate {
		value := fmt.Sprintf("('%s', '%s', %d, %d),", elem.Team1, elem.Team2, elem.Time, elem.Type)
		statement += value
	}
	statement = statement[:len(statement)-1]

	// insert next matches into next_match table
	if _, err := tx.Exec(statement); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	nextMatchesGenerate = make([]NextMatchTable, 0)

	for len(nextMatchesQueue) > 1 {
		match1 := nextMatchesQueue[0]
		match2 := nextMatchesQueue[1]
		nextMatchesQueue = nextMatchesQueue[2:]

		var next NextMatchTable
		next.Team1 = fmt.Sprintf("flag_type_%d", match1.Type)
		next.Team2 = fmt.Sprintf("flag_type_%d", match2.Type)
		next.Type = typeNumber
		nextMatchesGenerate = append(nextMatchesGenerate, next)
		nextMatchesQueue = append(nextMatchesQueue, next)
		typeNumber++
	}

	statement = fmt.Sprintf(`
		INSERT INTO 
			elimination_match(team1, team2, type)
		VALUES `)
	for _, elem := range nextMatchesGenerate {
		value := fmt.Sprintf("('%s', '%s', %d),", elem.Team1, elem.Team2, elem.Type)
		statement += value
	}
	statement = statement[:len(statement)-1]

	// insert next matches into elimination_match table
	if _, err := tx.Exec(statement); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	return tx.Commit()
}
