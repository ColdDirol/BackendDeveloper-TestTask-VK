package model

import (
	"BackendDeveloperVK-testTask/pkg/utils"
)

type User struct {
	ID       int    `json:"id" db:"id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password"`
}

type UserAuth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func CreateUserTable() error {
	_, err := utils.DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id SERIAL PRIMARY KEY,
            login VARCHAR(100),
            password VARCHAR(100),
            UNIQUE(login)
        )
    `)
	if err != nil {
		return err
	}

	return nil
}
