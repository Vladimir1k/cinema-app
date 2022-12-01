package hendler

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Vladimir1k/cinema-app"
	"github.com/Vladimir1k/cinema-app/pkg/repository"
	"github.com/Vladimir1k/cinema-app/pkg/service"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"strconv"
	"time"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	service.LogIn(w, r)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	service.Registration(r)
}

func CreatFilm(w http.ResponseWriter, r *http.Request) {
	err := acces(r)

	if err != nil {
		w.Write([]byte("ви не пройшли авторизацію"))
		return
	}
}

func ShowFilms(w http.ResponseWriter, r *http.Request) {
	err := acces(r)
	if err != nil {
		w.Write([]byte("ви не пройшли авторизацію"))
		return
	}
	w.Write([]byte("авторизацію пройдено"))
}

func FilmById(w http.ResponseWriter, r *http.Request) {
	err := acces(r)
	if err != nil {
		w.Write([]byte("ви не пройшли авторизацію"))
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	db, err := repository.OpenDB()
	if err != nil {
		log.Fatalf("filed toinitialize db %s", err)
	}

	defer db.Close()

	stmt := `SELECT * FROM film WHERE id = $1;`

	row := db.QueryRow(stmt, id)

	s := &cinema.Film{}

	err = row.Scan(&s.Id, &s.Name, &s.Genre, &s.DirectorId, &s.Rate, &s.Year, &s.Minutes)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			http.NotFound(w, r)

		} else {
			log.Fatal(err)
		}
	}

	fmt.Fprintf(w, "%v, %v, %v", s.Name, s.Genre, s.DirectorId)
}

func AddFavourites(w http.ResponseWriter, r *http.Request) {

}

func AddWishList(w http.ResponseWriter, r *http.Request) {

}

func ExportToCSV(w http.ResponseWriter, r *http.Request) {

}

func Exit(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   "",
		Expires: time.Unix(0, 0),
	})
	http.Redirect(w, r, "/signin", http.StatusSeeOther)
}

func acces(r *http.Request) error {
	c, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			// If the cookie is not set, return an unauthorized status
			return http.ErrNoCookie
		}
		// For any other type of error, return a bad request status
		return err
	}

	tknStr := c.Value

	claims := &service.Claims{}

	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return service.JwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			log.Println(http.StatusUnauthorized)
			return err
		}
		log.Println(http.StatusBadRequest)
		return err
	}
	if !tkn.Valid {
		log.Println(http.StatusUnauthorized)
		return err
	}
	return nil
}
