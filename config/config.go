package config

import "database/sql"

var (
	logger *Logger
)

func Initialize() (*sql.DB, error) {
	logger = GetLogger("config")

	// Initializate database connection
	db, err := ConnectDatabase()
	if err != nil {
		logger.Errorf("| pkg config | Initializate connection database error: %v", err)
		return nil, err
	}

	return db, nil
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}