package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/DevOps-Group-D/YouToFy-Authentication/models"
	"github.com/DevOps-Group-D/YouToFy-Authentication/services"
)

const TOKEN_EXPIRATION_HOURS = 24

func Register(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		errMsg := fmt.Sprintf("Error decoding invalid account: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	err = services.Register(account.Username, account.Password)
	if err != nil {
		errMsg := fmt.Sprintf("Error registering account: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(202)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var reqAccount *models.Account

	err := json.NewDecoder(r.Body).Decode(&reqAccount)
	if err != nil {
		errMsg := fmt.Sprintf("Error decoding invalid account: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	account, err := services.Login(reqAccount.Username, reqAccount.Password)
	if err != nil {
		errMsg := fmt.Sprintf("Error on account login: %s", err.Error())
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    account.SessionToken.String,
		Expires:  time.Now().Add(TOKEN_EXPIRATION_HOURS * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    account.Username,
		Expires:  time.Now().Add(TOKEN_EXPIRATION_HOURS * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "X-CSRF-Token",
		Value:    account.CsrfToken.String,
		Expires:  time.Now().Add(TOKEN_EXPIRATION_HOURS * time.Hour),
		Path:     "/",
		HttpOnly: false,
	})

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
}

func Authorize(w http.ResponseWriter, r *http.Request) {
	var reqAccount *models.Account

	err := json.NewDecoder(r.Body).Decode(&reqAccount)
	if err != nil {
		errMsg := fmt.Sprintf("Error decoding invalid account: %s", err.Error())
		http.Error(w, errMsg, http.StatusBadRequest)
		fmt.Println(errMsg)
		return
	}

	sessionToken, err := r.Cookie("session_token")
	if err != nil {
		errMsg := fmt.Sprintf("Error getting session_token cookie: %s", err.Error())
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}

	csrfToken := r.Header.Get("X-CSRF-Token")

	err = services.Authorize(reqAccount.Username, sessionToken.Value, csrfToken)
	if err != nil {
		errMsg := fmt.Sprintf("Unauthorized: %s", err.Error())
		http.Error(w, errMsg, http.StatusUnauthorized)
		fmt.Println(errMsg)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(200)
}
