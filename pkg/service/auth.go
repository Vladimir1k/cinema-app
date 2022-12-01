package service

import (
	"database/sql"
	"encoding/json"
	"github.com/Vladimir1k/cinema-app"
	"github.com/Vladimir1k/cinema-app/pkg/repository"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"time"
)

type Claims struct {
	username string `json:"login"`
	jwt.StandardClaims
}

var JwtKey = []byte("my_secret_key")

func LogIn(w http.ResponseWriter, r *http.Request) {
	creds := &cinema.User{}

	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		log.Println(http.StatusBadRequest)
		return
	}

	db, err := repository.OpenDB()
	if err != nil {
		log.Fatalf("filed open db: %s", err)
	}
	defer db.Close()

	result := db.QueryRow("SELECT password FROM users WHERE login=$1", creds.Login)
	if err != nil {
		log.Printf("sql query not work: %s", err)
	}

	storedCreds := &cinema.User{}
	// Store the obtained password in `storedCreds`
	err = result.Scan(&storedCreds.Password)
	if err != nil {
		// If an entry with the username does not exist, send an "Unauthorized"(401) status
		if err == sql.ErrNoRows {
			log.Println(http.StatusUnauthorized)
			return
		}
		// If the error is of any other type, send a 500 status
		log.Println(http.StatusInternalServerError)
		return
	}

	// Compare the stored hashed password, with the hashed version of the password that was received
	if err = bcrypt.CompareHashAndPassword([]byte(storedCreds.Password), []byte(creds.Password)); err != nil {
		// If the two passwords don't match, return a 401 status
		log.Fatalf("%s", http.StatusUnauthorized)
	}
	log.Println("you compleact autorithation")

	expirationTime := time.Now().Add(5 * time.Minute) // Устанавливаем время истечения

	claims := &Claims{
		username: creds.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenClaims.SignedString(JwtKey)
	if err != nil {
		log.Println(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
}
