package nextmatch

import (
	"database/sql"
	"fmt"
	"sort"
)

type Match struct {
	Team1 Team
	Team2 Team
}

type Team struct {
	Name  string
	Times []bool
	Group int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getTeamsGroup(allTeams []Team, groupNumber int) []Team {
	var teams []Team
	for _, elem := range teams {
		if elem.Group == groupNumber {
			teams = append(teams, elem)
		}
	}
	return teams
}

func combinationMatches(combMatches *[][]Match, currMatch []Match, mask int, pairs int, teams []Team) {

	if 2*pairs == len(teams) {
		// sort using unique data (team's name) from both teams in the match
		// this will be helpeful to take off equal elements in the future
		sort.Slice(currMatch, func(i, j int) bool {
			var comp1 string
			var comp2 string
			if currMatch[i].Team1.Name < currMatch[i].Team2.Name {
				comp1 = currMatch[i].Team1.Name + currMatch[i].Team2.Name
			} else {
				comp1 = currMatch[i].Team2.Name + currMatch[i].Team1.Name
			}

			if currMatch[j].Team1.Name < currMatch[j].Team2.Name {
				comp2 = currMatch[j].Team1.Name + currMatch[j].Team2.Name
			} else {
				comp2 = currMatch[j].Team2.Name + currMatch[j].Team1.Name
			}

			return comp1 < comp2
		})
		*combMatches = append(*combMatches, currMatch)
		return
	}

	for i := 0; i < len(teams); i++ {
		if ((1 << uint(i)) & mask) == 0 {
			// free position
			for j := i + 1; j < len(teams); j++ {
				if ((1 << uint(j)) & mask) == 0 {
					// free position
					// team with index i will play against team with index j
					var m Match
					m.Team1 = teams[i]
					m.Team2 = teams[j]
					currMatch = append(currMatch, m)

					// set that these positions are not available
					newMask := mask
					newMask = newMask | (1 << uint(i))
					newMask = newMask | (1 << uint(j))

					combinationMatches(combMatches, currMatch, newMask, pairs+1, teams)
					currMatch = currMatch[:len(currMatch)-1]
				}
			}
		}
	}
}

func getCombinationMatches(teams []Team) [][]Match {
	var combMatches [][]Match
	var currMatch []Match
	combinationMatches(&combMatches, currMatch, (1 << uint(len(teams))), 0, teams)

	// taking off equal elements
	equal := make(map[string]bool)
	var correctComb [][]Match
	for _, comb := range combMatches {
		var key string
		// create a unique key for the matches combination
		for _, match := range comb {
			key += match.Team1.Name
			key += match.Team2.Name
		}
		if _, ok := equal[key]; !ok {
			// this key does't exist
			correctComb = append(correctComb, comb)
			equal[key] = true
		}
	}

	return correctComb
}

func getTimePermutation(teamSize int, timeQt int) [][]int {
	return nil
}

// timeQt = how many times we have
// time = bit mask time where 0 is available time, start with (1 << timeQt)
func solve(groupNumber int, timeQt int, time int, allTeams []Team) int {
	if groupNumber == -1 {
		return 0
	}

	teamsGroup := getTeamsGroup(allTeams, groupNumber)

	matchCombination := getCombinationMatches(teamsGroup)

	// len(matchCombination[0]) indicates how many matches we have
	// this return all combinations of matches and times
	// each position can be either -1 where indicates that there is no match in this position
	// or greater then -1 which is the matche's index number
	timePermutation := getTimePermutation(len(matchCombination[0]), timeQt)

	bigAns := 0

	for _, combMatch := range matchCombination {

		ans := 0

		for _, combTime := range timePermutation {
			ok := true
			for h, t := range combTime {
				// h is the time (hour)
				// t is the match index inside combMatch
				if t != -1 {
					// match with index t is trying to play on time h
					if ((1 << uint(h)) & time) > 0 {
						// this time is not available
						ok = false
						break
					}
				}
			}
			if !ok {
				// this combination of times is not available
				continue
			}

			points := 0
			newTime := time

			for h, t := range combTime {
				if t != -1 {
					// match with index t plays on time h
					if combMatch[t].Team1.Times[h] == true {
						// check if the first team of the match wants to play on time h
						points++
					}
					if combMatch[t].Team2.Times[h] == true {
						// check if the second team of the match wants to play on time h
						points++
					}
					// set that time h is no avaible more
					newTime = (time | (1 << uint(h)))
				}
			}

			ans = max(ans, solve(groupNumber-1, timeQt, newTime, allTeams)+points)
		}

		bigAns = max(bigAns, ans)
	}

	return bigAns
}

func GenerateNextMatches(db *sql.DB) error {
	// [grupo][1<<horarios]

	var teams []Team

	for i := 0; i < 6; i++ {
		s := fmt.Sprintf("t%d", i+1)
		t := Team{
			Name: s,
		}
		teams = append(teams, t)
	}

	tt := getCombinationMatches(teams)

	for _, elem := range tt {
		fmt.Println(elem)
	}

	return fmt.Errorf("Not implemented yet")
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
