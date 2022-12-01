package main

import (
	"github.com/Vladimir1k/cinema-app/pkg/hendler"
	"github.com/Vladimir1k/cinema-app/pkg/repository"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/signin", hendler.SignIn)
	mux.HandleFunc("/signup", hendler.SignUp)
	mux.HandleFunc("/exit", hendler.Exit)
	mux.HandleFunc("/creatFilm", hendler.CreatFilm)
	mux.HandleFunc("/show/Films", hendler.ShowFilms)
	mux.HandleFunc("/show/Film", hendler.FilmById)
	mux.HandleFunc("/add/Favourites", hendler.AddFavourites)
	mux.HandleFunc("/add/Wishlist", hendler.AddWishList)
	mux.HandleFunc("/export", hendler.ExportToCSV)

	db, err := repository.OpenDB()
	if err != nil {
		log.Fatalf("filed toinitialize db %s", err)
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("filed toinitialize db %s", err)
	}

	log.Println("Запуск веб-сервера на http://127.0.0.1:8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalf("error ocured while running http server: %s", err)
	}
}
