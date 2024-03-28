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

func initLoginHandlers() {
	http.HandleFunc("/login", LoginHandler)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		loginUser(w, r)
	} else {
		utils.LOG.Error("method not allowed", logger.Err(errors.New("method not allowed")))
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func loginUser(w http.ResponseWriter, r *http.Request) {
	var userAuth model.UserAuth
	err := json.NewDecoder(r.Body).Decode(&userAuth)
	if err != nil {
		utils.LOG.Error("failed to decode actor data", logger.Err(err))
		http.Error(w, "failed to decode actor data", http.StatusInternalServerError)
		return
	}

	user, err := repository.GetUserByLogin(userAuth.Login)
	if err != nil || user == nil {
		utils.LOG.Error("user does not exist")
		http.Error(w, "user does not exist", http.StatusForbidden)
		return
	}

	if jwt.Sha256EncodeWithSalt(userAuth.Password) != user.Password {
		utils.LOG.Info("invalid credentials")
		http.Error(w, "invalid credentials", http.StatusForbidden)
		return
	}

	utils.LOG.Info("User has been logging in: ", slog.String("user", userAuth.Login))

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
