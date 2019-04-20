package group

import (
	"database/sql"
	"fmt"
	"sort"

	"github.com/AndreBtt/Computsal/components/previousmatch"
	"github.com/AndreBtt/Computsal/components/team"
)

func GetGroups(db *sql.DB) ([]GroupList, error) {
	statement := fmt.Sprintf(`SELECT DISTINCT group_number FROM team ORDER BY group_number`)
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	groups := []GroupList{}

	for rows.Next() {
		var g GroupList
		if err := rows.Scan(&g.Number); err != nil {
			return nil, err
		}
		if g.Number == -1 {
			continue
		}
		groups = append(groups, g)
	}

	return groups, nil
}

func GetGroupsDetail(db *sql.DB) ([]Group, error) {
	groupsNumber, err := GetGroups(db)
	if err != nil {
		return nil, err
	}
	groups := make([]Group, 0)
	for _, elem := range groupsNumber {
		g := Group{Number: elem.Number}
		if err := g.GetGroup(db); err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}

	return groups, nil
}

func UpdateGroup(db *sql.DB, groupID int, gp []GroupUpdateTeam) error {
	var addTeam []string
	var removeTeam []string

	for _, elem := range gp {
		if elem.Action == 1 {
			addTeam = append(addTeam, elem.Name)
		} else {
			removeTeam = append(removeTeam, elem.Name)
		}
	}

	// add teams
	if len(addTeam) > 0 {
		statement := fmt.Sprintf("UPDATE team SET group_number = %d WHERE", groupID)
		for _, elem := range addTeam {
			query := fmt.Sprintf(" name = '%s' OR", elem)
			statement += query
		}

		statement = statement[:len(statement)-2]

		if _, err := db.Exec(statement); err != nil {
			return err
		}
	}

	// remove teams
	if len(removeTeam) > 0 {
		statement := fmt.Sprintf("UPDATE team SET group_number = %d WHERE", -1)
		for _, elem := range removeTeam {
			query := fmt.Sprintf(" name = '%s' OR", elem)
			statement += query
		}

		statement = statement[:len(statement)-2]

		if _, err := db.Exec(statement); err != nil {
			return err
		}
	}

	return nil
}

func DeleteGroup(db *sql.DB, groupID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(`SET SQL_SAFE_UPDATES=0`); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	statement := fmt.Sprintf("UPDATE team SET group_number = -1 WHERE group_number = %d", groupID)
	if _, err := tx.Exec(statement); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	if _, err := tx.Exec(`SET SQL_SAFE_UPDATES=1`); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	return tx.Commit()
}

func CreateGroup(db *sql.DB, teams []GroupCreate) error {
	// search for the correct group number
	groups, err := GetGroups(db)
	if err != nil {
		return err
	}
	groupNumber := int(1)
	for _, elem := range groups {
		if groupNumber != elem.Number {
			break
		}
		groupNumber++
	}

	statement := fmt.Sprintf("UPDATE team SET group_number = %d WHERE", groupNumber)
	for _, elem := range teams {
		query := fmt.Sprintf(" name = '%s' OR", elem.Name)
		statement += query
	}

	statement = statement[:len(statement)-2]

	_, err = db.Exec(statement)
	return err
}

func (groupDetails *Group) GetGroup(db *sql.DB) error {
	statement := fmt.Sprintf(`
		SELECT
			team.name,
			team.id,
			team.photo
		FROM 
			team
		WHERE team.group_number = %d`, groupDetails.Number)

	rows, err := db.Query(statement)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var t team.TeamGroup
		if err := rows.Scan(&t.Name, &t.ID, &t.PhotoURL); err != nil {
			return err
		}
		groupDetails.Team = append(groupDetails.Team, t)
	}

	for i := 0; i < len(groupDetails.Team); i++ {
		var err error
		var previousMatches []previousmatch.PreviousMatchList
		if previousMatches, err = previousmatch.GetTeamPreviousMatches(db, groupDetails.Team[i].Name); err != nil {
			return err
		}

		for _, elem := range previousMatches {
			if elem.Team1 == groupDetails.Team[i].Name {
				if elem.Score1 > elem.Score2 {
					groupDetails.Team[i].Win++
				} else if elem.Score1 < elem.Score2 {
					groupDetails.Team[i].Lose++
				} else {
					groupDetails.Team[i].Draw++
				}
				groupDetails.Team[i].GoalsPro += elem.Score1
				groupDetails.Team[i].GoalsAgainst += elem.Score2
			} else {
				if elem.Score2 > elem.Score1 {
					groupDetails.Team[i].Win++
				} else if elem.Score2 < elem.Score1 {
					groupDetails.Team[i].Lose++
				} else {
					groupDetails.Team[i].Draw++
				}
				groupDetails.Team[i].GoalsPro += elem.Score2
				groupDetails.Team[i].GoalsAgainst += elem.Score1
			}
		}

		groupDetails.Team[i].Points = groupDetails.Team[i].Win*3 + groupDetails.Team[i].Draw
	}

	sort.Slice(groupDetails.Team, func(i, j int) bool {
		if groupDetails.Team[i].Points > groupDetails.Team[j].Points {
			return true
		} else if groupDetails.Team[i].Points < groupDetails.Team[j].Points {
			return false
		} else if groupDetails.Team[i].Win > groupDetails.Team[j].Win {
			return true
		} else if groupDetails.Team[i].Win < groupDetails.Team[j].Win {
			return false
		} else if groupDetails.Team[i].GoalsPro > groupDetails.Team[j].GoalsPro {
			return true
		} else if groupDetails.Team[i].GoalsPro < groupDetails.Team[j].GoalsPro {
			return false
		}
		return groupDetails.Team[i].GoalsAgainst < groupDetails.Team[j].GoalsAgainst
	})

	return nil
}
