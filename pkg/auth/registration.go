package auth

import (
	"BackendDeveloperVK-testTask/pkg/auth/jwt"
	"BackendDeveloperVK-testTask/pkg/auth/model"
	"BackendDeveloperVK-testTask/pkg/auth/repository"
	"BackendDeveloperVK-testTask/pkg/utils"
	"BackendDeveloperVK-testTask/pkg/utils/logger"
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"
)

func initRegistrationHandlers() {
	http.HandleFunc("/registration", RegistrationHandler)
}

func RegistrationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		registrateUser(w, r)
	} else {
		utils.LOG.Error("method not allowed", logger.Err(errors.New("method not allowed")))
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func registrateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.LOG.Error("failed to decode user data", logger.Err(err))
		http.Error(w, "failed to decode user data", http.StatusInternalServerError)
		return
	}

	user.Password = jwt.Sha256EncodeWithSalt(user.Password)

	err = repository.AddUser(user)
	if err != nil {
		utils.LOG.Error("failed to registration user", logger.Err(err))
		http.Error(w, "failed to registration user", http.StatusInternalServerError)
		return
	}

	utils.LOG.Info("New user: ", slog.String("user", user.Login))

	jwtToken, err := jwt.CreateToken(user.Login)
	if err != nil {
		utils.LOG.Error("failed to create token", logger.Err(err))
		http.Error(w, "failed to create token", http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(jwtToken)
	if err != nil {
		utils.LOG.Error("failed to marshal actors data", logger.Err(err))
		http.Error(w, "failed to marshal actors data", http.StatusInternalServerError)
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
