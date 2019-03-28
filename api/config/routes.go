package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	group "github.com/AndreBtt/Computsal/api/components/group"
	match "github.com/AndreBtt/Computsal/api/components/match"
	player "github.com/AndreBtt/Computsal/api/components/player"
	score "github.com/AndreBtt/Computsal/api/components/score"
	team "github.com/AndreBtt/Computsal/api/components/team"
	"github.com/gorilla/mux"
)

/* ---------------- PLAYER ROUTES --------------- */

func (a *App) createPlayers(w http.ResponseWriter, r *http.Request) {
	var players []player.PlayerCreate
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&players); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := player.CreatePlayers(a.DB, players); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) deletePlayer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	playerID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid player ID")
		return
	}

	if err := player.DeletePlayer(a.DB, playerID); err != nil {
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

func (a *App) updatePlayers(w http.ResponseWriter, r *http.Request) {
	var players []player.PlayerUpdate
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&players); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := player.UpdatePlayers(a.DB, players); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
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
	var t team.TeamTable
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
	teamID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid team ID")
		return
	}
	t := team.TeamTable{ID: teamID}

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
	var t team.TeamTable
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := t.UpdateTeam(a.DB); err != nil {
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
}

/* ---------------- SCORE ROUTES ----------------- */

func (a *App) getScores(w http.ResponseWriter, r *http.Request) {
	playerScore, err := score.GetScores(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, playerScore)
}

/* ---------------- PREVIOUS MATCHES ROUTES ----------------- */

func (a *App) getPreviousMatches(w http.ResponseWriter, r *http.Request) {
	previousMatches, err := match.GetPreviousMatches(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Matches not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, previousMatches)
}

func (a *App) getPreviousMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid match ID")
		return
	}
	matchDetails := match.PreviousMatch{ID: matchID}
	if err = matchDetails.GetPreviousMatch(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Match not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, matchDetails)
}

func (a *App) deletePreviousMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid match ID")
		return
	}

	if err := match.DeletePreviousMatch(a.DB, matchID); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Match Not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) createPreviousMatch(w http.ResponseWriter, r *http.Request) {
	match := match.NewMatch{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&match); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := match.CreateMatch(a.DB); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, match)
}

func (a *App) updatePreviousMatch(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	matchID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid match ID")
		return
	}

	ps := []score.PlayerIDScore{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&ps); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := score.UpdateScores(a.DB, matchID, ps); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

/* ---------------- NEXT MATCHES ROUTES ----------------- */

func (a *App) updateNextMatches(w http.ResponseWriter, r *http.Request) {
	nextMatches := []match.NextMatchUpdate{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&nextMatches); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := match.UpdateNextMatches(a.DB, nextMatches); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, nextMatches)
}

func (a *App) getNextMatches(w http.ResponseWriter, r *http.Request) {
	nextMatches, err := match.GetNextMatches(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Matches not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, nextMatches)
}

/* ---------------- GROUP ROUTES ----------------- */

func (a *App) getGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := group.GetGroups(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Groups not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, groups)
}

func (a *App) updateGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid group ID")
		return
	}

	gp := []group.GroupUpdateTeam{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&gp); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := group.UpdateGroup(a.DB, groupID, gp); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})

}

func (a *App) deleteGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupID, err := strconv.Atoi(vars["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid group ID")
		return
	}

	if err := group.DeleteGroup(a.DB, groupID); err != nil {
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

func (a *App) createGroup(w http.ResponseWriter, r *http.Request) {
	teams := []group.GroupCreate{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&teams); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := group.CreateGroup(a.DB, teams); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}
