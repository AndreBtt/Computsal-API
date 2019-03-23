package player

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

func TestPlayerAPI(t *testing.T) {

	var playerID int

	/* -------------  CREATE PLAYER -------------------- */

	// Ensure that 'Team' variable exists in Team table
	pCreate := Player{Name: "nameTest", Team: "gol++"}
	if err := pCreate.CreatePlayer(db); err != nil {
		t.Fatal(err)
	}

	if err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&playerID); err != nil {
		t.Fatal(err)
	}

	/* -------------  RETRIVE PLAYER ------------------- */

	pRetrive := Player{ID: playerID}
	if err := pRetrive.GetPlayer(db); err != nil {
		t.Fatal(err)
	}

	if pRetrive != pCreate {
		t.Fatal("Create player different than get player")
	}

	/* -------------  UPDATE PLAYER -------------------- */

	pUpdate := pRetrive
	pUpdate.Name = "nameUpdateTest"
	if err := pUpdate.UpdatePlayer(db); err != nil {
		t.Fatal(err)
	}

	/* -------------  RETRIVE UPDATED PLAYER ----------- */

	if err := pRetrive.GetPlayer(db); err != nil {
		t.Fatal(err)
	}

	if pRetrive != pUpdate {
		t.Fatal("Updated player different than get player")
	}

	/* -------------  DELETE PLAYER -------------------- */

	if err := pRetrive.DeletePlayer(db); err != nil {
		t.Fatal(err)
	}
}
