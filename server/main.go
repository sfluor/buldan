package main

import (
	"net/http"
	"os"

	"github.com/DataDog/datadog-go/v5/statsd"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	statsd, err := statsd.New("localhost:8125",
		statsd.WithNamespace("buldan.server"),
		statsd.WithErrorHandler(func(err error) { logger.Error("Statsd error", zap.Error(err)) }))
	if err != nil {
		logger.Fatal("failed to create statsd client", zap.Error(err))
	}

	statics := os.Args[1]
	listen := ":8080"
	if len(os.Args) == 3 {
		listen = os.Args[2]
	}

	srv := NewServer(statics, statsd, logger)

	logger.Info("Starting server", zap.String("address", listen))
	if err := http.ListenAndServe(listen, srv); err != nil {
		panic(err)
	}
}
