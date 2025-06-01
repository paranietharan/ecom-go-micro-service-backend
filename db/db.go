package db

import (
	"ecom-go-micro-service-backend/env"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // MySQL driver
	"github.com/jmoiron/sqlx"
)

type Database struct {
	db *sqlx.DB
}

func NewDatabase() (*Database, error) {
	dbUsername := env.GetEnvStr("DB_USER", "")
	dbPassword := env.GetEnvStr("DB_PASS", "")
	dbHost := env.GetEnvStr("DB_HOST", "localhost")
	dbPort := env.GetEnvStr("DB_PORT", "3306")
	dbName := env.GetEnvStr("DB_NAME", "mydb")

	if dbUsername == "" || dbPassword == "" || dbName == "" {
		return nil, fmt.Errorf("database credentials are not set in environment variables")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUsername, dbPassword, dbHost, dbPort, dbName)
	db, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() {
	if d.db != nil {
		err := d.db.Close()
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Database connection closed successfully")
		}
	}
}

func (d *Database) GetDB() *sqlx.DB {
	return d.db
}
