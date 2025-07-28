package models

import "github.com/google/uuid"

type Account struct {
	Id           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	Password     string    `json:"password"`
	SessionToken string    `json:"session_token"`
	CsrfToken    uuid.UUID `json:"csrf_token"`
}
