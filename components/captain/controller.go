package captain

import (
	"database/sql"
	"fmt"
)

func (cap *CaptainQuery) CaptainQuery(db *sql.DB, teamName string) error {
	statement := fmt.Sprintf("SELECT fk_captain_player, user_email FROM captain WHERE fk_captain_team = '%s'", teamName)
	if err := db.QueryRow(statement).Scan(&cap.PlayerID, &cap.Email); err != nil {
		return err
	}
	return nil
}

func (cap *CaptainCreate) CreateCaptain(db *sql.DB) error {
	statement := fmt.Sprintf(`
		INSERT INTO 
			captain(fk_captain_team, fk_captain_player, user_email)
		VALUES 
			('%s', %d, '%s')`, cap.Team, cap.PlayerID, cap.Email)
	_, err := db.Exec(statement)
	return err
}

func GetTeam(db *sql.DB, captainEmail string) (CaptainTeam, error) {
	statement := fmt.Sprintf("SELECT fk_captain_team as team FROM captain WHERE user_email = '%s'", captainEmail)
	var t CaptainTeam
	if err := db.QueryRow(statement).Scan(&t.Team); err != nil {
		return t, err
	}
	return t, nil
}
