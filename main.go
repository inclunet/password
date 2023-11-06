package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inclunet/password/pkg/password"
)

func main() {
	p := password.NewPassword()
	r := mux.NewRouter()
	r.HandleFunc("/api/password/new/{solo}", p.NewRound)
	r.HandleFunc("/api/password/{round}", p.GetRound)
	r.HandleFunc("/api/password/{round}/new", p.NewPlayer)
	r.HandleFunc("/api/password/{round}/{player}", p.GetPlayer)
	// r.HandleFunc("/api/card/{round}/{card}/autoplay", p.ToggleAutoplay)
	// r.HandleFunc("/api/card/{round}/1/0", p.Draw)
	// r.HandleFunc("/api/card/{round}/{card}/{number}", p.ToggleNumber)
	// r.HandleFunc("/qr/{round}/{card}", p.GetCardQR)
	// r.PathPrefix("/card/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "./bingo/build/index.html") })
	// r.PathPrefix("/").Handler(http.FileServer(http.Dir("./bingo/build/")))
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
