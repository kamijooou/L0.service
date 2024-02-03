package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/kamijooou/L0.service/internal/postgres/sql"
	"github.com/kamijooou/L0.service/internal/validator"
	"github.com/kamijooou/L0.service/pkg/log"
	"go.uber.org/zap"
)

func insertItem(ctx context.Context, tx pgx.Tx, item validator.Item) error {
	_, err := tx.Exec(
		ctx,
		sql.InsertItem,
		item.ChrtID,
		item.TrackNumber,
		item.Price,
		item.RID,
		item.Name,
		item.Sale,
		item.Size,
		item.TotalPrice,
		item.NmID,
		item.Brand,
		item.Status,
	)
	if err != nil {
		return err
	}

	return nil
}

func insertPayment(ctx context.Context, tx pgx.Tx, pay *validator.Payment) error {
	_, err := tx.Exec(
		ctx,
		sql.InsertPayment,
		pay.Transaction,
		pay.RequestID,
		pay.Currency,
		pay.Provider,
		pay.Amount,
		pay.PaymentDT,
		pay.Bank,
		pay.DeliveryCost,
		pay.GoodsTotal,
		pay.CustomFee,
	)
	if err != nil {
		return err
	}

	return nil
}

func insertDelivery(ctx context.Context, tx pgx.Tx, del *validator.Delivery) error {
	_, err := tx.Exec(
		ctx,
		sql.InsertDelivery,
		del.Name,
		del.Phone,
		del.Zip,
		del.City,
		del.Address,
		del.Region,
		del.Email,
	)
	if err != nil {
		return err
	}

	return nil
}

func insertOrder(ctx context.Context, tx pgx.Tx, ord *validator.Order) error {
	_, err := tx.Exec(
		ctx,
		sql.InsertOrder,
		ord.OrderUID,
		ord.TrackNumber,
		ord.Entry,
		ord.Locale,
		ord.InternalSignature,
		ord.CustomerID,
		ord.DeliveryService,
		ord.ShardKey,
		ord.SmID,
		ord.DateCreated,
		ord.OofShard,
	)
	if err != nil {
		return err
	}

	return nil
}

func Create(ctx context.Context, conn *pgx.Conn, msg *validator.Order) error {
	logger := log.LoggerFromContext(ctx)

	tx, err := conn.Begin(ctx)
	if err != nil {
		tx.Rollback(ctx)
		return err
	}

	if err := insertOrder(ctx, tx, msg); err != nil {
		tx.Rollback(ctx)
		logger.Error("insert order error:", zap.Error(err))
		return err
	}

	if err := insertDelivery(ctx, tx, msg.Delivery); err != nil {
		tx.Rollback(ctx)
		logger.Error("insert delivery error:", zap.Error(err))
		return err
	}

	if err := insertPayment(ctx, tx, msg.Payment); err != nil {
		tx.Rollback(ctx)
		logger.Error("insert payment error:", zap.Error(err))
		return err
	}

	for _, item := range msg.Items {
		if err := insertItem(ctx, tx, item); err != nil {
			tx.Rollback(ctx)
			logger.Error("insert item error:", zap.Error(err))
			return err
		}
	}

	if err := tx.Commit(ctx); err != nil {
		logger.Error("commit error:", zap.Error(err))
		return err
	}

	logger.Info("Order was created.")
	return nil
}
