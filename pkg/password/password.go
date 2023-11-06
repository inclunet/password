package password

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Password struct {
	Rounds Rounds
}

func (p *Password) GetRound(w http.ResponseWriter, r *http.Request) {
	Repply(w, p.Rounds.GetRound(GetUrlIntParam(r, "round")))
}

func (p *Password) GetPlayer(w http.ResponseWriter, r *http.Request) {
	Repply(w, p.Rounds.GetRound(GetUrlIntParam(r, "round")).GetPlayer(GetUrlIntParam(r, "player")))
}

func (p *Password) NewPlayer(w http.ResponseWriter, r *http.Request) {
	Repply(w, p.Rounds.GetRound(GetUrlIntParam(r, "round")).AddPlayer())
}

func (p *Password) NewRound(w http.ResponseWriter, r *http.Request) {
	solo := GetUrlBoolParam(r, "solo")
	round := p.Rounds.AddRound(solo)

	if solo {
		Repply(w, round.GetPlayer(1))
	} else {
		Repply(w, round)
	}
}

func GetUrlBoolParam(r *http.Request, param string) bool {
	if intParam := GetUrlStringParam(r, param); intParam == "true" {
		return true
	}

	return false
}

func GetUrlIntParam(r *http.Request, param string) int {
	if intParam, err := strconv.Atoi(GetUrlStringParam(r, param)); err == nil {
		return intParam
	}

	return 0
}

func GetQueryString(r *http.Request, key string, value string) string {
	query := r.URL.Query()

	if value, ok := query[key]; ok {
		return value[0]
	}

	return value
}

func GetUrlStringParam(r *http.Request, param string) string {
	vars := mux.Vars(r)
	return vars[param]
}

func Repply(w http.ResponseWriter, data any) {
	response, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func NewPassword() Password {
	return Password{}
}
