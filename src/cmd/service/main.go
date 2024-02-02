package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kamijooou/L0.service/internal/stan"
	"github.com/kamijooou/L0.service/internal/validator"
	"github.com/kamijooou/L0.service/pkg/log"

	"go.uber.org/zap"
)

func listen(ctx context.Context, msg <-chan []byte) {
	logger := log.LoggerFromContext(ctx)
	logger.Info("Listening for messages...")

	go func() {
		for {
			select {
			case msg := <-msg:
				msgStruct := &validator.Order{}
				if err := json.Unmarshal(msg, msgStruct); err != nil {
					logger.Error("JSON bad data:", zap.Error(err))
				}

				if err := msgStruct.Validate(ctx); err != nil {
					logger.Error("Validation error:", zap.Error(err))
				}
			case <-ctx.Done():
				logger.Info("Stop listening...")
				return
			}
		}
	}()

}

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	defer fmt.Println("Service stopped. Good bye...")

	logger, err := log.NewLogger()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logger.Sync()
	logger.Info("Starting L0.service...")

	logCtx := log.ContextWithLogger(ctx, logger)

	conn, err := stan.Connect(logCtx)
	if err != nil {
		logger.Error("STAN error")
		return
	}
	defer stan.Close(logCtx, conn.SC, conn.NC)

	msgCh, err := conn.Subcribe(logCtx)
	if err != nil {
		logger.Error("Subscribe error")
		return
	}

	listen(logCtx, msgCh)
}
