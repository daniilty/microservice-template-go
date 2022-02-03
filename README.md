# Microservice template for a Golang application.
## Features
  * Prometheus metrics
  * HTTP Healthcheck endpoint
  * Soon: jaeger tracing
  * Soon: GRPC healthcheck

## Configuration
### Add new db pinger:
./internal/mongo.go:
```go
func (m *mongoPinger) Ping(ctx context.Context) {
  ctx, _ = context.WithTimeout(ctx, pingTimeout)
  
  return client.Ping(ctx, readpref.Primary())
}
```
./internal/healthcheck/options.go:
```go
func WithMongoPinger(pinger db.Pinger) func(*checker) {
	return func(c *checker) {
		c.mongoPinger = pinger
	}
}
```

## ENV vars:
  * `HTTP_DEVOPS_ADDR` - addr for metics/healthcheck methods
