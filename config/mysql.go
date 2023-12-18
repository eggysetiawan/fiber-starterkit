package config

import (
	"database/sql"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func DBConn(dbname string) (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	if dbname != "" {
		dbName = dbname
	}

	dbTimeout := os.Getenv("DB_TIMEOUT")
	dbMaxConn, err := strconv.Atoi(os.Getenv("DB_MAX_CONN"))
	if err != nil {
		return nil, err
	}

	dbMaxIdleConn, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE_CONN"))
	if err != nil {
		return nil, err
	}

	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?timeout="+dbTimeout+"s")
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(dbMaxConn)
	db.SetMaxIdleConns(dbMaxIdleConn)

	return db, nil
}
