package handlers

import (
	"GranicnaPolicija/data"
	"GranicnaPolicija/db"
	"context"
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

type KeyPolicajac struct{}

type GPolicajacHandler struct {
	logger *log.Logger
	repo   db.GpRepo
}

type LogPolicajac struct {
	Jmbg     string
	Password string
}

type Claims struct {
	Jmbg string
	Id   string
	jwt.StandardClaims
}

type Jwt struct {
	Token string `json:"jwt"`
}

var jwtKey = []byte("secret_key")

func NewPolicajacHandler(l *log.Logger, ur db.GpRepo) *GPolicajacHandler {
	return &GPolicajacHandler{l, ur}
}

func (u *GPolicajacHandler) Register(rw http.ResponseWriter, h *http.Request) {
	policajac := h.Context().Value(KeyPolicajac{}).(*data.GPolicajac)

	if u.repo.NewGPolicajac(policajac) {
		rw.WriteHeader(http.StatusAccepted)
		rw.Write([]byte("202 - Accepted"))
		return
	}
}

func (u *GPolicajacHandler) LoginPolicajac(rw http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	var logged LogPolicajac
	err := decoder.Decode(&logged)

	fmt.Println(logged)

	if err != nil {
		http.Error(rw, "Unable to convert to json", http.StatusInternalServerError)
		u.logger.Println("Unable to convert to json :", err)
		return
	}
	u.logger.Println(logged)

	policajac, err := u.repo.Login(logged.Jmbg, logged.Password)
	if err != nil {
		http.Error(rw, "Unable to login", http.StatusInternalServerError)
		u.logger.Println("Unable to login", err)
		rw.WriteHeader(http.StatusUnauthorized)
		rw.Write([]byte("401 - Unauthorized"))
		return
	}

	expirationTime := time.Now().Add(time.Minute * 5)

	claims := &Claims{
		Jmbg: policajac.Jmbg,
		Id:   policajac.Id,
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
	rw.Write([]byte("406 - Not acceptable"))
}

func (u *GPolicajacHandler) MiddlewareGPValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		policajac := &data.GPolicajac{}
		err := policajac.FromJSON(h.Body)
		if err != nil {
			http.Error(rw, "Unable to decode json", http.StatusBadRequest)
			u.logger.Println(err)
			return
		}

		ctx := context.WithValue(h.Context(), KeyPolicajac{}, policajac)
		h = h.WithContext(ctx)

		next.ServeHTTP(rw, h)
	})

}

func (u *GPolicajacHandler) MiddlewareContentTypeSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, h *http.Request) {
		u.logger.Println("Method [", h.Method, "] - Hit path :", h.URL.Path)

		rw.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(rw, h)
	})
}
