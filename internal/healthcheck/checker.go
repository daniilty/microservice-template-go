package healthcheck

import (
	"context"

	"github.com/daniilty/microservice-template-go/internal/db"
)

// Checker - used for healthcheck.
type Checker interface {
	Check() *Info
}

type checker struct {
	dbPinger db.Pinger
}

// NewChecker - checker constructor.
func NewChecker(opts ...CheckerOption) Checker {
	c := &checker{}

	for i := range opts {
		opts[i](c)
	}

	return c
}

// Check - get app info.
func (c *checker) Check() *Info {
	dbStat := &Status{
		Ok: true,
	}

	err := c.dbPinger.Ping(context.TODO())
	if err != nil {
		dbStat.Ok = false
		dbStat.Message = err.Error()
	}

	return &Info{
		App: &BuildInfo{
			BuildTime:  BuildTime,
			CommitHash: CommitHash,
			Branch:     Branch,
		},
		DB: dbStat,
	}
}
