package nextmatch

import (
	"database/sql"
	"fmt"
	"sort"
	"strconv"
)

type PreviousMatch struct {
	Team1 string
	Team2 string
}

type Match struct {
	Team1 Team
	Team2 Team
}

type NextMatchGenerate struct {
	Team1 string
	Team2 string
	Time  int
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
	for _, elem := range allTeams {
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

func getCombinationMatches(teams []Team, prevMatches []PreviousMatch) [][]Match {
	var combMatches [][]Match
	var currMatch []Match
	combinationMatches(&combMatches, currMatch, (1 << uint(len(teams))), 0, teams)

	// saving previous matches in a map
	previousMatches := make(map[string]bool)
	for _, prevM := range prevMatches {
		var key1, key2 string
		key1 = prevM.Team1 + prevM.Team2
		key2 = prevM.Team2 + prevM.Team1
		previousMatches[key1] = true
		previousMatches[key2] = true
	}

	// taking off equal elements and past matches
	equal := make(map[string]bool)
	var correctComb [][]Match
	for _, comb := range combMatches {
		var key string
		// create a unique key for the matches combination
		validMatch := true
		for _, match := range comb {
			key += match.Team1.Name
			key += match.Team2.Name
			currMatch := match.Team1.Name + match.Team2.Name
			if previousMatches[currMatch] {
				validMatch = false
			}
		}
		if validMatch && equal[key] == false {
			// this key does't exist
			correctComb = append(correctComb, comb)
			equal[key] = true
		}
	}

	return correctComb
}

func permutations(arr []int) [][]int {
	var helper func([]int, int)
	res := [][]int{}

	helper = func(arr []int, n int) {
		if n == 1 {
			tmp := make([]int, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func getTimePermutation(teamSize int, timeQt int) [][]int {
	var v []int
	for i := 0; i < teamSize; i++ {
		v = append(v, i)
	}
	for len(v) < timeQt {
		v = append(v, -1)
	}

	perms := permutations(v)

	var correctPerms [][]int
	equal := make(map[string]bool)
	for _, p := range perms {
		var key string
		// create a unique key for the matches combination
		for _, teams := range p {
			key += strconv.Itoa(teams)
		}
		if _, ok := equal[key]; !ok {
			// this key does't exist
			correctPerms = append(correctPerms, p)
			equal[key] = true
		}
	}

	return correctPerms
}

func recoverAns(groupNumber int, times []string, totalAns int, allTeams []Team, dp *[][]int, prevMatches []PreviousMatch) error {
	timeQt := len(times)
	timeAns := (1 << uint(timeQt))
	acc := 0
	var nextMatches []NextMatchGenerate

	for group := groupNumber; group > 0; group-- {

		end := false

		teamsGroup := getTeamsGroup(allTeams, group)
		matchCombination := getCombinationMatches(teamsGroup, prevMatches)
		timePermutation := getTimePermutation(len(matchCombination[0]), timeQt)
		for _, combMatch := range matchCombination {
			for _, combTime := range timePermutation {
				ok := true
				for h, t := range combTime {
					if t != -1 {
						if ((1 << uint(h)) & timeAns) > 0 {
							ok = false
							break
						}
					}
				}
				if !ok {
					continue
				}

				points := 0
				newTime := timeAns
				for h, t := range combTime {
					if t != -1 {
						if combMatch[t].Team1.Times[h] == true {
							points++
						}
						if combMatch[t].Team2.Times[h] == true {
							points++
						}
						newTime = (newTime | (1 << uint(h)))
					}
				}

				currentAns := solve(groupNumber-1, timeQt, newTime, allTeams, dp, prevMatches) + points
				if acc+currentAns == totalAns {
					end = true
					for h, t := range combTime {
						if t != -1 {
							timeAns = (timeAns | (1 << uint(h)))
							var nxtM NextMatchGenerate
							nxtM.Team1 = combMatch[t].Team1.Name
							nxtM.Team2 = combMatch[t].Team2.Name
							nxtM.Time = h
							nextMatches = append(nextMatches, nxtM)
						}
					}
					acc += points
				}
				if end {
					break
				}
			}
			if end {
				break
			}
		}
	}

	for _, elem := range nextMatches {
		fmt.Println(elem.Team1, elem.Team2, times[elem.Time])
	}

	return fmt.Errorf("Not implemented yet")
}

// timeQt = how many times we have
// time = bit mask time where 0 is available time, start with (1 << timeQt)
func solve(groupNumber int, timeQt int, time int, allTeams []Team, dp *[][]int, prevMatches []PreviousMatch) int {
	if groupNumber == 0 {
		return 0
	}
	if (*dp)[groupNumber][time] != -1 {
		return (*dp)[groupNumber][time]
	}

	teamsGroup := getTeamsGroup(allTeams, groupNumber)
	matchCombination := getCombinationMatches(teamsGroup, prevMatches)

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
					// set time h is no available more
					newTime = (newTime | (1 << uint(h)))
				}
			}
			ans = max(ans, solve(groupNumber-1, timeQt, newTime, allTeams, dp, prevMatches)+points)
		}

		bigAns = max(bigAns, ans)
	}

	(*dp)[groupNumber][time] = max((*dp)[groupNumber][time], bigAns)

	return bigAns
}

func getAvailableTimes(db *sql.DB) (map[string]int, []string, int, error) {
	// get all available times
	statement := fmt.Sprintf(`
		SELECT 
			time.time
		FROM
			time
		ORDER BY time`)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, nil, 0, err
	}

	timeIdx := make(map[string]int)

	var times []string
	for rows.Next() {
		var t string
		if err := rows.Scan(&t); err != nil {
			return nil, nil, 0, err
		}
		timeIdx[t] = len(times)
		times = append(times, t)
	}
	rows.Close()
	timeQt := len(times)

	return timeIdx, times, timeQt, nil
}

func getPreviousMatches(db *sql.DB) ([]PreviousMatch, error) {
	statement := fmt.Sprintf(`
		SELECT 
			fk_match_team1 as team1,
			fk_match_team2 as team2
		FROM
			previous_match`)
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}

	var prevMatches []PreviousMatch

	for rows.Next() {
		var p PreviousMatch
		if err := rows.Scan(&p.Team1, &p.Team2); err != nil {
			return nil, err
		}
		prevMatches = append(prevMatches, p)
	}
	rows.Close()

	return prevMatches, nil
}

func getTeams(db *sql.DB, timeQt int) (int, []Team, map[string]int, error) {
	statement := fmt.Sprintf(`
		SELECT 
			name,
			group_number
		FROM 
			team`)
	rows, err := db.Query(statement)
	if err != nil {
		return 0, nil, nil, err
	}

	groupNumber := -1
	var teams []Team

	teamIdx := make(map[string]int)

	for rows.Next() {
		var t Team
		if err := rows.Scan(&t.Name, &t.Group); err != nil {
			return 0, nil, nil, err
		}
		t.Times = make([]bool, timeQt)
		for i := 0; i < timeQt; i++ {
			t.Times[i] = true
		}
		groupNumber = max(groupNumber, t.Group)
		teamIdx[t.Name] = len(teams)
		teams = append(teams, t)
	}
	rows.Close()

	return groupNumber, teams, teamIdx, nil
}

func GenerateNextMatches(db *sql.DB) error {

	// timeIdx idicates where a time "HH:MM:SS" is index on array times
	// times is all available times
	// timeQt is how many available times it has
	timeIdx, times, timeQt, err := getAvailableTimes(db)
	if err != nil {
		return err
	}

	// get all previous matches
	prevMatches, err := getPreviousMatches(db)
	if err != nil {
		return err
	}

	// get all teams with name e group_number
	groupNumber, teams, teamIdx, err := getTeams(db, timeQt)
	if err != nil {
		return err
	}

	// set team's time availability
	statement := fmt.Sprintf(`
		SELECT 
			fk_schedule_team AS team,
			time.time AS time
		FROM 
			schedule
		INNER JOIN
			time
				ON time.id = fk_schedule_time`)
	rows, err := db.Query(statement)
	if err != nil {
		return err
	}

	for rows.Next() {
		var currName string
		var currTime string
		if err := rows.Scan(&currName, &currTime); err != nil {
			return err
		}

		teams[teamIdx[currName]].Times[timeIdx[currTime]] = false
	}
	rows.Close()

	// set all dp positions to -1
	dp := make([][]int, groupNumber+1)
	for i := 0; i <= groupNumber; i++ {
		dp[i] = make([]int, (1 << uint(timeQt+1)))
		for j := 0; j < len(dp[i]); j++ {
			dp[i][j] = -1
		}
	}

	ans := solve(groupNumber, timeQt, (1 << uint(timeQt)), teams, &dp, prevMatches)

	// recover the answer and insert in the data base
	return recoverAns(groupNumber, times, ans, teams, &dp, prevMatches)
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
