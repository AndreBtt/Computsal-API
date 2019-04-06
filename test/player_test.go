package test

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/AndreBtt/Computsal/api/components/player"
	"github.com/AndreBtt/Computsal/api/components/team"
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

func TestPlayerAPI(t *testing.T) {

	/* -------------  CREATE TEAM -------------------- */

	tCreate := team.TeamTable{Name: "Fake Test Team", PhotoURL: "www.url.com.br", Group: -1}
	if err := tCreate.CreateTeam(db); err != nil {
		t.Fatal(err)
	}

	/* -------------  CREATE PLAYER -------------------- */

	// Ensure that 'Team' variable exists in Team table
	pCreate := player.PlayerTable{Name: "Fake Name Test", Team: "Fake Test Team"}
	if err := pCreate.CreatePlayer(db); err != nil {
		t.Fatal(err)
	}

	if err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&pCreate.ID); err != nil {
		t.Fatal(err)
	}

	/* -------------  RETRIVE PLAYER ------------------- */

	pRetrive := player.Player{ID: pCreate.ID}
	if err := pRetrive.GetPlayer(db); err != nil {
		t.Fatal(err)
	}

	if pRetrive.Name != pCreate.Name || pRetrive.Team != pCreate.Team || pRetrive.ID != pCreate.ID {
		t.Fatal("Create player different than get player")
	}

	/* -------------  UPDATE PLAYER -------------------- */

	pUpdate := pCreate
	pUpdate.Name = "FakeNameUpdateTest"
	if err := pUpdate.UpdatePlayer(db); err != nil {
		t.Fatal(err)
	}

	/* -------------  RETRIVE UPDATED PLAYER ----------- */

	if err := pRetrive.GetPlayer(db); err != nil {
		t.Fatal(err)
	}

	if pUpdate.Name != pRetrive.Name || pUpdate.Team != pRetrive.Team || pUpdate.ID != pRetrive.ID {
		t.Fatal("Create player different than get player")
	}

	/* -------------  DELETE PLAYER -------------------- */

	if err := pUpdate.DeletePlayer(db); err != nil {
		t.Fatal(err)
	}

	/* -------------  DELETE TEAM -------------------- */

	if err := tCreate.DeleteTeam(db); err != nil {
		t.Fatal(err)
	}
}
