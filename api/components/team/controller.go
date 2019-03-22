package team

import (
	"database/sql"
	"fmt"
)

func (t *Team) CreateTeam(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO team(name, photo, group) VALUES('%s', '%s', %d)", t.Name, t.PhotoURL, t.Group)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	return nil
}
