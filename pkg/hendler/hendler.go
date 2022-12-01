package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func singIn(w http.ResponseWriter, r *http.Request) {

}

func singUp(w http.ResponseWriter, r *http.Request) {

}

func creatFilm(w http.ResponseWriter, r *http.Request) {

}

func showFilms(w http.ResponseWriter, r *http.Request) {

}

func filmById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение фильма с ID %d...", id)
}

func addFavourites(w http.ResponseWriter, r *http.Request) {

}

func addWishList(w http.ResponseWriter, r *http.Request) {

}

func exportToCSV(w http.ResponseWriter, r *http.Request) {

}

func exit(w http.ResponseWriter, r *http.Request) {

}
