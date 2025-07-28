package servicesAcc

import (
	repositoriesAcc "github.com/DevOps-Group-D/YouToFy-API/repositories"
	"github.com/DevOps-Group-D/YouToFy-API/utils"
	"github.com/google/uuid"
)

const (
	REGISTER_QUERY = `INSERT INTO account (id, username, password) VALUES (%1, %2, %3) RETURNING id`
)

func Register(username string, password string) (uuid.UUID, error) {
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return uuid.Nil, err
	}

	return repositoriesAcc.Register(username, hashedPassword)
}
