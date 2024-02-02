package validator

import "time"

// TODO: implement the comments tasks in struct fields
type Item struct {
	ChrtID      uint64 `json:"chrt_id"`
	TrackNumber string `json:"track_number"`
	Price       uint64 `json:"price"`
	RID         string `json:"rid"`
	Name        string `json:"name"`
	Sale        uint64 `json:"sale"`
	Size        string `json:"size"`
	TotalPrice  uint64 `json:"total_price"`
	NmID        uint64 `json:"nm_id"`
	Brand       string `json:"brand"`
	Status      uint64 `json:"status"`
}

type Payment struct {
	Transaction  string `json:"transaction"`
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`
	Provider     string `json:"provider"`
	Amount       uint64 `json:"amount"`
	PaymentDT    uint64 `json:"payment_dt"`
	Bank         string `json:"bank"`
	DeliveryCost uint64 `json:"delivery_cost"`
	GoodsTotal   uint64 `json:"goods_total"`
	CustomFee    uint64 `json:"custom_fee"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Zip     string `json:"zip"`
	City    string `json:"city"`
	Address string `json:"address"`
	Region  string `json:"region"`
	Email   string `json:"email"`
}

type Order struct {
	OrderUID          string    `json:"order_uid"`    // validate unique from cache
	TrackNumber       string    `json:"track_number"` // validate unique from cache
	Entry             string    `json:"entry"`
	Delivery          *Delivery `json:"delivery"`
	Payment           *Payment  `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SmID              uint64    `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}
