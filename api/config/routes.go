package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	player "github.com/AndreBtt/Computsal/api/components/player"
	team "github.com/AndreBtt/Computsal/api/components/team"
)

type dataBase struct {
	db *sql.DB
}

func (a *App) getPlayers(w http.ResponseWriter, r *http.Request) {
	players, err := player.GetPlayers(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Players not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, players)
}

func (a *App) createTeam(w http.ResponseWriter, r *http.Request) {
	var t team.Team
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := t.CreateTeam(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, u)
}
