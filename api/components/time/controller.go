package time

import (
	"database/sql"
	"fmt"
)

func CreateTimes(db *sql.DB, times []TimeCreate) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	// delete all scheduled teams
	if _, err := tx.Exec(`TRUNCATE TABLE schedule`); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	// create new times
	statement := fmt.Sprintf("INSERT INTO adm_schedule(time) values ")
	for _, elem := range times {
		value := fmt.Sprintf("('%s'),", elem.Time)
		statement += value
	}
	statement = statement[:len(statement)-1]

	if _, err := tx.Exec(statement); err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}

	return tx.Commit()
}

func GetTimes(db *sql.DB) ([]TimeTable, error) {
	statement := fmt.Sprintf("SELECT id, time FROM adm_schedule ORDER BY time ASC")
	rows, err := db.Query(statement)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	times := []TimeTable{}

	for rows.Next() {
		var t TimeTable
		if err := rows.Scan(&t.ID, &t.Time); err != nil {
			return nil, err
		}
		times = append(times, t)
	}

	return times, nil
}

func UpdateTimes(db *sql.DB, times []TimeUpdate) error {
	var deleteElem []TimeDelete
	var updateElem []TimeTable

	for _, elem := range times {
		if elem.Action == 1 {
			updateElem = append(updateElem, TimeTable{ID: elem.ID, Time: elem.Time})
		} else {
			deleteElem = append(deleteElem, TimeDelete{ID: elem.ID})
		}
	}

	// delete all scheduled teams
	if _, err := db.Exec(`TRUNCATE TABLE schedule`); err != nil {
		return err
	}

	if err := delete(db, deleteElem); err != nil {
		return err
	}

	err := update(db, updateElem)
	return err
}

func update(db *sql.DB, times []TimeTable) error {
	if len(times) == 0 {
		return nil
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, elem := range times {
		statement := fmt.Sprintf("UPDATE adm_schedule SET time='%s' WHERE id=%d", elem.Time, elem.ID)
		if _, err := tx.Exec(statement); err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				return rollbackErr
			}
			return err
		}
	}

	return tx.Commit()
}

func delete(db *sql.DB, times []TimeDelete) error {
	if len(times) == 0 {
		return nil
	}

	statement := fmt.Sprintf("DELETE FROM adm_schedule WHERE")

	for _, elem := range times {
		value := fmt.Sprintf(" id = %d OR", elem.ID)
		statement += value
	}
	statement = statement[:len(statement)-2]

	_, err := db.Exec(statement)
	return err
}
