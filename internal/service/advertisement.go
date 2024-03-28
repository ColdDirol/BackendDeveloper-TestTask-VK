package service

import (
	"BackendDeveloperVK-testTask/internal/models"
	"BackendDeveloperVK-testTask/internal/repository"
	repository2 "BackendDeveloperVK-testTask/pkg/auth/repository"
	"BackendDeveloperVK-testTask/pkg/utils"
	"BackendDeveloperVK-testTask/pkg/utils/logger"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

func PostAdvertisement(w http.ResponseWriter, r *http.Request, login string) {
	var advertisement models.Advertisement
	err := json.NewDecoder(r.Body).Decode(&advertisement)
	if err != nil {
		utils.LOG.Error("failed to decode advertisement data", logger.Err(err))
		http.Error(w, "failed to decode advertisement data", http.StatusInternalServerError)
		return
	}

	user, err := repository2.GetUserByLogin(login)
	if err != nil {
		utils.LOG.Error("user does not exist", logger.Err(err), slog.String("login", login))
		http.Error(w, "user does not exist", http.StatusInternalServerError)
		return
	}
	advertisement.UserID = user.ID

	advertisementFromDB, err := repository.Add(&advertisement)
	if err != nil {
		utils.LOG.Error("failed to insert advertisement", logger.Err(err))
		http.Error(w, "failed to insert advertisement", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(advertisementFromDB)
	if err != nil {
		utils.LOG.Error("failed to marshal advertisement data", logger.Err(err))
		http.Error(w, "failed to marshal advertisement data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func GetAdvertisementByID(w http.ResponseWriter, id int) {
	advertisement, err := repository.GetByID(id)

	if err != nil {
		utils.LOG.Error("failed to get advertisement by id", logger.Err(err))
		http.Error(w, "failed to get advertisement by id", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(advertisement)
	if err != nil {
		utils.LOG.Error("failed to marshal advertisement data", logger.Err(err))
		http.Error(w, "failed to marshal advertisement data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func GetAllAdvertisementsSortedByDateDESC(w http.ResponseWriter, page int) {
	advertisement, err := repository.GetAllSortedByDateDESC(page)
	if err != nil {
		utils.LOG.Error("failed to get advertisement", logger.Err(err))
		http.Error(w, "failed to get advertisement", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(advertisement)
	if err != nil {
		utils.LOG.Error("failed to marshal advertisement data", logger.Err(err))
		http.Error(w, "failed to marshal advertisement data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func GetAllAdvertisementsSortedByDateASC(w http.ResponseWriter, page int) {
	advertisement, err := repository.GetAllSortedByDateASC(page)
	if err != nil {
		utils.LOG.Error("failed to get advertisement", logger.Err(err))
		http.Error(w, "failed to get advertisement", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(advertisement)
	if err != nil {
		utils.LOG.Error("failed to marshal advertisement data", logger.Err(err))
		http.Error(w, "failed to marshal advertisement data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func GetAllAdvertisementsSortedByCostDESC(w http.ResponseWriter, page int) {
	advertisement, err := repository.GetAllSortedByCostDESC(page)
	if err != nil {
		utils.LOG.Error("failed to get advertisement", logger.Err(err))
		http.Error(w, "failed to get advertisement", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(advertisement)
	if err != nil {
		utils.LOG.Error("failed to marshal advertisement data", logger.Err(err))
		http.Error(w, "failed to marshal advertisement data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func GetAllAdvertisementsSortedByCostASC(w http.ResponseWriter, page int) {
	advertisement, err := repository.GetAllSortedByCostASC(page)
	if err != nil {
		utils.LOG.Error("failed to get advertisement", logger.Err(err))
		http.Error(w, "failed to get advertisement", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(advertisement)
	if err != nil {
		utils.LOG.Error("failed to marshal advertisement data", logger.Err(err))
		http.Error(w, "failed to marshal advertisement data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func UpdateAdvertisement(w http.ResponseWriter, r *http.Request, login string) {
	var advertisement models.Advertisement
	err := json.NewDecoder(r.Body).Decode(&advertisement)
	if err != nil {
		utils.LOG.Error("failed to decode advertisement data", logger.Err(err))
		http.Error(w, "failed to decode advertisement data", http.StatusInternalServerError)
		return
	}

	user, err := repository2.GetUserByLogin(login)
	if err != nil {
		utils.LOG.Error("user does not exist", logger.Err(err), slog.String("login", login))
		http.Error(w, "user does not exist", http.StatusInternalServerError)
		return
	}
	if advertisement.UserID != user.ID {
		utils.LOG.Error("failed to update advertisement - the object does not belong to the user", slog.String("userID", strconv.Itoa(user.ID)), slog.String("ownerID", strconv.Itoa(advertisement.UserID)))
		http.Error(w, "failed to update advertisement - the object does not belong to the user", http.StatusForbidden)
		return
	}

	advertisementFromDB, err := repository.Update(&advertisement)
	if err != nil {
		utils.LOG.Error("failed to update advertisement", logger.Err(err))
		http.Error(w, "failed to update advertisement", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(advertisementFromDB)
	if err != nil {
		utils.LOG.Error("failed to marshal advertisement data", logger.Err(err))
		http.Error(w, "failed to marshal advertisement data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}

func DeleteAdvertisement(w http.ResponseWriter, r *http.Request, advertisementID int, login string) {
	advertisement, err := repository.GetByID(advertisementID)
	if err != nil {
		utils.LOG.Error("failed to get advertisement by id", logger.Err(err))
		http.Error(w, "failed to get advertisement by id", http.StatusBadRequest)
		return
	}

	user, err := repository2.GetUserByLogin(login)
	if err != nil {
		utils.LOG.Error("user does not exist", logger.Err(err), slog.String("login", login))
		http.Error(w, "user does not exist", http.StatusInternalServerError)
		return
	}
	if advertisement.UserID != user.ID {
		utils.LOG.Error("failed to delete advertisement - the object does not belong to the user", slog.String("userID", strconv.Itoa(user.ID)), slog.String("ownerID", strconv.Itoa(advertisement.UserID)))
		http.Error(w, "failed to delete advertisement - the object does not belong to the user", http.StatusForbidden)
		return
	}

	advertisement, err = repository.DeleteByID(advertisementID)
	if err != nil {
		utils.LOG.Error("failed to delete advertisement", logger.Err(err))
		http.Error(w, "failed to delete advertisement", http.StatusBadRequest)
		return
	}

	jsonBytes, err := json.Marshal(advertisement)
	if err != nil {
		utils.LOG.Error("failed to marshal advertisement data", logger.Err(err))
		http.Error(w, "failed to marshal advertisement data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonBytes)
	if err != nil {
		utils.LOG.Error("failed to write response", logger.Err(err))
		http.Error(w, "failed to write response", http.StatusInternalServerError)
		return
	}
}
