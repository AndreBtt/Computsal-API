package team

import (
	"database/sql"
	"fmt"
)

func (t *TeamTable) CreateTeam(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO team(name, photo, group_number) VALUES('%s', '%s', %d)", t.Name, t.PhotoURL, t.Group)
	if _, err := db.Exec(statement); err != nil {
		return err
	}
	if err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&t.ID); err != nil {
		return err
	}
	return nil
}

func (t *TeamTable) UpdateTeam(db *sql.DB) error {
	statement := fmt.Sprintf(`
		UPDATE 
			team
		SET
			name = '%s',
			photo = '%s',
			group_number = %d
		WHERE
			id = %d`, t.Name, t.PhotoURL, t.Group, t.ID)
	_, err := db.Exec(statement)
	return err
}

func (t *TeamTable) DeleteTeam(db *sql.DB) error {
	statement := fmt.Sprintf(`
		DELETE FROM
			team 
		WHERE 
			id = %d`, t.ID)
	_, err := db.Exec(statement)
	return err
}

func GetTeams(db *sql.DB) ([]TeamTable, error) {
	statement := fmt.Sprintf("SELECT id, name, photo, group_number FROM team")
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teams := []TeamTable{}

	for rows.Next() {
		var t TeamTable
		if err := rows.Scan(&t.ID, &t.Name, &t.PhotoURL, &t.Group); err != nil {
			return nil, err
		}
		teams = append(teams, t)
	}

	return teams, nil
}
