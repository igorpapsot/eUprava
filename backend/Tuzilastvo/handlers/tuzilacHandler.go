package handlers

import (
	"Tuzilastvo/data"
	"Tuzilastvo/db"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt"
)

type TuzilacHandler struct {
	logger *log.Logger
	repo   db.Repo
}

type KeyTuzilac struct{}

func NewTuzilacHandler(l *log.Logger, ur db.Repo) *TuzilacHandler {
	return &TuzilacHandler{l, ur}
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

func (u *TuzilacHandler) LoginUser(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var logged LogUser
	err := decoder.Decode(&logged)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}
	u.logger.Println(logged)

	tuzilac, err := u.repo.Login(logged.Jmbg, logged.Lozinka)
	if err != nil {
		http.Error(rw, "Unable to login", http.StatusInternalServerError)
		u.logger.Println("Unable to login", err)
		rw.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Jmbg: tuzilac.Jmbg,
		Role: "tuzilac",
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

func (u *TuzilacHandler) Register(rw http.ResponseWriter, h *http.Request) {
	tuzilac := h.Context().Value(KeyTuzilac{}).(*data.Tuzilac)

	if u.repo.Register(tuzilac) {
		rw.WriteHeader(http.StatusAccepted)
		return
	}

	rw.WriteHeader(http.StatusNotAcceptable)
}

func (u *TuzilacHandler) GetTuzilac(rw http.ResponseWriter, h *http.Request) {
	vars := mux.Vars(h)
	var jmbg = vars["jmbg"]
	t, err := u.repo.GetTuzilac(jmbg)

	err = t.ToJSON(rw)

	if err != nil {
		http.Error(rw, unableToConvertToJson, http.StatusInternalServerError)
		u.logger.Println(unableToConvertToJson, " :", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
}

func (u *TuzilacHandler) MiddlewareTuzilacValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		tuzilac := &data.Tuzilac{}
		err := tuzilac.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		_, err = u.repo.GetTuzilac(tuzilac.Jmbg)
		if err == nil {
			rw.WriteHeader(http.StatusNotAcceptable)
			return
		}

		ctx := context.WithValue(h.Context(), KeyTuzilac{}, tuzilac)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})
}
