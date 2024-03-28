package main

import (
	"BackendDeveloperVK-testTask/internal/handler"
	"BackendDeveloperVK-testTask/pkg/auth"
	"BackendDeveloperVK-testTask/pkg/utils"
	"BackendDeveloperVK-testTask/pkg/utils/database"
	"BackendDeveloperVK-testTask/pkg/utils/logger"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"time"
)

func waitForPostgres() {
	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable", utils.CONFIG.Database.Host, utils.CONFIG.Database.Port, utils.CONFIG.Database.DBName, utils.CONFIG.Database.Username, utils.CONFIG.Database.Password)

	maxAttempts := 200
	attempt := 0
	for {
		db, err := sql.Open("postgres", dsn)
		if err != nil {
			utils.LOG.Error("Error connecting to PostgreSQL: %v\n", err)
		} else {
			defer db.Close()

			if err := db.Ping(); err == nil {
				utils.LOG.Info("PostgreSQL is up - executing command")
				break
			}
		}

		attempt++
		if attempt > maxAttempts {
			utils.LOG.Error("Exceeded max attempts, giving up.")
			os.Exit(1)
		}

		utils.LOG.Info("PostgreSQL is unavailable - wait", slog.String("attempt", strconv.Itoa(attempt)), slog.String("attempt", strconv.Itoa(maxAttempts)))
		time.Sleep(1 * time.Second)
	}
}

func main() {
	utils.InitConfig("config.json")

	log := logger.SetupLogger(utils.CONFIG.Env)
	log.Info("You are in mode:", slog.String("env", utils.CONFIG.Env))

	waitForPostgres()

	db, err := database.InitDatabase(
		utils.CONFIG.Database.Host,
		utils.CONFIG.Database.Port,
		utils.CONFIG.Database.Username,
		utils.CONFIG.Database.Password,
		utils.CONFIG.Database.DBName,
		utils.CONFIG.Database.SSLMode,
	)
	if err != nil {
		log.Error("failed to init database", logger.Err(err))
		os.Exit(1)
	}
	defer db.Close()

	auth.InitAuth()
	handler.InitAdvertisementHandler()

	runApplication()
}

func runApplication() {
	utils.LOG.Info("Server start listening on port: ", slog.String("port", utils.CONFIG.HTTPServer.Port))
	httpAddress := utils.CONFIG.HTTPServer.Host + ":" + utils.CONFIG.HTTPServer.Port
	err := http.ListenAndServe(httpAddress, nil)
	if err != nil {
		utils.LOG.Error("Failed to start server", logger.Err(err))
		os.Exit(1)
	}
}
