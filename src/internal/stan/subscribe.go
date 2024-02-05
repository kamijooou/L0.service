package stan

import (
	"context"

	"github.com/kamijooou/L0.service/pkg/log"
	"github.com/nats-io/stan.go"
	"go.uber.org/zap"
)

var channelName, durableName string = "l0-channel", "l0"

func (c *Connection) Subcribe(ctx context.Context) (<-chan []byte, error) {
	logger := log.LoggerFromContext(ctx)

	msgCh := make(chan []byte)
	mcb := func(msg *stan.Msg) {
		logger.Info("MESSAGE RECEIVED")
		msgCh <- msg.Data
	}

	_, err := c.SC.Subscribe(
		channelName,
		mcb,
		stan.DeliverAllAvailable(),
		stan.DurableName(durableName),
	)
	if err != nil {
		logger.Error("subscribe error:", zap.Error(err))
		return nil, err
	}
	logger.Info("subscribed to channel")
	return msgCh, nil
}
