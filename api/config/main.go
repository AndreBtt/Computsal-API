package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	connectionString := fmt.Sprintf("%s:%s@/%s", user, password, dbname)

	var err error
	a.DB, err = sql.Open("mysql", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()
	a.initializeRoutes()
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/players", a.getPlayers).Methods("GET")
	a.Router.HandleFunc("/players", a.createPlayer).Methods("POST")
	a.Router.HandleFunc("/players/{id}", a.deletePlayer).Methods("DELETE")
	a.Router.HandleFunc("/players/{id}", a.updatePlayer).Methods("UPDATE")
	a.Router.HandleFunc("/players/{id}", a.getPlayer).Methods("GET")

	a.Router.HandleFunc("/teams", a.getTeams).Methods("GET")
	a.Router.HandleFunc("/teams", a.createTeam).Methods("POST")
	a.Router.HandleFunc("/teams/{name}", a.deleteTeam).Methods("DELETE")
	a.Router.HandleFunc("/teams/{name}", a.updateTeam).Methods("UPDATE")
	a.Router.HandleFunc("/teams/{name}", a.getTeam).Methods("GET")
	a.Router.HandleFunc("/teams/{name}/players", a.getTeamPlayers).Methods("GET")

	a.Router.HandleFunc("/score/{matchID}", a.getScoreMatch).Methods("GET")
	a.Router.HandleFunc("/score", a.getScore).Methods("GET")
}

func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func main() {
	a := App{}
	a.Initialize("root", "andre1995", "Computsal")
	defer a.DB.Close()
	a.Run(":8080")
}