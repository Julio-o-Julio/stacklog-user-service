package config

import (
	"database/sql"
	"fmt"
)

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	password = "1234"
	dbname = "postgres"
)

func ConnectDatabase() (*sql.DB, error) {
	logger = GetLogger("database")

	psqlInfo := fmt.Sprintf("host:%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		logger.Errorf("| pkg config | Connect database error: %v", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Errorf("| pkg config | Ping database error: %v", err)
		return nil, err
	}

	logger.Infof("| pkg config | Connected database to: %v", dbname)

	return db, nil
}