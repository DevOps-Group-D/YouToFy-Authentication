package repositories

import (
	"github.com/DevOps-Group-D/YouToFy-API/database"
	"github.com/DevOps-Group-D/YouToFy-API/models"
)

const (
	INSERT_QUERY = `INSERT INTO account (username, password) VALUES ($1, $2)`
	SELECT_QUERY = `SELECT * FROM account WHERE username = $1`
	UPDATE_QUERY = `UPDATE account SET password = $2, session_token = $3, csrf_token = $4 WHERE username = $1`
)

func Insert(username string, password string) error {
	conn, err := database.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRow(INSERT_QUERY, username, password)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}

func Get(username string) (*models.Account, error) {
	conn, err := database.Connect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	row := conn.QueryRow(SELECT_QUERY, username)
	if row.Err() != nil {
		return nil, row.Err()
	}

	account := &models.Account{}
	err = row.Scan(&account.Username, &account.Password, &account.SessionToken, &account.CsrfToken)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func Update(account *models.Account) error {
	conn, err := database.Connect()
	if err != nil {
		return err
	}
	defer conn.Close()

	row := conn.QueryRow(UPDATE_QUERY, account.Username, account.Password, account.SessionToken.String, account.CsrfToken.String)
	if row.Err() != nil {
		return row.Err()
	}

	return nil
}
