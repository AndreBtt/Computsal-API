package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	player "github.com/AndreBtt/Computsal/api/components/player"
	team "github.com/AndreBtt/Computsal/api/components/team"
	"github.com/gorilla/mux"
)

/* ---------------- PLAYER ROUTES --------------- */

func (a *App) createPlayer(w http.ResponseWriter, r *http.Request) {
	var p player.PlayerTable
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := p.CreatePlayer(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
}

func (a *App) deletePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid player ID")
		return
	}
	p := player.PlayerTable{ID: playerID}

	if err := p.DeletePlayer(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Player not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) updatePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid player ID")
		return
	}

	var p player.PlayerTable
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&p); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	p.ID = playerID

	if err := p.UpdatePlayer(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, p)
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

func (a *App) getPlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid player ID")
		return
	}
	p := player.Player{ID: playerID}

	if err := p.GetPlayer(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Player not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, p)
}

/* ---------------- TEAM ROUTES ----------------- */

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

	respondWithJSON(w, http.StatusCreated, t)
}

func (a *App) deleteTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	t := team.Team{Name: vars["name"]}

	if err := t.DeleteTeam(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Team not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) updateTeam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var t team.Team
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := t.UpdateTeam(a.DB, vars["name"]); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, t)
}

func (a *App) getTeams(w http.ResponseWriter, r *http.Request) {
	teams, err := team.GetTeams(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Teams not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, teams)
}

func (a *App) getTeam(w http.ResponseWriter, r *http.Request) {
	// expect to receive team's name
	vars := mux.Vars(r)
	t := team.Team{Name: vars["name"]}

	if err := t.GetTeam(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Team not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, t)
}

func (a *App) getTeamPlayers(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teams, err := team.GetPlayers(a.DB, vars["name"])
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Team not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, teams)
}
