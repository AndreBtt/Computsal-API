package group

import (
	"database/sql"
	"fmt"
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
