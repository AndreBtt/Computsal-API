package nextmatch

import (
	"database/sql"
	"fmt"
	"sort"
)

// func f(pairID int, pairs []pair, times int, qtTimes int, ans *[][]state, currState []state) {

// 	if pairID == len(pairs) {
// 		state = append(state, currState)
// 		return
// 	}

// 	for t := 0; t < qtTimes; t++ {
// 		if !(times & (1 << t)) {
// 			points := 0
// 			currState = append(currState, state{time: t, point: points})
// 			f(pairID+1, pairs, teams, (times | (1 << t)), qtTimes, ans, currState )
// 			currState = currState[:len(currState)-1]
// 		}
// 	}

// 	return
// }

// struct pair type {
// 	team1
// 	team2
// }

// func solve(groupNumber int, times int, qtTimes int, allTeams []team) int {

// 	if groupNumber == -1 {
// 		return 0
// 	}

// 	teams := getTeamsByGroup(group, allTeams)
// 	// [(t1,t2), (t3,t4)], [(t1,t3), (t2,t4)], ....
// 	pairsCombination := getAllPairs(teams)

// 	bigAns := 0

// 	for _, p := pairsCombination {
// 		for _, teamPairs := p {
// 			var possibleTimeMatch [][]State
// 			var currentState []State
// 			// gerou todas a combinacoes de (horario, satisfacao) para as duplas
// 			f(0, teamPairs, qtTimes, possibleTimeMatch, currentState)

// 			ans := -1

// 			for j := range likelyAns {
// 				newTimes = times
// 				totalSatisf := 0
// 				for _, comb := range likelyAns[j] {
// 					totalSatisf += comb.satisf
// 					newTimes = newTimes | (1 << comb.time)
// 				}
// 				ans = max(ans, solve(groupNumber-1, newTimes, qtTimes, allTeams) + totalSatisf)
// 			}
// 		}

// 		bigAns = max(ans, bigAns)
// 	}

// 	return bigAns
// }

// func GenerateNextMatches(db *sql.DB) error {
// 	// [grupo][1<<horarios]

// 	statement := fmt.Sprintf(`
// 		SELECT DISTINCT
// 			count(id) as total
// 		FROM
// 			time`)

// 	rows, err := db.Query(statement)
// 	if err != nil {
// 		return nil, err
// 	}

// 	totalTimes := 0
// 	for rows.Next() {
// 		if err := rows.Scan(&totalTimes); err != nil {
// 			return nil, err
// 		}
// 	}
// 	rows.Close()

// 	statement = fmt.Sprintf(`
// 		SELECT
// 			max(group_number) as group_number
// 		FROM
// 			team`)
// 	rows, err := db.Query(statement)
// 	if err != nil {
// 		return nil, err
// 	}

// 	highestGroup := 0
// 	for rows.Next() {
// 		if err := rows.Scan(&highestGroup); err != nil {
// 			return nil, err
// 		}
// 	}
// 	rows.Close()

// 	resp := solve(highestGroup, (1 << totalTimes), totalTimes, todosOsTimes)

// 	// recuperar()

// }

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

	matches := []NextMatchList{}

	for rows.Next() {
		var newMatch NextMatchList
		if err := rows.Scan(&newMatch.ID, &newMatch.Team1, &newMatch.Team2, &newMatch.Time, &newMatch.Type); err != nil {
			return nil, err
		}
		matches = append(matches, newMatch)
	}
	rows.Close()

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

	for rows.Next() {
		var newMatch NextMatchList
		newMatch.Time = ""
		if err := rows.Scan(&newMatch.ID, &newMatch.Team1, &newMatch.Team2, &newMatch.Type); err != nil {
			return nil, err
		}
		matches = append(matches, newMatch)
	}
	rows.Close()

	// sort by time and then by type
	sort.Slice(matches, func(i, j int) bool {
		if matches[i].Time == "" && matches[j].Time == "" {
			return matches[i].Type < matches[j].Type
		}
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

	// delete all elimination matches
	if _, err := tx.Exec(`TRUNCATE TABLE elimination_match`); err != nil {
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

	// don't have next matches, just final game
	if len(nextMatchesGenerate) == 0 {
		return tx.Commit()
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
