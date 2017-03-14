package env

import (
	"flag"
	"github.com/plutov/go-docker-compose-skeleton/pkg/db"
	"github.com/plutov/go-docker-compose-skeleton/pkg/queue"
	"github.com/plutov/go-docker-compose-skeleton/pkg/shared"
)

// Context struct
type Context struct {
	Config *shared.Config
	Db     *db.Mysql
	Queue  *queue.Beanstalkd
}

// New func
func New() (Context, error) {
	c := shared.Config{}
	flag.StringVar(&c.Addr, "addr", ":8080", "API address")
	flag.StringVar(&c.DbConn, "db", "root@(db:3306)/app", "DB connection string")
	flag.StringVar(&c.QueueConn, "queue", "127.0.0.1:11300", "Queue connection string")
	flag.Parse()

	d, dbErr := db.Init(&c)
	if dbErr != nil {
		return Context{}, dbErr
	}

	q, qErr := queue.Init(&c)
	if qErr != nil {
		return Context{}, qErr
	}

	return Context{
		Config: &c,
		Db:     d,
		Queue:  q,
	}, nil
}
