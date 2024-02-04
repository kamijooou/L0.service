package sql

const InsertOrderSQL = `
insert into orders (
    order_uid,
    track_number,
    entry,
    locale,
    internal_signature,
    customer_id,
    delivery_service,
    shardkey,
    sm_id,
    date_created,
    oof_shard
)
values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11
);
`

const InsertDeliverySQL = `
insert into deliveries (
    order_id,
    name,
    phone,
    zip,
    city,
    address,
    region,
    email
)
values (
    (select currval(pg_get_serial_sequence('orders','id'))),
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7
);
`

const InsertPaymentSQL = `
insert into payments (
    transaction,
    request_id,
    currency,
    provider,
    amount,
    payment_dt,
    bank,
    delivery_cost,
    goods_total,
    custom_fee
)
values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10
);
`

const InsertItemSQL = `
insert into items (
    chrt_id,
    track_number,
    price,
    rid,
    name,
    sale,
    size,
    total_price,
    nm_id,
    brand,
    status
)
values (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6,
    $7,
    $8,
    $9,
    $10,
    $11
);
`
