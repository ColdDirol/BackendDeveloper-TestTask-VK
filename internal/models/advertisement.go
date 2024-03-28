package models

import (
	"BackendDeveloperVK-testTask/pkg/utils"
	"time"
)

type Advertisement struct {
	ID           int       `json:"id" db:"id"`
	UserID       int       `json:"user_id" db:"user_id"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	Cost         int       `json:"cost" db:"cost"`
	Date         time.Time `json:"date" db:"date"`
	ImageAddress string    `json:"image_address" db:"image_address"`
}

func CreateAdvertisementTable() error {
	_, err := utils.DB.Exec(`
        CREATE TABLE IF NOT EXISTS advertisements (
            id SERIAL PRIMARY KEY,
            user_id INTEGER,
            title VARCHAR(100),
            description VARCHAR(1024),
            cost INTEGER,
            date TIMESTAMP,
            image_address VARCHAR(256)
        )
    `)
	if err != nil {
		return err
	}

	return nil
}
