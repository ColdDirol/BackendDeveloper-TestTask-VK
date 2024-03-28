package utils

import (
	"database/sql"
	"log/slog"
)

var DB *sql.DB

var LOG *slog.Logger

var CONFIG *Config
