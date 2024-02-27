package store

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/takumi616/go-english-vocabulary-api/config"
)

func ConnectToDatabse(ctx context.Context, config *config.Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s", config.DatabaseHost, config.DatabasePort, config.DatabaseUser, config.DatabasePassword, config.DatabaseName, config.DatabaseSSLMODE)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Printf("Failed to open database: %v", err)
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	//Verify a connection to the database is still alive, establishing a connection if necessary.
	if err = db.PingContext(ctx); err != nil {
		log.Printf("No connection to the database: %v", err)
		return nil, err
	}

	sqlxDb := sqlx.NewDb(db, "postgres")
	return sqlxDb, nil
}
