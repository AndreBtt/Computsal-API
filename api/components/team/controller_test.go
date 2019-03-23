package team

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func TestMain(m *testing.M) {
	user := "root"
	password := "andre1995"
	dbname := "Computsal"

	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)
	var err error
	db, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	code := m.Run()

	os.Exit(code)
}

func TestTeamAPI(t *testing.T) {

	var teamID int

	/* -------------  CREATE TEAM -------------------- */

	tCreate := Team{Name: "testTeam", PhotoURL: "www.url.com.br", Group: 1}
	if err := tCreate.CreateTeam(db); err != nil {
		t.Fatal(err)
	}

	if err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&teamID); err != nil {
		t.Fatal(err)
	}

	/* -------------  RETRIVE TEAM ------------------- */

	tRetrive := Team{Name: "testTeam"}
	if err := tRetrive.GetTeam(db); err != nil {
		t.Fatal(err)
	}

	if tRetrive != tCreate {
		t.Fatal("Create team different than get team")
	}

	/* -------------  UPDATE TEAM -------------------- */

	tUpdate := tRetrive
	tUpdate.Name = "testTeamUpdate"
	if err := tUpdate.UpdateTeam(db, tRetrive.Name); err != nil {
		t.Fatal(err)
	}

	/* -------------  RETRIVE UPDATED TEAM ----------- */

	tRetriveUpdate := Team{Name: "testTeamUpdate"}
	if err := tRetriveUpdate.GetTeam(db); err != nil {
		t.Fatal(err)
	}

	if tRetriveUpdate != tUpdate {
		t.Fatal("Updated team different than get team")
	}

	/* -------------  DELETE TEAM -------------------- */

	if err := tRetriveUpdate.DeleteTeam(db); err != nil {
		t.Fatal(err)
	}

}
