package services

import (
	"database/sql"
	"fmt"

	"github.com/DevOps-Group-D/YouToFy-API/errors"
	"github.com/DevOps-Group-D/YouToFy-API/models"
	"github.com/DevOps-Group-D/YouToFy-API/repositories"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
)

func Register(username string, password string) error {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	return repositories.Insert(username, hashedPassword)
}

func Login(username string, password string) (*models.Account, error) {
	account, err := repositories.Get(username)
	if err != nil {
		return nil, err
	}

	fmt.Println(account)

	err = utils.CheckHashedPassword(account.Password, password)
	if err != nil {
		return nil, err
	}

	sessionToken, err := utils.GenerateToken(32)
	if err != nil {
		return nil, err
	}

	csrfToken, err := utils.GenerateToken(32)
	if err != nil {
		return nil, err
	}

	account.SessionToken = sql.NullString{String: sessionToken}
	account.CsrfToken = sql.NullString{String: csrfToken}

	err = repositories.Update(account)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func Authorize(username string, sessionToken string, csrfToken string) error {
	if sessionToken == "" {
		return &errors.UnauthorizedError{}
	}

	if csrfToken == "" {
		return &errors.UnauthorizedError{}
	}

	account, err := repositories.Get(username)
	if err != nil {
		return err
	}

	if sessionToken != account.SessionToken.String {
		fmt.Println("4", username, sessionToken, csrfToken)
		return &errors.UnauthorizedError{}
	}

	if csrfToken != account.CsrfToken.String {
		fmt.Println("5", username, sessionToken, csrfToken)
		return &errors.UnauthorizedError{}
	}

	return nil
}
