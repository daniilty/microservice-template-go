package healthcheck

import "github.com/daniilty/microservice-template-go/internal/db"

// CheckerOption - di option.
type CheckerOption func(*checker)

// WithDBPinger - set db pinger.
func WithDBPinger(pinger db.Pinger) func(*checker) {
	return func(c *checker) {
		c.dbPinger = pinger
	}
}
