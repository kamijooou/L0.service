package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/jackc/pgx/v5"
	"github.com/kamijooou/L0.service/internal/postgres"
	"github.com/kamijooou/L0.service/internal/stan"
	"github.com/kamijooou/L0.service/internal/validator"
	"github.com/kamijooou/L0.service/pkg/log"

	"go.uber.org/zap"
)

func listen(ctx context.Context, msg <-chan []byte, conn *pgx.Conn) {
	logger := log.LoggerFromContext(ctx)
	logger.Info("Listening for messages...")

	go func() {
		for {
			select {
			case msg := <-msg:
				msgStruct := &validator.Order{}
				if err := json.Unmarshal(msg, msgStruct); err != nil {
					logger.Error("JSON bad data:", zap.Error(err))
					break
				}

				if err := msgStruct.Validate(ctx); err != nil {
					logger.Error("Validation error:", zap.Error(err))
					break
				}

				if err := postgres.Create(ctx, conn, msgStruct); err != nil {
					logger.Error("postgres.Create() error")
				}
			case <-ctx.Done():
				logger.Info("Stop listening...")
				return
			}
		}
	}()
}

func main() {
	fmt.Println("Starting L0.service...")
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	defer fmt.Println("Service stopped. Good bye...")

	logger, err := log.NewLogger()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logger.Sync()

	logCtx := log.ContextWithLogger(ctx, logger)

	pgxConn, err := postgres.Connect(logCtx)
	if err != nil {
		logger.Error("PGX error")
		return
	}
	defer postgres.Close(logCtx, pgxConn)

	stanConn, err := stan.Connect(logCtx)
	if err != nil {
		logger.Error("STAN error")
		return
	}
	defer stan.Close(logCtx, stanConn.SC, stanConn.NC)

	msgCh, err := stanConn.Subcribe(logCtx)
	if err != nil {
		logger.Error("Subscribe error")
		return
	}

	listen(logCtx, msgCh, pgxConn)
}
