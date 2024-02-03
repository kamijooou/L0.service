package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/kamijooou/L0.service/pkg/log"
	"go.uber.org/zap"
)

var url = "postgres://l0_user:pass@localhost:5433/l0_test"

func Connect(ctx context.Context) (*pgx.Conn, error) {
	logger := log.LoggerFromContext(ctx)
	logger.Info("Connecting to postgres...")

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		logger.Error("Connection failed:", zap.Error(err))
		return nil, err
	}

	logger.Info("Connecting to postgres was successfull!")
	return conn, nil
}

func Close(ctx context.Context, conn *pgx.Conn) {
	logger := log.LoggerFromContext(ctx)
	conn.Close(ctx)
	logger.Info("Connection to DB was closed.")
}
