package repository

import (
	"BackendDeveloperVK-testTask/pkg/auth/model"
	"BackendDeveloperVK-testTask/pkg/utils"
)

func AddUser(user model.User) error {
	_, err := utils.DB.Exec(`
        INSERT INTO users (login, password) VALUES ($1, $2)
    `, user.Login, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func GetUserByLogin(login string) (*model.User, error) {
	var user model.User
	err := utils.DB.QueryRow("SELECT id, login, password FROM users WHERE login = $1", login).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func GetUserByID(id int) (*model.User, error) {
	var user model.User
	err := utils.DB.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(&user.ID, &user.Login, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteAllUsers() error {
	_, err := utils.DB.Exec("DELETE FROM users")
	if err != nil {
		return err
	}
	return nil
}
