package app

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/mhaatha/go-rest-api/helper"
)

func NewDB(dbPassword string) *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("root:%v@tcp(localhost:3306)/go_rest_api", dbPassword))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
