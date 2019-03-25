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

func TestScoreAPI(t *testing.T) {

	/* -------------  CREATE TWO TEAM -------------------- */

	tCreate := team.TeamTable{Name: "Fake Test Team 1", PhotoURL: "www.url.com.br", Group: -1}
	if err := tCreate.CreateTeam(db); err != nil {
		t.Fatal(err)
	}

	tCreate = team.TeamTable{Name: "Fake Test Team 2", PhotoURL: "www.url.com.br", Group: -1}
	if err := tCreate.CreateTeam(db); err != nil {
		t.Fatal(err)
	}

	/* -------------  CREATE PLAYER -------------------- */

	pCreate := player.PlayerTable{Name: "Fake Name Test", Team: "Fake Test Team 1"}
	if err := pCreate.CreatePlayer(db); err != nil {
		t.Fatal(err)
	}

	if err := db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&pCreate.ID); err != nil {
		t.Fatal(err)
	}

	/* -------------  CREATE MATCH -------------------- */

	// testar as funcoes de Score
	// deletar o jogador
	// deletar os times

}
