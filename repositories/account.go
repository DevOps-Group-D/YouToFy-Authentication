package repositoriesAcc

import (
	"github.com/DevOps-Group-D/YouToFy-API/database"
	"github.com/google/uuid"
)

const (
	REGISTER_QUERY = `INSERT INTO account (id, username, password) VALUES ($1, $2, $3) RETURNING id`
)

func Register(username string, password string) (uuid.UUID, error) {
	conn, err := database.Connect()
	if err != nil {
		return uuid.Nil, err
	}
	defer conn.Close()

	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.Nil, err
	}

	row := conn.QueryRow(REGISTER_QUERY, id, username, password)
	if row.Err() != nil {
		return uuid.Nil, row.Err()
	}

	return id, nil
}
