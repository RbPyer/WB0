package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Config struct {
	Host string
	Port string
	Username string
	Password string
	DBName string
	SSLMode string
}


func NewPostgresDB(c Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
	c.Host, c.Port, c.Username, c.DBName, c.Password, c.SSLMode))
	if err != nil {
		return nil, err
	}
	return db, nil
}