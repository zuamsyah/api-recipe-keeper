package db

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"api-recipe-keeper/config"
)

var db *gorm.DB

func Connect() *gorm.DB {
	host := config.Get("DB_HOST").String()
	port := config.Get("DB_PORT").String()
	name := config.Get("DB_NAME").String()
	user := config.Get("DB_USER").String()
	password := config.Get("DB_PASSWORD").String()

	dbConnPool, err := gorm.Open("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, password, host, port, name))
	if err != nil {
		log.Fatal(err)
	}

	if err := dbConnPool.DB().Ping(); err != nil {
		log.Fatal(err)
	}

	if config.Get("DB_IS_DEBUG").Bool() {
		dbConnPool = dbConnPool.Debug()
	}

	maxOpenConns := config.Get("DB_MAX_OPEN_CONNS").Int()
	maxIdleConns := config.Get("DB_MAX_IDLE_CONNS").Int()
	connMaxLifetime := config.Get("DB_CONN_MAX_LIFETIME").Duration()

	dbConnPool.DB().SetMaxIdleConns(maxIdleConns)
	dbConnPool.DB().SetMaxOpenConns(maxOpenConns)
	dbConnPool.DB().SetConnMaxLifetime(connMaxLifetime)

	db = dbConnPool
	return db
}
