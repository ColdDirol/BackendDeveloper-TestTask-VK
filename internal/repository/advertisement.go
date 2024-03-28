package repository

import (
	"BackendDeveloperVK-testTask/internal/models"
	"BackendDeveloperVK-testTask/pkg/utils"
	"time"
)

func Add(adv *models.Advertisement) (*models.Advertisement, error) {
	_, err := utils.DB.Exec(`
        INSERT INTO advertisements (user_id, title, description, cost, date, image_address)
        VALUES ($1, $2, $3, $4, $5, $6)
    `, adv.UserID, adv.Title, adv.Description, adv.Cost, time.Now(), adv.ImageAddress)
	if err != nil {
		return nil, err
	}

	newAdv := &models.Advertisement{}
	err = utils.DB.QueryRow(`
        SELECT id, user_id, title, description, cost, date, image_address 
        FROM advertisements 
        WHERE id = (SELECT MAX(id) FROM advertisements WHERE user_id = $1 AND title = $2 AND description = $3 AND cost = $4 AND image_address = $5)
	`, adv.UserID, adv.Title, adv.Description, adv.Cost, adv.ImageAddress).Scan(
		&newAdv.ID, &newAdv.UserID, &newAdv.Title, &newAdv.Description, &newAdv.Cost, &newAdv.Date, &newAdv.ImageAddress)

	if err != nil {
		return nil, err
	}

	return newAdv, nil
}

func GetByID(id int) (*models.Advertisement, error) {
	adv := &models.Advertisement{}
	err := utils.DB.QueryRow("SELECT * FROM advertisements WHERE id = $1", id).Scan(
		&adv.ID, &adv.UserID, &adv.Title, &adv.Description, &adv.Cost, &adv.Date, &adv.ImageAddress)
	if err != nil {
		return nil, err
	}
	return adv, nil
}

// сначала свежие
func GetAllSortedByDateDESC(page int) ([]models.Advertisement, error) {
	offset := (page - 1) * utils.CONFIG.PageElements
	rows, err := utils.DB.Query("SELECT * FROM advertisements ORDER BY date DESC LIMIT $1 OFFSET $2", utils.CONFIG.PageElements, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var advs []models.Advertisement
	for rows.Next() {
		var adv models.Advertisement
		err := rows.Scan(&adv.ID, &adv.UserID, &adv.Title, &adv.Description, &adv.Cost, &adv.Date, &adv.ImageAddress)
		if err != nil {
			return nil, err
		}
		advs = append(advs, adv)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return advs, nil
}

// сначала старые
func GetAllSortedByDateASC(page int) ([]models.Advertisement, error) {
	offset := (page - 1) * utils.CONFIG.PageElements
	rows, err := utils.DB.Query("SELECT * FROM advertisements ORDER BY date ASC LIMIT $1 OFFSET $2", utils.CONFIG.PageElements, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var advs []models.Advertisement
	for rows.Next() {
		var adv models.Advertisement
		err := rows.Scan(&adv.ID, &adv.UserID, &adv.Title, &adv.Description, &adv.Cost, &adv.Date, &adv.ImageAddress)
		if err != nil {
			return nil, err
		}
		advs = append(advs, adv)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return advs, nil
}

// сначала дешевые
func GetAllSortedByCostASC(page int) ([]models.Advertisement, error) {
	offset := (page - 1) * utils.CONFIG.PageElements
	rows, err := utils.DB.Query("SELECT * FROM advertisements ORDER BY cost ASC LIMIT $1 OFFSET $2", utils.CONFIG.PageElements, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var advs []models.Advertisement
	for rows.Next() {
		var adv models.Advertisement
		err := rows.Scan(&adv.ID, &adv.UserID, &adv.Title, &adv.Description, &adv.Cost, &adv.Date, &adv.ImageAddress)
		if err != nil {
			return nil, err
		}
		advs = append(advs, adv)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return advs, nil
}

// сначала дорогие
func GetAllSortedByCostDESC(page int) ([]models.Advertisement, error) {
	offset := (page - 1) * utils.CONFIG.PageElements
	rows, err := utils.DB.Query("SELECT * FROM advertisements ORDER BY cost DESC LIMIT $1 OFFSET $2", utils.CONFIG.PageElements, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var advs []models.Advertisement
	for rows.Next() {
		var adv models.Advertisement
		err := rows.Scan(&adv.ID, &adv.UserID, &adv.Title, &adv.Description, &adv.Cost, &adv.Date, &adv.ImageAddress)
		if err != nil {
			return nil, err
		}
		advs = append(advs, adv)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return advs, nil
}

func Update(adv *models.Advertisement) (*models.Advertisement, error) {
	_, err := utils.DB.Exec(`
        UPDATE advertisements
        SET title=$1, description=$2, cost=$3, date=$4, image_address=$5
        WHERE id=$6
    `, adv.Title, adv.Description, adv.Cost, time.Now(), adv.ImageAddress, adv.ID)
	if err != nil {
		return nil, err
	}
	adv, err = GetByID(adv.ID)
	if err != nil {
		return nil, err
	}

	return adv, nil
}

func DeleteByID(id int) (*models.Advertisement, error) {
	adv, err := GetByID(id)
	if err != nil {
		return nil, err
	}
	_, err = utils.DB.Exec("DELETE FROM advertisements WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return adv, nil
}
