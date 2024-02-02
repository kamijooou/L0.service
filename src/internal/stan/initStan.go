package stan

import (
	"context"
	"sync"

	"github.com/kamijooou/L0.service/pkg/log"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"go.uber.org/zap"
)

var (
	clusterID, clientID string = "test-cluster", "l0-sub"
	URL                 string = stan.DefaultNatsURL
)

type Connection struct {
	SC stan.Conn
	NC *nats.Conn
}

func Close(ctx context.Context, sc stan.Conn, nc *nats.Conn) {
	logger := log.LoggerFromContext(ctx)
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		<-ctx.Done()
		sc.Close()
		nc.Close()
		logger.Info("nats-streaming connections was closed.")
	}()

	wg.Wait()
}

func Connect(ctx context.Context) (*Connection, error) {
	logger := log.LoggerFromContext(ctx)

	logger.Info("Connecting to NATS...")
	opts := []nats.Option{nats.Name("L0 STAN subscriber")}
	nc, err := nats.Connect(URL, opts...)
	if err != nil {
		logger.Error("NATS error:", zap.Error(err))
		return nil, err
	}
	logger.Info("Connection to NATS was successful!")

	logger.Info("Connecting to STAN...")
	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		logger.Error("STAN error:", zap.Error(err))
		return nil, err
	}
	logger.Info(
		"STAN connected to:",
		zap.String("URL", URL),
		zap.String("clusterID", clusterID),
		zap.String("clientID", clientID),
	)

	return &Connection{sc, nc}, nil
}
