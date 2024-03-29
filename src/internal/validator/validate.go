package validator

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/kamijooou/L0.service/pkg/log"
)

func (i Item) validate(trackNum string) error {
	if i.ChrtID <= 0 {
		return errors.New("bad item.chrt_id")
	}

	if i.TrackNumber != trackNum {
		return errors.New("item.track_number doesn't equals to order.track_number")
	}

	if i.RID == "" {
		return errors.New("item.rid is empty")
	}

	if i.Name == "" {
		return errors.New("item.name is empty")
	}

	if i.NmID <= 0 {
		return errors.New("bad item.nm_id")
	}

	if i.Status <= 0 {
		return errors.New("bad item.status")
	}

	return nil
}

// TODO: make it concurrent with cancel channel
func (msg *Order) validateItems(ctx context.Context) error {
	for _, item := range msg.Items {
		if err := item.validate(msg.TrackNumber); err != nil {
			return err
		}
	}

	return nil
}

func (p *Payment) validatePayment(uid string) error {
	if p.Transaction != uid {
		return errors.New("order.order_uid doesn't equals to payment.transaction")
	}

	if p.Currency == "" {
		return errors.New("payment.currency is empty")
	}

	if p.Provider == "" {
		return errors.New("payment.provider is empty")
	}

	if p.Amount <= 0 {
		return errors.New("bad payment.amount")
	}

	if p.PaymentDT <= 0 {
		return errors.New("bad payment.payment_dt")
	}

	if p.Bank == "" {
		return errors.New("payment.bank is empty")
	}

	return nil
}

func (d *Delivery) validateDelivery() error {
	if matched, _ := regexp.MatchString(`^\+\d{10}$`, d.Phone); !matched {
		return errors.New("invalid phone number")
	}

	if d.City == "" {
		return errors.New("delivery.city is empty")
	}

	if d.Address == "" {
		return errors.New("delivery.address is empty")
	}

	if d.Region == "" {
		return errors.New("delivery.region is empty")
	}

	if matched, _ := regexp.MatchString(`^([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z0-9._-]+)$`, d.Email); !matched {
		return errors.New("invalid email")
	}

	return nil
}

func (msg *Order) Validate(ctx context.Context) error {
	logger := log.LoggerFromContext(ctx)

	if msg.OrderUID == "" {
		return errors.New("order_uid is empty")
	}

	if msg.TrackNumber == "" {
		return errors.New("track_number is empty")
	}

	if msg.Delivery == nil {
		return errors.New("there is no delivery")
	}

	if msg.Payment == nil {
		return errors.New("there is no payment")
	}

	if len(msg.Items) == 0 {
		return errors.New("there is no items")
	}

	if msg.CustomerID == "" {
		return errors.New("customer_id is empty")
	}

	if msg.SmID <= 0 {
		return errors.New("bad sm_id")
	}

	var zeroTime time.Time
	if msg.DateCreated == zeroTime {
		return errors.New("date_created is empty")
	}

	if err := msg.Delivery.validateDelivery(); err != nil {
		return err
	}

	if err := msg.Payment.validatePayment(msg.OrderUID); err != nil {
		return err
	}

	if err := msg.validateItems(ctx); err != nil {
		return err
	}

	logger.Info("message validated")
	return nil
}
