package cache

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/kamijooou/L0.service/internal/postgres"
	"github.com/kamijooou/L0.service/internal/validator"
	"github.com/kamijooou/L0.service/pkg/log"
	"go.uber.org/zap"
)

var CACHE = map[string]*validator.Order{}

func Init(ctx context.Context, conn *pgx.Conn) error {
	logger := log.LoggerFromContext(ctx)
	logger.Info("Initializing cache...")

	orders, err := postgres.ReadAll(ctx, conn)
	if err != nil {
		logger.Error("cache: read from db fails")
		return err
	}

	for _, order := range orders {
		CACHE[order.OrderUID] = order
	}

	logger.Info("Cache initialization was successful!")
	return nil
}

func Get(ctx context.Context, orderUID string) (*validator.Order, error) {
	logger := log.LoggerFromContext(ctx)

	result, ok := CACHE[orderUID]
	if !ok {
		logger.Error("Cache GET: there is no order with requested UID:", zap.String("UID", orderUID))
		return nil, errors.New("bad order_uid")
	}

	return result, nil
}
