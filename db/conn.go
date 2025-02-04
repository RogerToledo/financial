package db

import (
	"database/sql"
	"fmt"

	"github.com/me/financial/config"
	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	sc := config.DB()

	conn, err := sql.Open("postgres", sc.StringConn)
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, fmt.Errorf("Error pinging database: %v", err)
	}

	return conn, nil
}
