package DB

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

const (
	host     = "localhost"
	port     = 5432
	user_    = "postgres"
	password = "postgres"
	dbname   = "dpay_helpdesk"
)

func init_DB() (*sqlx.DB, error) {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user_, password, host, port, dbname)
	db, err := sqlx.Open("postgres", dbURL)
	return db, err
}
