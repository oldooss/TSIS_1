package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type Player struct {
	Name        string `json:"name"`
	Position    string `json:"position"`
	Nationality string `json:"nationality"`
}

var players = []Player{
	{"Michael Jordan", "ASG", "American"},
	{"Magic Johnson", "PG", "American"},
	{"Dirk Nowitzki", "PF", "German"},
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, player := range players {
		if player.Name == params["name"] {
			json.NewEncoder(w).Encode(player)
			return
		}
	}
	http.NotFound(w, r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Real Madrid App is healthy!"))
}
