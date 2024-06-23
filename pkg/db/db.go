package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type ConfigDB struct {
	Driver   string
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}

func NewDB(cfg ConfigDB) *sqlx.DB {
	// Database connection string
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s", cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name)

	// Initialize a mysql database connection
	db, err := sqlx.Connect(cfg.Driver, dsn)
	if err != nil {
		panic("Failed to connect to the database: " + err.Error())
	}

	// Verify the connection to the database is still alive
	err = db.Ping()
	if err != nil {
		panic("Failed to ping the database: " + err.Error())
	}

	return db
}
