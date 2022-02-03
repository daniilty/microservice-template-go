package db

import "context"

// Pinger - interface to check db health.
type Pinger interface {
	// Ping - pong.
	Ping(ctx context.Context) error
}

// delete this...
// you can define your mongoPinger, redisPinger, etc...
type pinger struct{}

// Ping - ping your db.
func (p *pinger) Ping(ctx context.Context) error {
	return nil
}

func NewPinger() Pinger {
	return &pinger{}
}
