package config

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

//Config type contains application
type Config struct {
	DB       *sqlx.DB
	Sessions *redis.Pool
}

//ApplicationConfig returns a config struct containing application specific dependencies
func ApplicationConfig() (*Config, error) {
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("PG_USER")
	redisAddr := os.Getenv("REDIS_ADDR")
	//dbPassword := os.Getenv("DB_PASSWORD")
	pqlInfo := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable",
		pgHost, pgPort, dbUser, dbName)
	redisPool := &redis.Pool{
		MaxIdle:     10,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddr)
		},
	}
	db, err := sqlx.Connect("postgres", pqlInfo)

	if err != nil {
		return &Config{}, err
	}

	return &Config{
		DB:       db,
		Sessions: redisPool,
	}, nil
}
