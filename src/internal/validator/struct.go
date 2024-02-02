package validator

import "time"

// TODO: implement the comments tasks in struct fields

type Item struct {
	ChrtID      uint64 `json:"chrt_id"`      // validate not null
	TrackNumber string `json:"track_number"` // validate equal to order.track_number
	Price       uint64 `json:"price"`        // validate not null
	RID         string `json:"rid"`          // validate not null
	Name        string `json:"name"`         // validate not null
	Sale        uint64 `json:"sale"`
	Size        string `json:"size"`        // validate not null
	TotalPrice  uint64 `json:"total_price"` // validate not null
	NmID        uint64 `json:"nm_id"`       // validate not null
	Brand       string `json:"brand"`
	Status      uint64 `json:"status"` // validate not null
}

type Payment struct {
	Transaction  string `json:"transaction"` // validate equal to order_uid
	RequestID    string `json:"request_id"`
	Currency     string `json:"currency"`      // validate not null
	Provider     string `json:"provider"`      // validate not null
	Amount       uint64 `json:"amount"`        // validate not null
	PaymentDT    uint64 `json:"payment_dt"`    // validate not null
	Bank         string `json:"bank"`          // validate not null
	DeliveryCost uint64 `json:"delivery_cost"` // validate not null
	GoodsTotal   uint64 `json:"goods_total"`   // validate not null
	CustomFee    uint64 `json:"custom_fee"`
}

type Delivery struct {
	Name    string `json:"name"`
	Phone   string `json:"phone"` // validate not null
	Zip     string `json:"zip"`
	City    string `json:"city"`    // validate not null
	Address string `json:"address"` // validate not null
	Region  string `json:"region"`  // validate not null
	Email   string `json:"email"`   // validate not null
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
	CustomerID        string    `json:"customer_id"` // validate not null
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SmID              uint64    `json:"sm_id"`        // validate not null
	DateCreated       time.Time `json:"date_created"` // validate not null
	OofShard          string    `json:"oof_shard"`
}
