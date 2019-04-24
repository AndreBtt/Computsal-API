package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/AndreBtt/Computsal/components/captain"
	"github.com/AndreBtt/Computsal/components/group"
	"github.com/AndreBtt/Computsal/components/nextmatch"
	"github.com/AndreBtt/Computsal/components/player"
	"github.com/AndreBtt/Computsal/components/previousmatch"
	"github.com/AndreBtt/Computsal/components/schedule"
	"github.com/AndreBtt/Computsal/components/score"
	"github.com/AndreBtt/Computsal/components/team"
	"github.com/AndreBtt/Computsal/components/time"
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

func (a *App) deletePlayers(w http.ResponseWriter, r *http.Request) {
	var players []player.PlayerID
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&players); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := player.DeletePlayers(a.DB, players); err != nil {
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
	var t team.TeamCreate
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&t); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := team.CreateTeam(a.DB, t); err != nil {
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

	if err := team.DeleteTeam(a.DB, teamID); err != nil {
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
	var t team.TeamUpdate
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
	vars := mux.Vars(r)
	teamName := vars["teamName"]

	teamDetails := team.Team{Name: teamName}

	err := teamDetails.GetTeam(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Team not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondWithJSON(w, http.StatusOK, teamDetails)
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
	previousMatches, err := previousmatch.GetPreviousMatches(a.DB)
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
	matchDetails := previousmatch.PreviousMatch{ID: matchID}
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

	if err := previousmatch.DeletePreviousMatch(a.DB, matchID); err != nil {
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
	match := previousmatch.NewMatch{}
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

func (a *App) generateNextMatches(w http.ResponseWriter, r *http.Request) {
	err := nextmatch.GenerateNextMatches(a.DB)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func (a *App) updateNextMatches(w http.ResponseWriter, r *http.Request) {
	nextMatches := []nextmatch.NextMatchUpdate{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&nextMatches); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := nextmatch.UpdateNextMatches(a.DB, nextMatches); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getNextMatches(w http.ResponseWriter, r *http.Request) {
	nextMatches, err := nextmatch.GetNextMatches(a.DB)
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

func (a *App) createNextMatches(w http.ResponseWriter, r *http.Request) {
	nextMatches := []nextmatch.NextMatchCreate{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&nextMatches); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := nextmatch.CreateNextMatches(a.DB, nextMatches); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
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

func (a *App) getGroupsDetail(w http.ResponseWriter, r *http.Request) {
	groups, err := group.GetGroupsDetail(a.DB)
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
	groupID, err := strconv.Atoi(vars["groupNumber"])
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
	groupID, err := strconv.Atoi(vars["groupNumber"])
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

func (a *App) getGroup(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	groupNumber, err := strconv.Atoi(vars["groupNumber"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid group number")
		return
	}
	groupDetails := group.Group{Number: groupNumber}
	if err := groupDetails.GetGroup(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Group not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, groupDetails)
}

/* ---------------- CAPTAIN ROUTES ----------------- */

func (a *App) getCaptain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamName := vars["teamName"]

	var cap captain.CaptainQuery

	if err := cap.CaptainQuery(a.DB, teamName); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, cap)
}

func (a *App) getTeamCaptain(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	var team captain.CaptainTeam
	var err error

	if team, err = captain.GetTeam(a.DB, email); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, team)
}

/* ---------------- TIMES ROUTES ----------------- */

func (a *App) createTimes(w http.ResponseWriter, r *http.Request) {
	times := []time.TimeCreate{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&times); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := time.CreateTimes(a.DB, times); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusCreated, map[string]string{"result": "success"})
}

func (a *App) getTimes(w http.ResponseWriter, r *http.Request) {
	times, err := time.GetTimes(a.DB)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Times not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, times)
}

func (a *App) updateDeleteTimes(w http.ResponseWriter, r *http.Request) {
	times := []time.TimeUpdate{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&times); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := time.UpdateTimes(a.DB, times); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

/* ---------------- SCHEDULE ROUTES ----------------- */

func (a *App) getTeamSchedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamName := vars["teamName"]

	times, err := schedule.GetAvailableTimes(a.DB, teamName)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			respondWithError(w, http.StatusNotFound, "Times not found")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, times)
}

func (a *App) updateTeamSchedule(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	teamName := vars["teamName"]

	schedules := []schedule.TimeUpdate{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&schedules); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer r.Body.Close()

	if err := schedule.UpdateSchedule(a.DB, schedules, teamName); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
