package queue

import (
	"github.com/kr/beanstalk"
	"github.com/plutov/go-docker-compose-skeleton/pkg/shared"
)

// Beanstalkd struct
type Beanstalkd struct {
	Conn   *beanstalk.Conn
	Config *shared.Config
}

// Init func
func Init(c *shared.Config) (*Beanstalkd, error) {
	conn, err := beanstalk.Dial("tcp", c.QueueConn)
	if err != nil {
		return nil, err
	}

	return &Beanstalkd{
		Conn:   conn,
		Config: c,
	}, nil
}
