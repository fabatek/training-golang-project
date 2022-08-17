package config

import (
	"database/sql"
	"faba_traning_project/utils"
	"strconv"

	_ "github.com/lib/pq"
)

func InitDatabase() (*sql.DB, error) {

	dbUrl := utils.GetWithDefault("AWSTIO_DB_URL", "postgres://postgres:postgres@localhost:5432/dev_db?sslmode=disable")

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	maxOpenConnect, err := strconv.Atoi(utils.GetWithDefault("AWSTIO_MAX_OPEN_CONNS", "10"))
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConnect)

	return db, nil
}
