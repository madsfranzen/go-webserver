package database

import (
    "context"
    "fmt"
    "os"

    "github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() error {
    databaseUrl := os.Getenv("DATABASE_URL")
    if databaseUrl == "" {
        return fmt.Errorf("DATABASE_URL not set")
    }

    pool, err := pgxpool.New(context.Background(), databaseUrl)
    if err != nil {
        return err
    }

    // Test connection
    err = pool.Ping(context.Background())
    if err != nil {
        return err
    }

    Pool = pool
    return nil
}

func Close() {
    if Pool != nil {
        Pool.Close()
    }
}

