package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/daniilty/microservice-template-go/internal/db"
	"github.com/daniilty/microservice-template-go/internal/healthcheck"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

const (
	// set more descriptive exit codes...
	exitCodeNotOK = 2
)

func run() (int, error) {
	cfg, err := loadEnvConfig()
	if err != nil {
		return exitCodeNotOK, err
	}

	appInfo := healthcheck.NewChecker(healthcheck.WithDBPinger(db.NewPinger()))

	http.DefaultServeMux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		bb, err := json.Marshal(appInfo.Check())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
			return
		}

		w.Write(bb)
	})
	http.DefaultServeMux.Handle("/metrics", promhttp.Handler())

	devopsServer := &http.Server{
		Addr:    cfg.httpDevopsAddr,
		Handler: http.DefaultServeMux,
	}

	wg := &sync.WaitGroup{}

	loggerCfg := zap.NewProductionConfig()

	logger, err := loggerCfg.Build()
	if err != nil {
		return exitCodeNotOK, err
	}

	sugared := logger.Sugar()

	wg.Add(1)
	go func() {
		sugared.Infow("Server started.", "server", "devops", "addr", devopsServer.Addr)
		err := devopsServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return
			}
			sugared.Errorw("Listen and serve.", "server", "devops", "err", err)
		}
	}()

	term := make(chan os.Signal, 1)
	signal.Notify(term, syscall.SIGINFO, syscall.SIGTERM, os.Interrupt)

	<-term
	sugared.Infow("Server shutdown.", "server", "devops")
	devopsServer.Shutdown(context.Background())

	return 0, nil
}

func main() {
	code, err := run()
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(code)
	}
}
