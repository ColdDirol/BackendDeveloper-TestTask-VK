package handler

import (
	"BackendDeveloperVK-testTask/internal/service"
	"BackendDeveloperVK-testTask/pkg/auth/jwt"
	"BackendDeveloperVK-testTask/pkg/auth/middleware"
	"BackendDeveloperVK-testTask/pkg/utils"
	"BackendDeveloperVK-testTask/pkg/utils/logger"
	"net/http"
	"strconv"
)

func InitAdvertisementHandler() {
	http.HandleFunc("/advertisement/sortByDate/DESC/", middleware.Middleware(advertisementSortByDateDESCHandler))
	http.HandleFunc("/advertisement/sortByDate/ASC/", middleware.Middleware(advertisementSortByDateASCHandler))
	http.HandleFunc("/advertisement/sortByCost/DESC/", middleware.Middleware(advertisementSortByCostDESCHandler))
	http.HandleFunc("/advertisement/sortByCost/ASC/", middleware.Middleware(advertisementSortByCostASCHandler))
	http.HandleFunc("/advertisement/", middleware.Middleware(advertisementByIDHandler))
	http.HandleFunc("/advertisement", middleware.Middleware(advertisementHandler))
}

func advertisementSortByDateDESCHandler(w http.ResponseWriter, r *http.Request, claims *jwt.Claims) {
	pageNum, err := strconv.Atoi(utils.ExtractLastStringParameterFromURL(r.URL.Path))
	if err != nil {
		utils.LOG.Info("Error converting string to number", logger.Err(err))
		return
	}

	if r.Method == http.MethodGet {
		service.GetAllAdvertisementsSortedByDateDESC(w, pageNum)
	} else {
		utils.LOG.Info("invalid http method")
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func advertisementSortByDateASCHandler(w http.ResponseWriter, r *http.Request, claims *jwt.Claims) {
	pageNum, err := strconv.Atoi(utils.ExtractLastStringParameterFromURL(r.URL.Path))
	if err != nil {
		utils.LOG.Info("Error converting string to number", logger.Err(err))
		return
	}

	if r.Method == http.MethodGet {
		service.GetAllAdvertisementsSortedByDateASC(w, pageNum)
	} else {
		utils.LOG.Info("invalid http method")
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func advertisementSortByCostDESCHandler(w http.ResponseWriter, r *http.Request, claims *jwt.Claims) {
	pageNum, err := strconv.Atoi(utils.ExtractLastStringParameterFromURL(r.URL.Path))
	if err != nil {
		utils.LOG.Info("Error converting string to number", logger.Err(err))
		return
	}

	if r.Method == http.MethodGet {
		service.GetAllAdvertisementsSortedByCostDESC(w, pageNum)
	} else {
		utils.LOG.Info("invalid http method")
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func advertisementSortByCostASCHandler(w http.ResponseWriter, r *http.Request, claims *jwt.Claims) {
	pageNum, err := strconv.Atoi(utils.ExtractLastStringParameterFromURL(r.URL.Path))
	if err != nil {
		utils.LOG.Info("Error converting string to number", logger.Err(err))
		return
	}

	if r.Method == http.MethodGet {
		service.GetAllAdvertisementsSortedByCostASC(w, pageNum)
	} else {
		utils.LOG.Info("invalid http method")
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func advertisementByIDHandler(w http.ResponseWriter, r *http.Request, claims *jwt.Claims) {
	id, err := strconv.Atoi(utils.ExtractLastStringParameterFromURL(r.URL.Path))
	if err != nil {
		utils.LOG.Info("Error converting string to number", logger.Err(err))
		return
	}
	switch r.Method {
	case http.MethodGet:
		service.GetAdvertisementByID(w, id)
	case http.MethodDelete:
		if claims == nil {
			utils.LOG.Info("invalid jwt token")
			http.Error(w, "invalid jwt token", http.StatusForbidden)
		}
		service.DeleteAdvertisement(w, r, id, claims.Login)
	default:
		utils.LOG.Info("invalid http method")
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}

func advertisementHandler(w http.ResponseWriter, r *http.Request, claims *jwt.Claims) {
	switch r.Method {
	case http.MethodPost:
		if claims == nil {
			utils.LOG.Info("invalid jwt token")
			http.Error(w, "invalid jwt token", http.StatusForbidden)
		}
		service.PostAdvertisement(w, r, claims.Login)
	case http.MethodPut:
		if claims == nil {
			utils.LOG.Info("invalid jwt token")
			http.Error(w, "invalid jwt token", http.StatusForbidden)
		}
		service.UpdateAdvertisement(w, r, claims.Login)
	default:
		utils.LOG.Info("invalid http method")
		http.Error(w, "invalid http method", http.StatusMethodNotAllowed)
	}
}
