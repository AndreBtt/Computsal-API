package test

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

func TestScoreAPI(t *testing.T) {

	// criar dois times
	// criar um jogador de um time
	// testar as funcoes de Score
	// deletar o jogador
	// deletar os times

}
