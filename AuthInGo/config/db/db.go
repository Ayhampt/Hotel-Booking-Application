package config

import (
	env "AuthInGo/config/env"
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

func SetupDB() (*sql.DB,error) {

	cfg := mysql.NewConfig()

	cfg.User = env.GetString("DB_USER", "root")
	cfg.Passwd = env.GetString("DB_PASS", "")
	cfg.Net = env.GetString("DB_NET", "tcp")
	cfg.Addr = env.GetString("ADDRESS", "127.0.0.1:3306")
	cfg.DBName = env.GetString("DB_NAME", "")

	fmt.Println("Connecting to database with config:", cfg.FormatDSN())
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Println("Error connecting to database:", err)
		return nil, err
	}

	if pingErr := db.Ping(); pingErr != nil {
		fmt.Println("Error pinging database:", pingErr)
		return nil, pingErr
	}

	fmt.Println("Successfully connected to database")
	return db, nil

}