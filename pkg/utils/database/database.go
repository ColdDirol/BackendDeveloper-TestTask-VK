package database

import (
	"BackendDeveloperVK-testTask/internal/models"
	"BackendDeveloperVK-testTask/pkg/auth/model"
	"BackendDeveloperVK-testTask/pkg/utils"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func InitDatabase(host string, port string, username string, password string, dbname string, sslmode string) (*sql.DB, error) {
	db, err := initPostgres(host, port, username, password, dbname, sslmode)
	if err != nil {
		return nil, err
	}

	utils.DB = db

	if err = model.CreateUserTable(); err != nil {
		utils.LOG.Error("Create user table error")
		return nil, err
	}
	if err = models.CreateAdvertisementTable(); err != nil {
		utils.LOG.Error("Create advertisement table error")
		return nil, err
	}

	return db, nil
}

func initPostgres(host string, port string, username string, password string, dbname string, sslmode string) (*sql.DB, error) {
	url := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=disable",
		username,
		password,
		host,
		port,
		dbname)

	db, err := sql.Open("postgres", url)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
