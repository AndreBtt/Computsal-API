package schedule

import (
	"database/sql"
	"fmt"
)

func GetAvailableTimes(db *sql.DB, teamID int) ([]TimeAvailable, error) {

	statement := fmt.Sprintf(`SELECT
			name
		FROM
			team
		WHERE id = %d`, teamID)

	var teamName string

	if err := db.QueryRow(statement).Scan(&teamName); err != nil {
		return nil, err
	}

	statement = fmt.Sprintf(`
		SELECT
			time.id AS id,
			time.time AS time,
			CASE WHEN schedule.fk_schedule_team IS NOT NULL 
			THEN 0
			ELSE 1
		END AS availability
		FROM
			time
		LEFT JOIN
			schedule
				ON schedule.fk_schedule_time = time.id AND
				schedule.fk_schedule_team = "%s"`, teamName)

	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	times := []TimeAvailable{}

	for rows.Next() {
		var t TimeAvailable
		if err := rows.Scan(&t.TimeID, &t.Time, &t.Availability); err != nil {
			return nil, err
		}
		times = append(times, t)
	}

	return times, nil
}

func UpdateSchedule(db *sql.DB, schedules []TimeUpdate, teamID int) error {
	var available, notAvailable []TimeUpdate

	for _, elem := range schedules {
		if elem.Availability == true {
			available = append(available, elem)
		} else {
			notAvailable = append(notAvailable, elem)
		}
	}

	statement := fmt.Sprintf(`SELECT
			name
		FROM
			team
		WHERE id = %d`, teamID)

	var teamName string

	if err := db.QueryRow(statement).Scan(&teamName); err != nil {
		return err
	}

	if err := createNotAvailableTime(db, notAvailable, teamName); err != nil {
		return err
	}

	err := removeAvailableTime(db, available, teamName)
	return err
}

func createNotAvailableTime(db *sql.DB, schedules []TimeUpdate, teamName string) error {
	if len(schedules) == 0 {
		return nil
	}

	statement := fmt.Sprintf("INSERT INTO schedule(fk_schedule_team, fk_schedule_time) values")

	for _, elem := range schedules {
		value := fmt.Sprintf("('%s', %d),", teamName, elem.TimeID)
		statement += value
	}

	statement = statement[:len(statement)-1]
	_, err := db.Exec(statement)
	return err
}

func removeAvailableTime(db *sql.DB, schedules []TimeUpdate, teamName string) error {
	if len(schedules) == 0 {
		return nil
	}

	statement := fmt.Sprintf("DELETE FROM schedule WHERE")

	for _, elem := range schedules {
		value := fmt.Sprintf(" (fk_schedule_team = '%s' AND fk_schedule_time = %d) OR", teamName, elem.TimeID)
		statement += value
	}

	statement = statement[:len(statement)-2]
	_, err := db.Exec(statement)
	return err

}
