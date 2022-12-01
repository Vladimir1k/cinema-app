package service

import (
	"encoding/json"
	"github.com/Vladimir1k/cinema-app"
	"github.com/Vladimir1k/cinema-app/pkg/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

func Registration(r *http.Request) {
	creds := &cinema.User{}

	err := json.NewDecoder(r.Body).Decode(creds)
	if err != nil {
		log.Println(http.StatusBadRequest)
		return
	}

	hashPassword, err := HashPassword(creds.Password)

	db, err := repository.OpenDB()
	if err != nil {
		log.Fatalf("filed toinitialize db %s", err)
	}
	defer db.Close()

	sqlStatement := `INSERT INTO users (login, password, age) VALUES ($1, $2, $3);`
	_, err = db.Exec(sqlStatement, creds.Login, hashPassword, creds.Age)
	if err != nil {
		log.Printf("регистрация не удалась: %s", err)
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(hashedPassword), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
