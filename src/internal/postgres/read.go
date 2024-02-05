package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/kamijooou/L0.service/internal/postgres/sql"
	"github.com/kamijooou/L0.service/internal/validator"
	"github.com/kamijooou/L0.service/pkg/log"
	"go.uber.org/zap"
)

func ReadAll(ctx context.Context, conn *pgx.Conn) ([]*validator.Order, error) {
	logger := log.LoggerFromContext(ctx)

	rows, err := conn.Query(ctx, sql.SelectSQL)
	if err != nil {
		logger.Error("postgres reading fails:", zap.Error(err))
		return nil, err
	}

	orders := make([]*validator.Order, 0, 256)
	idx := 0
	for rows.Next() {
		order := &validator.Order{Delivery: &validator.Delivery{}, Payment: &validator.Payment{}}
		if err := scanRow(rows, order); err != nil {
			logger.Error("scan error:", zap.Error(err))
			return nil, err
		}

		if idx != 0 {
			if orders[idx-1].OrderUID == order.OrderUID {
				orders[idx-1].Items = append(orders[idx-1].Items, order.Items...)
				continue
			}
		}

		orders = append(orders, order)
		idx++
	}

	return orders, nil
}
