package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ffelipelimao/booking/config"

	_ "github.com/go-sql-driver/mysql"
)

func NewDatabase() *sql.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:3307)/%s?parseTime=true", config.DB_USER, config.DB_PASSWORD, config.DB_HOST, config.DB_DATABASE)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err.Error())
	}

	return db
}

func Close(db *sql.DB) {
	db.Close()
}
