package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/plutov/go-docker-compose-skeleton/api/pkg/shared"
)

// Mysql struct
type Mysql struct {
	Conn   *sql.DB
	Config *shared.Config
}

// Init func
func Init(c *shared.Config) (*Mysql, error) {
	conn, err := sql.Open("mysql", c.DbConn)
	if err != nil {
		return nil, err
	}

	return &Mysql{
		Conn:   conn,
		Config: c,
	}, nil
}
