package postgres

import (
	"github.com/jackc/pgx/v5"
	"github.com/kamijooou/L0.service/internal/validator"
)

func scanRow(rows pgx.Rows, order *validator.Order) error {
	item := validator.Item{}

	err := rows.Scan(
		&(order.OrderUID),
		&(order.TrackNumber),
		&(order.Entry),

		&(order.Delivery.Name),
		&(order.Delivery.Phone),
		&(order.Delivery.Zip),
		&(order.Delivery.City),
		&(order.Delivery.Address),
		&(order.Delivery.Region),
		&(order.Delivery.Email),

		&(order.Payment.Transaction),
		&(order.Payment.RequestID),
		&(order.Payment.Currency),
		&(order.Payment.Provider),
		&(order.Payment.Amount),
		&(order.Payment.PaymentDT),
		&(order.Payment.Bank),
		&(order.Payment.DeliveryCost),
		&(order.Payment.GoodsTotal),
		&(order.Payment.CustomFee),

		&(item.ChrtID),
		&(item.TrackNumber),
		&(item.Price),
		&(item.RID),
		&(item.Name),
		&(item.Sale),
		&(item.Size),
		&(item.TotalPrice),
		&(item.NmID),
		&(item.Brand),
		&(item.Status),

		&(order.Locale),
		&(order.InternalSignature),
		&(order.CustomerID),
		&(order.DeliveryService),
		&(order.ShardKey),
		&(order.SmID),
		&(order.DateCreated),
		&(order.OofShard),
	)
	if err != nil {
		return err
	}

	order.Items = append(order.Items, item)
	return nil
}
