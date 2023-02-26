package postgres

import (
	"fmt"
	"os"
	"strconv"
	"taskService/config"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DbConnect *gorm.DB
}

func New(db *config.Config) (*Postgres, error) {

	dbInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.Address, strconv.Itoa(db.Port), db.User, db.Password, db.Basename)

	// Opens a new DB and attempts a Ping
	dbConn, err := gorm.Open("postgres", dbInfo)

	if err != nil {
		return nil, err
	}
	if _, ok := os.LookupEnv("POSTGRES_LOG"); ok {
		dbConn.LogMode(true)
	}
	dbConn.Update()
	pg := &Postgres{
		DbConnect: dbConn,
	}

	return pg, nil
}
