package stan

import (
	"context"

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
	<-ctx.Done()
	sc.Close()
	nc.Close()
	logger.Info("nats-streaming connections was closed")
}

func Connect(ctx context.Context) (*Connection, error) {
	logger := log.LoggerFromContext(ctx)

	opts := []nats.Option{nats.Name("L0 STAN subscriber")}
	nc, err := nats.Connect(URL, opts...)
	if err != nil {
		logger.Error("NATS error:", zap.Error(err))
		return nil, err
	}
	logger.Info("connected to NATS")

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc))
	if err != nil {
		logger.Error("STAN error:", zap.Error(err))
		return nil, err
	}
	logger.Sugar().Infof("STAN connected to URL (%s) clusterID (%s) clientID (%s)", URL, clusterID, clientID)

	return &Connection{sc, nc}, nil
}
