package middleware

import (
	"BackendDeveloperVK-testTask/pkg/auth/jwt"
	"BackendDeveloperVK-testTask/pkg/utils"
	"BackendDeveloperVK-testTask/pkg/utils/logger"
	"errors"
	"net/http"
)

func Middleware(next func(http.ResponseWriter, *http.Request, *jwt.Claims)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		authorize, claims, err := jwt.VerifyToken(token)
		if err != nil {
			utils.LOG.Info("wrong jwt", logger.Err(err))
			authorize = false
			claims = nil
		}
		if !authorize {
			claims = nil
			switch r.Method {
			case http.MethodGet:
			default:
				utils.LOG.Error("not enough rights", logger.Err(errors.New("not enough rights")))
				http.Error(w, "not enough rights", http.StatusForbidden)
				return
			}
		}

		next(w, r, claims)
	}
}
