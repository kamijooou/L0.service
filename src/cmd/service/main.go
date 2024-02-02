package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/kamijooou/L0.service/internal/stan"
	"github.com/kamijooou/L0.service/pkg/log"

	"go.uber.org/zap"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	// defer fmt.Println("Service stopped. Good bye...")

	logger, err := zap.NewDevelopment()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer logger.Sync()
	logger.Info("Starting L0.service...")

	workCtx := log.ContextWithLogger(ctx, logger)

	conn, err := stan.Connect(workCtx)
	if err != nil {
		logger.Error("STAN error")
		return
	}
	defer stan.Close(workCtx, conn.SC, conn.NC)

	msgCh, err := conn.Subcribe(workCtx)
	if err != nil {
		logger.Error("Subscribe error")
		return
	}

	logger.Info("Wait for messages...")
	for {
		select {
		case msg := <-msgCh:
			fmt.Println(string(msg))
		case <-ctx.Done():
			return
		}
	}
}
