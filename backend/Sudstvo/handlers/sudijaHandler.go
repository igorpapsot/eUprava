package handlers

import (
	"Sudstvo/data"
	"Sudstvo/db"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type SudijaHandler struct {
	logger *log.Logger
	repo   db.Repo
}

type KeySudija struct{}

func NewSudijaHandler(l *log.Logger, ur db.Repo) *SudijaHandler {
	return &SudijaHandler{l, ur}
}

type LogUser struct {
	Jmbg    string `json:"jmbg"`
	Lozinka string `json:"lozinka"`
}

type Claims struct {
	Jmbg string `json:"username"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type Jwt struct {
	Token string `json:"jwt"`
}

var jwtKey = []byte("secret_key")

func (u *SudijaHandler) LoginUser(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var logged LogUser
	err := decoder.Decode(&logged)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}
	u.logger.Println(logged)

	sudija, err := u.repo.Login(logged.Jmbg, logged.Lozinka)
	if err != nil {
		http.Error(rw, "Unable to login", http.StatusInternalServerError)
		u.logger.Println("Unable to login", err)
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Jmbg: sudija.Jmbg,
		Role: "sudija",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	var userResponse Jwt
	userResponse.Token = tokenString
	jsonUser, err := json.Marshal(userResponse)
	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	rw.Write(jsonUser)
	return

	rw.WriteHeader(http.StatusNotAcceptable)
}

func (u *SudijaHandler) Register(rw http.ResponseWriter, h *http.Request) {
	sudija := h.Context().Value(KeySudija{}).(*data.Sudija)

	if u.repo.Register(sudija) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
}

func (u *SudijaHandler) MiddlewareSudijaValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		sudija := &data.Sudija{}
		err := sudija.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		_, err = u.repo.GetSudija(sudija.Jmbg)
		if err == nil {
			rw.WriteHeader(http.StatusNotAcceptable)
			return
		}

		ctx := context.WithValue(h.Context(), KeySudija{}, sudija)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
