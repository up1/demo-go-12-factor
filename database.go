package database

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

type DataBase struct {
	Client *sqlx.DB
}

func NewDatabase() (*DataBase, error) {
	connectionString := viper.GetString("db_config")
	dbConn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return &DataBase{}, fmt.Errorf("Database initial failed: %w", err)
	}
	return &DataBase{
		Client: dbConn,
	}, nil
}
