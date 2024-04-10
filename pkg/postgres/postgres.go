package postgres

import (
	"fmt"
	_ "github.com/lib/pq"
	postgres "gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strconv"
	"taskService/config"
	"time"
)

type Postgres struct {
	DbConnect *gorm.DB
}

func New(db *config.Config) (*Postgres, error) {

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.PG.Address, strconv.Itoa(db.PG.Port), db.PG.User, db.PG.Password, db.PG.Basename)

	// Opens a new DB and attempts a Ping
	dbConn, err := gorm.Open(postgres.Open(dbInfo), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	sqlDb, err := dbConn.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDb.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDb.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDb.SetConnMaxLifetime(time.Hour)
	if err != nil {
		return nil, err
	}
	pg := &Postgres{
		DbConnect: dbConn,
	}

	return pg, nil
}
