package models

import "database/sql"

type Account struct {
	Username     string         `json:"username"`
	Password     string         `json:"password"`
	SessionToken sql.NullString `json:"session_token"`
	CsrfToken    sql.NullString `json:"csrf_token"`
	Spotify      *Spotify       `json:"spotify"`
	Youtube      *Youtube       `json:"youtube"`
}

type Spotify struct {
	AccessToken string `json:"access_token"`
}

type Youtube struct {
	AccessToken string `json:"access_token"`
}
